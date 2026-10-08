use uuid::Uuid;

use super::{PolicyNumber, Term};

// Entidad: campos privados, se construye a partir de value objects ya validados.
#[derive(Clone)]
pub struct Policy {
    id: Uuid,
    number: PolicyNumber,
    term: Term,
}

impl Policy {
    pub fn new(number: PolicyNumber, term: Term) -> Self {
        Self {
            id: Uuid::new_v4(),
            number,
            term,
        }
    }

    pub fn id(&self) -> Uuid {
        self.id
    }

    pub fn number(&self) -> &PolicyNumber {
        &self.number
    }

    pub fn term(&self) -> Term {
        self.term
    }
}
