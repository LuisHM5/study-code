// Círculo 2: casos de uso. Un caso de uso por archivo, con su Input y su Output.
pub mod gateways;
pub mod get_policy;
pub mod issue_policy;

use crate::domain::DomainError;

#[derive(Debug, thiserror::Error)]
pub enum AppError {
    #[error(transparent)]
    Domain(#[from] DomainError),
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
}
