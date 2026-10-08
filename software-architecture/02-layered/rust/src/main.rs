//! Capas por tipo técnico: handler -> service -> repository -> datos.
//! Las carpetas se llaman por lo que el código ES, no por el negocio.

mod error;
mod handlers {
    pub mod policy;
}
mod models {
    pub mod policy;
}
mod repositories {
    pub mod policy;
}
mod services {
    pub mod policy;
}

use std::sync::Arc;

use axum::{
    routing::{get, post},
    Router,
};

use repositories::policy::PolicyRepository;
use services::policy::PolicyService;

#[tokio::main]
async fn main() {
    let service = Arc::new(PolicyService::new(PolicyRepository::default()));

    let app = Router::new()
        .route("/policies", post(handlers::policy::issue))
        .route("/policies/{id}", get(handlers::policy::get_by_id))
        .with_state(service);

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
