use super::{DomainError, PolicyNumber, PolicyRepository};

// Servicio de dominio: una regla que la entidad sola no puede verificar.
pub async fn assert_unique(
    policies: &dyn PolicyRepository,
    number: &PolicyNumber,
) -> Result<(), DomainError> {
    if policies.exists_by_number(number).await? {
        return Err(DomainError::Duplicated);
    }
    Ok(())
}
