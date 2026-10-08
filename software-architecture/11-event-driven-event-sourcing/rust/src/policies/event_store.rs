use std::{collections::HashMap, sync::Mutex};

use uuid::Uuid;

use super::{aggregate::PolicyError, events::PolicyEvent};

// Un stream de eventos por póliza. Solo se agrega al final, nunca se edita.
#[derive(Default)]
pub struct EventStore {
    streams: Mutex<HashMap<Uuid, Vec<PolicyEvent>>>,
}

impl EventStore {
    pub fn read_stream(&self, id: Uuid) -> Vec<PolicyEvent> {
        self.streams
            .lock()
            .unwrap()
            .get(&id)
            .cloned()
            .unwrap_or_default()
    }

    // Concurrencia optimista: solo agrega si nadie escribió desde que se leyó el stream.
    pub fn append(
        &self,
        id: Uuid,
        expected_version: usize,
        events: &[PolicyEvent],
    ) -> Result<(), PolicyError> {
        let mut streams = self.streams.lock().unwrap();
        let stream = streams.entry(id).or_default();
        if stream.len() != expected_version {
            return Err(PolicyError::Conflict);
        }
        stream.extend_from_slice(events);
        Ok(())
    }
}
