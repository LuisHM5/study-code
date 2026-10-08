use chrono::NaiveDate;

use super::DomainError;

// Value object: una vigencia inválida no puede existir.
#[derive(Clone, Copy, PartialEq, Debug)]
pub struct Term {
    start: NaiveDate,
    end: NaiveDate,
}

impl Term {
    pub fn new(start: NaiveDate, end: NaiveDate) -> Result<Self, DomainError> {
        if end <= start {
            return Err(DomainError::InvalidTerm);
        }
        Ok(Self { start, end })
    }

    pub fn start(&self) -> NaiveDate {
        self.start
    }

    pub fn end(&self) -> NaiveDate {
        self.end
    }
}
