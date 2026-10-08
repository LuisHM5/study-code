use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    response::{IntoResponse, Response},
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use uuid::Uuid;

use super::{
    aggregate::{Policy, PolicyError},
    events::PolicyEvent,
    service::PolicyService,
};

pub fn router(service: Arc<PolicyService>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/{id}", get(get_by_id))
        .route("/policies/{id}/renew", post(renew))
        .route("/policies/{id}/events", get(events))
        .with_state(service)
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssueRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Deserialize)]
struct RenewRequest {
    months: u32,
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
struct PolicyResponse {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

impl PolicyResponse {
    fn new(id: Uuid, policy: Policy) -> Self {
        Self {
            id,
            number: policy.number,
            start_date: policy.start_date,
            end_date: policy.end_date,
        }
    }
}

async fn issue(
    State(service): State<Arc<PolicyService>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let id = service
        .issue(req.number, req.start_date, req.end_date)
        .map_err(to_response)?;
    Ok((StatusCode::CREATED, Json(json!({ "id": id }))))
}

async fn renew(
    State(service): State<Arc<PolicyService>>,
    Path(id): Path<Uuid>,
    body: Result<Json<RenewRequest>, JsonRejection>,
) -> Result<Json<PolicyResponse>, Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let policy = service.renew(id, req.months).map_err(to_response)?;
    Ok(Json(PolicyResponse::new(id, policy)))
}

async fn get_by_id(
    State(service): State<Arc<PolicyService>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyResponse>, Response> {
    let policy = service.get(id).map_err(to_response)?;
    Ok(Json(PolicyResponse::new(id, policy)))
}

async fn events(
    State(service): State<Arc<PolicyService>>,
    Path(id): Path<Uuid>,
) -> Result<Json<Vec<PolicyEvent>>, Response> {
    Ok(Json(service.history(id).map_err(to_response)?))
}

fn to_response(err: PolicyError) -> Response {
    let status = match err {
        PolicyError::InvalidTerm | PolicyError::InvalidMonths => StatusCode::BAD_REQUEST,
        PolicyError::Duplicated | PolicyError::Conflict => StatusCode::CONFLICT,
        PolicyError::NotFound => StatusCode::NOT_FOUND,
    };
    message(status, &err.to_string())
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
