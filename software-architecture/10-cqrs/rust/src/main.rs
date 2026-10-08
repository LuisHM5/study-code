//! CQRS sin bus: los comandos pasan por el dominio y el repositorio de escritura;
//! las consultas leen de un modelo de lectura separado, hecho a la medida.

mod policies;

#[tokio::main]
async fn main() {
    let app = policies::router();

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
