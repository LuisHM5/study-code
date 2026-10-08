use async_trait::async_trait;

use crate::policy::{
    domain::{Policy, PolicyError},
    ports::Notifier,
};

// Adaptador de salida. En producción sería SMTP, SES, etc.
pub struct ConsoleNotifier;

#[async_trait]
impl Notifier for ConsoleNotifier {
    async fn policy_issued(&self, policy: &Policy) -> Result<(), PolicyError> {
        println!("[notificación] póliza emitida {}", policy.number);
        Ok(())
    }
}
