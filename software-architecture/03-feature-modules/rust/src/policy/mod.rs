// API pública de la feature: lo que no se re-exporta aquí es privado.
mod handler;
mod model;
mod repository;
mod service;

use std::sync::Arc;

use axum::{
    routing::{get, post},
    Router,
};

pub use service::{Notifier, PolicyService};

pub fn router(service: Arc<PolicyService>) -> Router {
    Router::new()
        .route("/policies", post(handler::issue))
        .route("/policies/{id}", get(handler::get_by_id))
        .with_state(service)
}
