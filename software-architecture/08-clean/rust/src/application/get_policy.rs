use chrono::NaiveDate;
use uuid::Uuid;

use super::{gateways::PolicyGateway, AppError};

pub struct GetPolicyOutput {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub struct GetPolicy<P> {
    pub policies: P,
}

impl<P: PolicyGateway> GetPolicy<P> {
    pub async fn execute(&self, id: Uuid) -> Result<GetPolicyOutput, AppError> {
        let policy = self
            .policies
            .find_by_id(id)
            .await?
            .ok_or(AppError::NotFound)?;
        Ok(GetPolicyOutput {
            id: policy.id(),
            number: policy.number().to_owned(),
            start_date: policy.start_date(),
            end_date: policy.end_date(),
        })
    }
}
