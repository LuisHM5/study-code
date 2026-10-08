// Lo que los casos de uso necesitan de afuera. Los implementa el círculo 3.
use async_trait::async_trait;
use uuid::Uuid;

use super::AppError;
use crate::domain::Policy;

#[async_trait]
pub trait PolicyGateway: Send + Sync {
    async fn exists_by_number(&self, number: &str) -> Result<bool, AppError>;
    async fn save(&self, policy: &Policy) -> Result<(), AppError>;
    async fn find_by_id(&self, id: Uuid) -> Result<Option<Policy>, AppError>;
}

#[async_trait]
pub trait NotificationGateway: Send + Sync {
    async fn policy_issued(&self, number: &str) -> Result<(), AppError>;
}
