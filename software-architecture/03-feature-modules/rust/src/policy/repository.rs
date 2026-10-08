use std::{collections::HashMap, sync::Mutex};

use uuid::Uuid;

use super::model::{NewPolicy, Policy};

// Privado de la feature: nadie fuera de `policy` puede tocar estos datos.
#[derive(Default)]
pub struct PolicyRepository {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

impl PolicyRepository {
    pub fn exists_by_number(&self, number: &str) -> bool {
        self.policies
            .lock()
            .unwrap()
            .values()
            .any(|p| p.number == number)
    }

    pub fn insert(&self, new: NewPolicy) -> Policy {
        let policy = Policy {
            id: Uuid::new_v4(),
            number: new.number,
            start_date: new.start_date,
            end_date: new.end_date,
        };
        self.policies
            .lock()
            .unwrap()
            .insert(policy.id, policy.clone());
        policy
    }

    pub fn find_by_id(&self, id: Uuid) -> Option<Policy> {
        self.policies.lock().unwrap().get(&id).cloned()
    }
}
