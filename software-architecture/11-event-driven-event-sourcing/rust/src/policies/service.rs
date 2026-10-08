use std::{collections::HashSet, sync::Mutex};

use chrono::NaiveDate;
use tokio::sync::broadcast::Sender;
use uuid::Uuid;

use super::{
    aggregate::{Policy, PolicyError},
    event_store::EventStore,
    events::{PolicyEvent, Published},
};

pub struct PolicyService {
    store: EventStore,
    // Proyección mínima para la regla de unicidad del número.
    numbers: Mutex<HashSet<String>>,
    bus: Sender<Published>,
}

impl PolicyService {
    pub fn new(bus: Sender<Published>) -> Self {
        Self {
            store: EventStore::default(),
            numbers: Mutex::default(),
            bus,
        }
    }

    pub fn issue(
        &self,
        number: String,
        start: NaiveDate,
        end: NaiveDate,
    ) -> Result<Uuid, PolicyError> {
        let events = Policy::issue(number.clone(), start, end)?;
        if !self.numbers.lock().unwrap().insert(number) {
            return Err(PolicyError::Duplicated);
        }

        let id = Uuid::new_v4();
        self.store.append(id, 0, &events)?;
        self.publish(id, events);
        Ok(id)
    }

    pub fn renew(&self, id: Uuid, months: u32) -> Result<Policy, PolicyError> {
        // Cargar = reproducir el historial.
        let history = self.history(id)?;
        let policy = Policy::from_history(&history);

        let events = policy.renew(months)?;
        self.store.append(id, policy.version, &events)?;

        let policy = events.iter().fold(policy, Policy::apply);
        self.publish(id, events);
        Ok(policy)
    }

    pub fn get(&self, id: Uuid) -> Result<Policy, PolicyError> {
        Ok(Policy::from_history(&self.history(id)?))
    }

    pub fn history(&self, id: Uuid) -> Result<Vec<PolicyEvent>, PolicyError> {
        let events = self.store.read_stream(id);
        if events.is_empty() {
            return Err(PolicyError::NotFound);
        }
        Ok(events)
    }

    fn publish(&self, policy_id: Uuid, events: Vec<PolicyEvent>) {
        for event in events {
            let _ = self.bus.send(Published { policy_id, event });
        }
    }
}
