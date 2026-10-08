use std::sync::Arc;

use chrono::NaiveDate;
use uuid::Uuid;

use super::{
    domain::{Policy, PolicyError},
    ports::{Notifier, PolicyRepository},
};

// Caso de uso: solo conoce el dominio y los puertos.
pub struct IssuePolicy {
    pub policies: Arc<dyn PolicyRepository>,
    pub notifier: Arc<dyn Notifier>,
}

impl IssuePolicy {
    pub async fn execute(
        &self,
        number: String,
        start: NaiveDate,
        end: NaiveDate,
    ) -> Result<Uuid, PolicyError> {
        if self.policies.exists_by_number(&number).await? {
            return Err(PolicyError::Duplicated);
        }
        let policy = Policy::issue(number, start, end)?;
        self.policies.save(&policy).await?;
        self.notifier.policy_issued(&policy).await?;
        Ok(policy.id)
    }
}
