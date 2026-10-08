use std::{collections::HashMap, sync::Mutex};

use uuid::Uuid;

use crate::policies::domain::Policy;

// Almacén de escritura: guarda el agregado. No se usa para consultar pantallas.
#[derive(Default)]
pub struct WriteRepository {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

impl WriteRepository {
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
}
