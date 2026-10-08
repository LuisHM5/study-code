use chrono::NaiveDate;
use uuid::Uuid;

use super::{
    gateways::{NotificationGateway, PolicyGateway},
    AppError,
};
use crate::domain::Policy;

// Modelos propios de la frontera: no son el JSON de HTTP ni la entidad.
pub struct IssuePolicyInput {
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub struct IssuePolicyOutput {
    pub id: Uuid,
}

// Genérico sobre sus gateways: en tests se sustituyen por dobles sin tocar nada.
pub struct IssuePolicy<P, N> {
    pub policies: P,
    pub notifier: N,
}

impl<P: PolicyGateway, N: NotificationGateway> IssuePolicy<P, N> {
    pub async fn execute(&self, input: IssuePolicyInput) -> Result<IssuePolicyOutput, AppError> {
        if self.policies.exists_by_number(&input.number).await? {
            return Err(AppError::Duplicated);
        }
        let policy = Policy::new(input.number, input.start_date, input.end_date)?;
        self.policies.save(&policy).await?;
        self.notifier.policy_issued(policy.number()).await?;
        Ok(IssuePolicyOutput { id: policy.id() })
    }
}
