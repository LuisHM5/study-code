// Dominio puro: sin axum, sin base de datos, sin serde.
use chrono::NaiveDate;
use uuid::Uuid;

#[derive(Debug, thiserror::Error)]
pub enum PolicyError {
    #[error("vigencia inválida")]
    InvalidTerm,
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
}

#[derive(Clone)]
pub struct Policy {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

impl Policy {
    pub fn issue(
        number: String,
        start_date: NaiveDate,
        end_date: NaiveDate,
    ) -> Result<Self, PolicyError> {
        if end_date <= start_date {
            return Err(PolicyError::InvalidTerm);
        }
        Ok(Self {
            id: Uuid::new_v4(),
            number,
            start_date,
            end_date,
        })
    }
}
