use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use serde::Serialize;
use uuid::Uuid;

use crate::policies::PoliciesApi;

// El siniestro guarda el id y el número de la póliza como datos propios,
// sin referenciar la entidad interna de policies.
#[derive(Clone, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Claim {
    pub id: Uuid,
    pub policy_id: Uuid,
    pub policy_number: String,
    pub description: String,
}

pub struct PolicyNotFound;

pub struct Claims {
    policies: Arc<dyn PoliciesApi>,
    claims: Mutex<HashMap<Uuid, Claim>>,
}

impl Claims {
    pub fn new(policies: Arc<dyn PoliciesApi>) -> Self {
        Self {
            policies,
            claims: Mutex::default(),
        }
    }

    pub fn report(&self, policy_id: Uuid, description: String) -> Result<Claim, PolicyNotFound> {
        let policy = self.policies.find_by_id(policy_id).ok_or(PolicyNotFound)?;

        let claim = Claim {
            id: Uuid::new_v4(),
            policy_id: policy.id,
            policy_number: policy.number,
            description,
        };
        self.claims.lock().unwrap().insert(claim.id, claim.clone());
        Ok(claim)
    }

    pub fn get(&self, id: Uuid) -> Option<Claim> {
        self.claims.lock().unwrap().get(&id).cloned()
    }
}
