//! Modular por feature: una carpeta por concepto de negocio.
//! main.rs solo arma los módulos; cada uno expone lo mínimo con `pub`.

mod notification;
mod platform;
mod policy;

use std::sync::Arc;

use axum::Router;

#[tokio::main]
async fn main() {
    let notifier = Arc::new(notification::NotificationService);
    let policy_service = Arc::new(policy::PolicyService::new(notifier));

    // Cada feature aporta su propio router.
    let app = Router::new().merge(policy::router(policy_service));

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
