use async_trait::async_trait;

use crate::{
    application::PolicyNotifier,
    domain::{DomainError, Policy},
};

pub struct ConsoleNotifier;

#[async_trait]
impl PolicyNotifier for ConsoleNotifier {
    async fn policy_issued(&self, policy: &Policy) -> Result<(), DomainError> {
        println!("[notificación] póliza emitida {}", policy.number().as_str());
        Ok(())
    }
}
