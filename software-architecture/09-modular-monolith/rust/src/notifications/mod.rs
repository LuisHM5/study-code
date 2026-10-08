// Módulo de notificaciones: reacciona a eventos. policies no sabe que existe.
use tokio::sync::broadcast::Receiver;

use crate::policies::PolicyIssued;

pub fn start(mut events: Receiver<PolicyIssued>) {
    tokio::spawn(async move {
        while let Ok(event) = events.recv().await {
            println!(
                "[notificación] póliza emitida {} (id {})",
                event.number, event.policy_id
            );
        }
    });
}
