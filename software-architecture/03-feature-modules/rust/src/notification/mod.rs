use crate::policy::Notifier;

pub struct NotificationService;

// La feature de notificaciones implementa el contrato que pide `policy`.
impl Notifier for NotificationService {
    fn policy_issued(&self, number: &str) {
        println!("[notificación] póliza emitida {number}");
    }
}
