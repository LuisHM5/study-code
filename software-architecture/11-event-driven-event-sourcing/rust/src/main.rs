//! Event-Driven + Event Sourcing.
//! - Event sourcing: de la póliza no se guarda el estado, se guardan sus eventos.
//! - Event-driven: cada evento se publica en un bus y notifications reacciona.

mod notifications;
mod policies;

#[tokio::main]
async fn main() {
    let (bus, _) = tokio::sync::broadcast::channel(16);

    notifications::start(bus.subscribe());
    let app = policies::router(bus);

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
