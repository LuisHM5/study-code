//! Screaming Architecture: el árbol dice "pólizas: emisión y consulta",
//! no "handlers, services, models". Lo técnico queda aparte en platform/.

mod platform;
mod policies;

use std::sync::Arc;

use axum::{
    routing::{get, post},
    Router,
};

use policies::PolicyBook;

#[tokio::main]
async fn main() {
    let book = Arc::new(PolicyBook::default());

    let app = Router::new()
        .route("/policies", post(platform::http::issue_policy))
        .route("/policies/{id}", get(platform::http::find_policy))
        .with_state(book);

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
