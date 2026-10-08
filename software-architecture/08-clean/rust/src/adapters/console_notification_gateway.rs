use async_trait::async_trait;

use crate::application::{gateways::NotificationGateway, AppError};

pub struct ConsoleNotificationGateway;

#[async_trait]
impl NotificationGateway for ConsoleNotificationGateway {
    async fn policy_issued(&self, number: &str) -> Result<(), AppError> {
        println!("[notificación] póliza emitida {number}");
        Ok(())
    }
}
