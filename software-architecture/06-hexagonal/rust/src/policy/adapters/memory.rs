use std::{collections::HashMap, sync::Mutex};

use async_trait::async_trait;
use uuid::Uuid;

use crate::policy::{
    domain::{Policy, PolicyError},
    ports::PolicyRepository,
};

// Adaptador de salida. En producción sería PgPolicyRepository con sqlx.
#[derive(Default)]
pub struct InMemoryPolicyRepository {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

#[async_trait]
impl PolicyRepository for InMemoryPolicyRepository {
    async fn exists_by_number(&self, number: &str) -> Result<bool, PolicyError> {
        let policies = self.policies.lock().unwrap();
        Ok(policies.values().any(|p| p.number == number))
    }

    async fn save(&self, policy: &Policy) -> Result<(), PolicyError> {
        self.policies
            .lock()
            .unwrap()
            .insert(policy.id, policy.clone());
        Ok(())
    }

    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, PolicyError> {
        Ok(self.policies.lock().unwrap().get(&id).cloned())
    }
}
