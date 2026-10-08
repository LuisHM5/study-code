// Suscriptor: reacciona a los eventos. Quien los publica no sabe que existe.
use tokio::sync::broadcast::Receiver;

use crate::policies::{PolicyEvent, Published};

pub fn start(mut bus: Receiver<Published>) {
    tokio::spawn(async move {
        while let Ok(Published { policy_id, event }) = bus.recv().await {
            match event {
                PolicyEvent::PolicyIssued { number, .. } => {
                    println!("[notificación] póliza emitida {number}");
                }
                PolicyEvent::PolicyRenewed { new_end_date } => {
                    println!("[notificación] póliza {policy_id} renovada hasta {new_end_date}");
                }
            }
        }
    });
}
