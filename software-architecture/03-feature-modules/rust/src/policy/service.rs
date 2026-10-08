use std::sync::Arc;

use uuid::Uuid;

use super::{
    model::{NewPolicy, Policy},
    repository::PolicyRepository,
};
use crate::platform::error::AppError;

// El contrato se declara donde se consume; otra feature lo implementa.
pub trait Notifier: Send + Sync {
    fn policy_issued(&self, number: &str);
}

pub struct PolicyService {
    repo: PolicyRepository,
    notifier: Arc<dyn Notifier>,
}

impl PolicyService {
    pub fn new(notifier: Arc<dyn Notifier>) -> Self {
        Self {
            repo: PolicyRepository::default(),
            notifier,
        }
    }

    pub fn issue(&self, new: NewPolicy) -> Result<Policy, AppError> {
        if new.end_date <= new.start_date {
            return Err(AppError::InvalidTerm);
        }
        if self.repo.exists_by_number(&new.number) {
            return Err(AppError::Duplicated);
        }
        let policy = self.repo.insert(new);
        self.notifier.policy_issued(&policy.number);
        Ok(policy)
    }

    pub fn get(&self, id: Uuid) -> Result<Policy, AppError> {
        self.repo.find_by_id(id).ok_or(AppError::NotFound)
    }
}
