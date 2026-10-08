//! Transaction Script: cada handler hace todo, de arriba a abajo.
//! Sin capas: validación, reglas, datos y notificación en el mismo lugar.

use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use uuid::Uuid;

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssuePolicy {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Clone, Serialize)]
#[serde(rename_all = "camelCase")]
struct Policy {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

// La "base de datos": un mapa en memoria.
type Db = Arc<Mutex<HashMap<Uuid, Policy>>>;
type ApiError = (StatusCode, Json<Value>);

fn error(status: StatusCode, message: &str) -> ApiError {
    (status, Json(json!({ "message": message })))
}

async fn issue_policy(
    State(db): State<Db>,
    body: Result<Json<IssuePolicy>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), ApiError> {
    let Json(input) = body.map_err(|_| error(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;

    if input.end_date <= input.start_date {
        return Err(error(StatusCode::BAD_REQUEST, "vigencia inválida"));
    }

    let mut policies = db.lock().unwrap();
    if policies.values().any(|p| p.number == input.number) {
        return Err(error(StatusCode::CONFLICT, "número duplicado"));
    }

    let policy = Policy {
        id: Uuid::new_v4(),
        number: input.number,
        start_date: input.start_date,
        end_date: input.end_date,
    };
    policies.insert(policy.id, policy.clone());

    println!("[notificación] póliza emitida {}", policy.number);
    Ok((StatusCode::CREATED, Json(json!({ "id": policy.id }))))
}

async fn get_policy(State(db): State<Db>, Path(id): Path<Uuid>) -> Result<Json<Policy>, ApiError> {
    db.lock()
        .unwrap()
        .get(&id)
        .cloned()
        .map(Json)
        .ok_or_else(|| error(StatusCode::NOT_FOUND, "póliza no encontrada"))
}

#[tokio::main]
async fn main() {
    let db: Db = Arc::default();
    let app = Router::new()
        .route("/policies", post(issue_policy))
        .route("/policies/{id}", get(get_policy))
        .with_state(db);

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
