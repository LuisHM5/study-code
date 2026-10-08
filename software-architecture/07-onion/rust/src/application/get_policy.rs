use std::sync::Arc;

use chrono::NaiveDate;
use uuid::Uuid;

use crate::domain::{DomainError, PolicyRepository};

// DTO de salida: la entidad de dominio no sale del anillo de aplicación.
pub struct PolicyDto {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub struct GetPolicyService {
    pub policies: Arc<dyn PolicyRepository>,
}

impl GetPolicyService {
    pub async fn execute(&self, id: Uuid) -> Result<PolicyDto, DomainError> {
        let policy = self
            .policies
            .find_by_id(id)
            .await?
            .ok_or(DomainError::NotFound)?;
        Ok(PolicyDto {
            id: policy.id(),
            number: policy.number().as_str().to_owned(),
            start_date: policy.term().start(),
            end_date: policy.term().end(),
        })
    }
}
