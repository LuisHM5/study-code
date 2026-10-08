use chrono::{Months, NaiveDate};

use super::events::PolicyEvent;

#[derive(Debug, thiserror::Error)]
pub enum PolicyError {
    #[error("vigencia inválida")]
    InvalidTerm,
    #[error("los meses de renovación deben ser mayores a cero")]
    InvalidMonths,
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
    #[error("la póliza fue modificada por otra operación, reintenta")]
    Conflict,
}

// Estado actual de una póliza. Nunca se guarda: se calcula a partir de sus eventos.
#[derive(Default)]
pub struct Policy {
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
    pub version: usize,
}

impl Policy {
    // DECIDIR: valida y devuelve los eventos resultantes, sin mutar nada.
    pub fn issue(
        number: String,
        start_date: NaiveDate,
        end_date: NaiveDate,
    ) -> Result<Vec<PolicyEvent>, PolicyError> {
        if end_date <= start_date {
            return Err(PolicyError::InvalidTerm);
        }
        Ok(vec![PolicyEvent::PolicyIssued {
            number,
            start_date,
            end_date,
        }])
    }

    pub fn renew(&self, months: u32) -> Result<Vec<PolicyEvent>, PolicyError> {
        if months == 0 {
            return Err(PolicyError::InvalidMonths);
        }
        Ok(vec![PolicyEvent::PolicyRenewed {
            new_end_date: self.end_date + Months::new(months),
        }])
    }

    // APLICAR: transición de estado pura. Único lugar donde cambia el estado.
    pub fn apply(mut self, event: &PolicyEvent) -> Self {
        match event {
            PolicyEvent::PolicyIssued {
                number,
                start_date,
                end_date,
            } => {
                self.number = number.clone();
                self.start_date = *start_date;
                self.end_date = *end_date;
            }
            PolicyEvent::PolicyRenewed { new_end_date } => self.end_date = *new_end_date,
        }
        self.version += 1;
        self
    }

    // El estado es un fold sobre el historial.
    pub fn from_history(events: &[PolicyEvent]) -> Self {
        events.iter().fold(Self::default(), Self::apply)
    }
}
