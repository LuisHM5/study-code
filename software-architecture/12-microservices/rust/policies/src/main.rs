//! Servicio de pólizas: dueño de sus datos. Para notificar llama por red
//! al servicio de notificaciones; si este falla, la emisión no se cae.

mod http;
mod notifications_client;
mod policy;

use std::sync::Arc;

use notifications_client::NotificationsClient;
use policy::PolicyStore;

#[tokio::main]
async fn main() {
    let notifications_url =
        std::env::var("NOTIFICATIONS_URL").unwrap_or_else(|_| "http://localhost:3002".into());

    let app = http::router(Arc::new(http::AppState {
        policies: PolicyStore::default(),
        notifications: NotificationsClient::new(notifications_url),
    }));

    let port = std::env::var("PORT").unwrap_or_else(|_| "3001".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("policies escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
