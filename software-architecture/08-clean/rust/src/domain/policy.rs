use chrono::NaiveDate;
use uuid::Uuid;

#[derive(Debug, thiserror::Error)]
pub enum DomainError {
    #[error("vigencia inválida")]
    InvalidTerm,
}

#[derive(Clone)]
pub struct Policy {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

impl Policy {
    pub fn new(
        number: String,
        start_date: NaiveDate,
        end_date: NaiveDate,
    ) -> Result<Self, DomainError> {
        if end_date <= start_date {
            return Err(DomainError::InvalidTerm);
        }
        Ok(Self {
            id: Uuid::new_v4(),
            number,
            start_date,
            end_date,
        })
    }

    pub fn id(&self) -> Uuid {
        self.id
    }

    pub fn number(&self) -> &str {
        &self.number
    }

    pub fn start_date(&self) -> NaiveDate {
        self.start_date
    }

    pub fn end_date(&self) -> NaiveDate {
        self.end_date
    }
}
