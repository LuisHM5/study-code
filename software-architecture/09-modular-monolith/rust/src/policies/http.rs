use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    response::Response,
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use uuid::Uuid;

use super::{app::Service, domain::PolicyError};
use crate::platform::error;

pub fn router(service: Arc<Service>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/{id}", get(get_by_id))
        .with_state(service)
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssueRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
struct PolicyResponse {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

async fn issue(
    State(service): State<Arc<Service>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| error(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let policy = service
        .issue(req.number, req.start_date, req.end_date)
        .map_err(to_response)?;
    Ok((StatusCode::CREATED, Json(json!({ "id": policy.id }))))
}

async fn get_by_id(
    State(service): State<Arc<Service>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyResponse>, Response> {
    let policy = service.get(id).map_err(to_response)?;
    Ok(Json(PolicyResponse {
        id: policy.id,
        number: policy.number,
        start_date: policy.start_date,
        end_date: policy.end_date,
    }))
}

fn to_response(err: PolicyError) -> Response {
    let status = match err {
        PolicyError::InvalidTerm => StatusCode::BAD_REQUEST,
        PolicyError::Duplicated => StatusCode::CONFLICT,
        PolicyError::NotFound => StatusCode::NOT_FOUND,
    };
    error(status, &err.to_string())
}
