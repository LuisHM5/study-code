use std::sync::Arc;

use chrono::NaiveDate;
use uuid::Uuid;

use super::PolicyNotifier;
use crate::domain::{uniqueness, DomainError, Policy, PolicyNumber, PolicyRepository, Term};

pub struct IssuePolicyCommand {
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub struct IssuePolicyService {
    pub policies: Arc<dyn PolicyRepository>,
    pub notifier: Arc<dyn PolicyNotifier>,
}

impl IssuePolicyService {
    pub async fn execute(&self, cmd: IssuePolicyCommand) -> Result<Uuid, DomainError> {
        let number = PolicyNumber::parse(&cmd.number)?;
        let term = Term::new(cmd.start_date, cmd.end_date)?;
        uniqueness::assert_unique(self.policies.as_ref(), &number).await?;

        let policy = Policy::new(number, term);
        self.policies.save(&policy).await?;
        self.notifier.policy_issued(&policy).await?;
        Ok(policy.id())
    }
}
