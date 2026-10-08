// Anillo central: modelo de dominio, servicios de dominio e interfaces de repositorio.
// No depende de ningún otro módulo del proyecto.
mod number;
mod policy;
mod repository;
mod term;
pub mod uniqueness;

pub use number::PolicyNumber;
pub use policy::Policy;
pub use repository::PolicyRepository;
pub use term::Term;

#[derive(Debug, thiserror::Error)]
pub enum DomainError {
    #[error("el número de póliza no puede estar vacío")]
    EmptyNumber,
    #[error("vigencia inválida")]
    InvalidTerm,
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
}
