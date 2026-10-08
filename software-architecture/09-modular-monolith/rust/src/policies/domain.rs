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
    pub fn issue(number: String, start: NaiveDate, end: NaiveDate) -> Result<Self, PolicyError> {
        if end <= start {
            return Err(PolicyError::InvalidTerm);
        }
        Ok(Self {
            id: Uuid::new_v4(),
            number,
            start_date: start,
            end_date: end,
        })
    }
}
