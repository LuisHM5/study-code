//! Servicio de notificaciones: proceso y despliegue independientes.
//! Otros servicios le hablan solo por red (HTTP + JSON).

use axum::{extract::rejection::JsonRejection, http::StatusCode, routing::post, Json, Router};
use serde::Deserialize;
use serde_json::{json, Value};
use uuid::Uuid;

// Contrato del servicio. Cambiarlo afecta a todos sus clientes.
#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct PolicyIssued {
    policy_id: Uuid,
    number: String,
}

async fn notify(body: Result<Json<PolicyIssued>, JsonRejection>) -> (StatusCode, Json<Value>) {
    let Ok(Json(event)) = body else {
        return (
            StatusCode::BAD_REQUEST,
            Json(json!({ "message": "cuerpo inválido" })),
        );
    };
    println!(
        "[notificación] póliza emitida {} (id {})",
        event.number, event.policy_id
    );
    (StatusCode::ACCEPTED, Json(json!({ "status": "accepted" })))
}

#[tokio::main]
async fn main() {
    let app = Router::new().route("/notifications", post(notify));

    let port = std::env::var("PORT").unwrap_or_else(|_| "3002".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("notifications escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
