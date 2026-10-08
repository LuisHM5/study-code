use std::{collections::HashMap, sync::Mutex};

use async_trait::async_trait;
use uuid::Uuid;

use crate::domain::{DomainError, Policy, PolicyNumber, PolicyRepository};

#[derive(Default)]
pub struct InMemoryPolicyRepository {
    policies: Mutex<HashMap<Uuid, Policy>>,
}

#[async_trait]
impl PolicyRepository for InMemoryPolicyRepository {
    async fn exists_by_number(&self, number: &PolicyNumber) -> Result<bool, DomainError> {
        let policies = self.policies.lock().unwrap();
        Ok(policies.values().any(|p| p.number() == number))
    }

    async fn save(&self, policy: &Policy) -> Result<(), DomainError> {
        self.policies
            .lock()
            .unwrap()
            .insert(policy.id(), policy.clone());
        Ok(())
    }

    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, DomainError> {
        Ok(self.policies.lock().unwrap().get(&id).cloned())
    }
}
