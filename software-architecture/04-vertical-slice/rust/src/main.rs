//! Vertical Slice: un archivo por caso de uso, con su request, su lógica y su handler.

mod app;
mod policy;

use axum::Router;

#[tokio::main]
async fn main() {
    let app = Router::new()
        .merge(policy::router())
        .with_state(app::AppContext::default());

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
