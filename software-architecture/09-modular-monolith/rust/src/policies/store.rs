use std::{collections::HashMap, sync::Mutex};

use uuid::Uuid;

use super::domain::Policy;

// Los datos de pólizas son de este módulo: nadie más puede leerlos ni escribirlos.
#[derive(Default)]
pub struct Store {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

impl Store {
    pub fn exists_by_number(&self, number: &str) -> bool {
        self.policies
            .lock()
            .unwrap()
            .values()
            .any(|p| p.number == number)
    }

    pub fn save(&self, policy: &Policy) {
        self.policies
            .lock()
            .unwrap()
            .insert(policy.id, policy.clone());
    }

    pub fn find_by_id(&self, id: Uuid) -> Option<Policy> {
        self.policies.lock().unwrap().get(&id).cloned()
    }
}
