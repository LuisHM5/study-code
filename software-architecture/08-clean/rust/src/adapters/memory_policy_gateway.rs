use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use async_trait::async_trait;
use uuid::Uuid;

use crate::{
    application::{gateways::PolicyGateway, AppError},
    domain::Policy,
};

// Clonar comparte el mismo almacén (como clonar un pool de conexiones).
#[derive(Clone, Default)]
pub struct MemoryPolicyGateway {
    policies: Arc<Mutex<HashMap<Uuid, Policy>>>,
}

#[async_trait]
impl PolicyGateway for MemoryPolicyGateway {
    async fn exists_by_number(&self, number: &str) -> Result<bool, AppError> {
        let policies = self.policies.lock().unwrap();
        Ok(policies.values().any(|p| p.number() == number))
    }

    async fn save(&self, policy: &Policy) -> Result<(), AppError> {
        self.policies
            .lock()
            .unwrap()
            .insert(policy.id(), policy.clone());
        Ok(())
    }

    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, AppError> {
        Ok(self.policies.lock().unwrap().get(&id).cloned())
    }
}
