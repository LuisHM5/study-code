use uuid::Uuid;

use crate::{
    error::AppError,
    models::policy::{NewPolicy, Policy},
    repositories::policy::PolicyRepository,
};

// Lógica de negocio. No sabe nada de HTTP.
pub struct PolicyService {
    repo: PolicyRepository,
}

impl PolicyService {
    pub fn new(repo: PolicyRepository) -> Self {
        Self { repo }
    }

    pub fn issue(&self, new: NewPolicy) -> Result<Policy, AppError> {
        if new.end_date <= new.start_date {
            return Err(AppError::InvalidTerm);
        }
        if self.repo.exists_by_number(&new.number) {
            return Err(AppError::Duplicated);
        }
        let policy = self.repo.insert(new);
        println!("[notificación] póliza emitida {}", policy.number);
        Ok(policy)
    }

    pub fn get(&self, id: Uuid) -> Result<Policy, AppError> {
        self.repo.find_by_id(id).ok_or(AppError::NotFound)
    }
}
