use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    response::{IntoResponse, Response},
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::Deserialize;
use serde_json::{json, Value};
use uuid::Uuid;

use crate::{
    notifications_client::NotificationsClient,
    policy::{Policy, PolicyError, PolicyStore},
};

pub struct AppState {
    pub policies: PolicyStore,
    pub notifications: NotificationsClient,
}

pub fn router(state: Arc<AppState>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/{id}", get(get_by_id))
        .with_state(state)
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssueRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

async fn issue(
    State(state): State<Arc<AppState>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;

    let policy = state
        .policies
        .issue(req.number, req.start_date, req.end_date)
        .map_err(|err| match err {
            PolicyError::InvalidTerm => message(StatusCode::BAD_REQUEST, "vigencia inválida"),
            PolicyError::Duplicated => message(StatusCode::CONFLICT, "número duplicado"),
        })?;

    state.notifications.policy_issued(&policy).await;
    Ok((StatusCode::CREATED, Json(json!({ "id": policy.id }))))
}

async fn get_by_id(
    State(state): State<Arc<AppState>>,
    Path(id): Path<Uuid>,
) -> Result<Json<Policy>, Response> {
    state
        .policies
        .find_by_id(id)
        .map(Json)
        .ok_or_else(|| message(StatusCode::NOT_FOUND, "póliza no encontrada"))
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
