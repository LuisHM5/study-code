use super::DomainError;

// Value object: solo se puede construir válido.
#[derive(Clone, PartialEq, Debug)]
pub struct PolicyNumber(String);

impl PolicyNumber {
    pub fn parse(raw: &str) -> Result<Self, DomainError> {
        let value = raw.trim();
        if value.is_empty() {
            return Err(DomainError::EmptyNumber);
        }
        Ok(Self(value.to_owned()))
    }

    pub fn as_str(&self) -> &str {
        &self.0
    }
}
