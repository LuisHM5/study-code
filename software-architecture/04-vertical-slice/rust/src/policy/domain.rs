use chrono::NaiveDate;
use uuid::Uuid;

use crate::app::AppError;

// Lo único compartido entre slices: el dato y las reglas puras.
#[derive(Clone)]
pub struct Policy {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub fn validate_term(start: NaiveDate, end: NaiveDate) -> Result<(), AppError> {
    if end <= start {
        return Err(AppError::InvalidTerm);
    }
    Ok(())
}
