use std::{collections::HashMap, sync::Mutex};

use chrono::NaiveDate;
use uuid::Uuid;

#[derive(Debug, thiserror::Error)]
pub enum PolicyError {
    #[error("el número de póliza no puede estar vacío")]
    EmptyNumber,
    #[error("vigencia inválida")]
    InvalidTerm,
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
}

// Los tipos también gritan: un PolicyNumber no es "un String cualquiera".
#[derive(Clone, PartialEq)]
pub struct PolicyNumber(String);

impl PolicyNumber {
    pub fn parse(raw: &str) -> Result<Self, PolicyError> {
        let value = raw.trim();
        if value.is_empty() {
            return Err(PolicyError::EmptyNumber);
        }
        Ok(Self(value.to_owned()))
    }

    pub fn as_str(&self) -> &str {
        &self.0
    }
}

#[derive(Clone, Copy)]
pub struct PolicyTerm {
    start: NaiveDate,
    end: NaiveDate,
}

impl PolicyTerm {
    pub fn new(start: NaiveDate, end: NaiveDate) -> Result<Self, PolicyError> {
        if end <= start {
            return Err(PolicyError::InvalidTerm);
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

#[derive(Clone)]
pub struct Policy {
    pub id: Uuid,
    pub number: PolicyNumber,
    pub term: PolicyTerm,
}

// El libro de pólizas emitidas (en memoria).
#[derive(Default)]
pub struct PolicyBook {
    pub(super) policies: Mutex<HashMap<Uuid, Policy>>,
}
