use async_trait::async_trait;
use uuid::Uuid;

use super::{DomainError, Policy, PolicyNumber};

// La interfaz vive en el dominio; la implementación, en infrastructure.
#[async_trait]
pub trait PolicyRepository: Send + Sync {
    async fn exists_by_number(&self, number: &PolicyNumber) -> Result<bool, DomainError>;
    async fn save(&self, policy: &Policy) -> Result<(), DomainError>;
    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, DomainError>;
}
