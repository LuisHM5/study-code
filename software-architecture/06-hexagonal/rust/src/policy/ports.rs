// Puertos de salida: lo que el núcleo necesita del mundo exterior.
use async_trait::async_trait;
use uuid::Uuid;

use super::domain::{Policy, PolicyError};

#[async_trait]
pub trait PolicyRepository: Send + Sync {
    async fn exists_by_number(&self, number: &str) -> Result<bool, PolicyError>;
    async fn save(&self, policy: &Policy) -> Result<(), PolicyError>;
    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, PolicyError>;
}

#[async_trait]
pub trait Notifier: Send + Sync {
    async fn policy_issued(&self, policy: &Policy) -> Result<(), PolicyError>;
}
