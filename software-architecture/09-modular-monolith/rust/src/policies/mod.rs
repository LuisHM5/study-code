// API PÚBLICA del módulo de pólizas. Todo lo demás (`mod` sin `pub`) es
// inaccesible desde claims o notifications: el compilador hace cumplir la frontera.
mod app;
mod domain;
mod http;
mod store;

use std::sync::Arc;

use axum::Router;
use tokio::sync::broadcast;
use uuid::Uuid;

/// Contrato hacia otros módulos: un resumen, nunca la entidad interna.
pub struct PolicySummary {
    pub id: Uuid,
    pub number: String,
}

/// Evento público.
#[derive(Clone, Debug)]
pub struct PolicyIssued {
    pub policy_id: Uuid,
    pub number: String,
}

pub type PolicyEvents = broadcast::Sender<PolicyIssued>;

pub trait PoliciesApi: Send + Sync {
    fn find_by_id(&self, id: Uuid) -> Option<PolicySummary>;
}

pub struct PoliciesModule {
    pub api: Arc<dyn PoliciesApi>,
    pub router: Router,
}

pub fn build(events: PolicyEvents) -> PoliciesModule {
    let service = Arc::new(app::Service::new(events));
    PoliciesModule {
        api: service.clone(),
        router: http::router(service),
    }
}
