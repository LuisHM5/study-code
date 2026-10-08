// Servicios de aplicación: orquestan un caso de uso; las reglas viven en el dominio.
pub mod get_policy;
pub mod issue_policy;

use async_trait::async_trait;

use crate::domain::{DomainError, Policy};

#[async_trait]
pub trait PolicyNotifier: Send + Sync {
    async fn policy_issued(&self, policy: &Policy) -> Result<(), DomainError>;
}
