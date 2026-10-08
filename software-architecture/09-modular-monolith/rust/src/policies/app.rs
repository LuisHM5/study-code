use chrono::NaiveDate;
use uuid::Uuid;

use super::{
    domain::{Policy, PolicyError},
    store::Store,
    PoliciesApi, PolicyEvents, PolicyIssued, PolicySummary,
};

pub struct Service {
    store: Store,
    events: PolicyEvents,
}

impl Service {
    pub fn new(events: PolicyEvents) -> Self {
        Self {
            store: Store::default(),
            events,
        }
    }

    pub fn issue(
        &self,
        number: String,
        start: NaiveDate,
        end: NaiveDate,
    ) -> Result<Policy, PolicyError> {
        if self.store.exists_by_number(&number) {
            return Err(PolicyError::Duplicated);
        }
        let policy = Policy::issue(number, start, end)?;
        self.store.save(&policy);

        // Publica el hecho; no sabe quién escucha.
        let _ = self.events.send(PolicyIssued {
            policy_id: policy.id,
            number: policy.number.clone(),
        });
        Ok(policy)
    }

    pub fn get(&self, id: Uuid) -> Result<Policy, PolicyError> {
        self.store.find_by_id(id).ok_or(PolicyError::NotFound)
    }
}

// La fachada pública se implementa sobre el servicio interno.
impl PoliciesApi for Service {
    fn find_by_id(&self, id: Uuid) -> Option<PolicySummary> {
        self.store.find_by_id(id).map(|p| PolicySummary {
            id: p.id,
            number: p.number,
        })
    }
}
