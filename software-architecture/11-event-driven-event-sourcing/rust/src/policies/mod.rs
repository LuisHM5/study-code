mod aggregate;
mod event_store;
mod events;
mod http;
mod service;

use std::sync::Arc;

use axum::Router;
use tokio::sync::broadcast::Sender;

pub use events::{PolicyEvent, Published};

pub fn router(bus: Sender<Published>) -> Router {
    http::router(Arc::new(service::PolicyService::new(bus)))
}
