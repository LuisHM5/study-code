use std::{collections::HashMap, sync::Mutex};

use chrono::NaiveDate;
use serde::Serialize;
use uuid::Uuid;

#[derive(Clone, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Policy {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub enum PolicyError {
    InvalidTerm,
    Duplicated,
}

// Los datos de este servicio. Ningún otro servicio puede leerlos directamente.
#[derive(Default)]
pub struct PolicyStore {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

impl PolicyStore {
    pub fn issue(
        &self,
        number: String,
        start_date: NaiveDate,
        end_date: NaiveDate,
    ) -> Result<Policy, PolicyError> {
        if end_date <= start_date {
            return Err(PolicyError::InvalidTerm);
        }
        let mut policies = self.policies.lock().unwrap();
        if policies.values().any(|p| p.number == number) {
            return Err(PolicyError::Duplicated);
        }
        let policy = Policy {
            id: Uuid::new_v4(),
            number,
            start_date,
            end_date,
        };
        policies.insert(policy.id, policy.clone());
        Ok(policy)
    }

    pub fn find_by_id(&self, id: Uuid) -> Option<Policy> {
        self.policies.lock().unwrap().get(&id).cloned()
    }
}
