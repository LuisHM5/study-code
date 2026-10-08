use std::sync::Arc;

use uuid::Uuid;

use super::{
    domain::{Policy, PolicyError},
    ports::PolicyRepository,
};

pub struct GetPolicy {
    pub policies: Arc<dyn PolicyRepository>,
}

impl GetPolicy {
    pub async fn execute(&self, id: Uuid) -> Result<Policy, PolicyError> {
        self.policies
            .find_by_id(id)
            .await?
            .ok_or(PolicyError::NotFound)
    }
}
