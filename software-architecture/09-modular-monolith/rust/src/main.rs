//! Monolito modular: un solo proceso, módulos con fronteras estrictas.
//! Cada módulo expone una API pública mínima; su interior es privado.
//! En un proyecto grande cada módulo sería un crate del workspace.

mod claims;
mod notifications;
mod platform;
mod policies;

use axum::Router;

#[tokio::main]
async fn main() {
    // Bus de eventos en proceso: policies publica, notifications escucha.
    let (events, _) = tokio::sync::broadcast::channel(16);

    notifications::start(events.subscribe());
    let policies = policies::build(events);
    // claims recibe solo la API pública de policies.
    let claims = claims::build(policies.api.clone());

    let app = Router::new().merge(policies.router).merge(claims);

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
