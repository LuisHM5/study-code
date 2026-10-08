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

use crate::policy::{domain::PolicyError, get_policy::GetPolicy, issue_policy::IssuePolicy};

pub struct UseCases {
    pub issue_policy: IssuePolicy,
    pub get_policy: GetPolicy,
}

pub fn router(use_cases: Arc<UseCases>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/{id}", get(get_by_id))
        .with_state(use_cases)
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
    State(uc): State<Arc<UseCases>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let id = uc
        .issue_policy
        .execute(req.number, req.start_date, req.end_date)
        .await
        .map_err(to_response)?;
    Ok((StatusCode::CREATED, Json(json!({ "id": id }))))
}

async fn get_by_id(
    State(uc): State<Arc<UseCases>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyResponse>, Response> {
    let policy = uc.get_policy.execute(id).await.map_err(to_response)?;
    Ok(Json(PolicyResponse {
        id: policy.id,
        number: policy.number,
        start_date: policy.start_date,
        end_date: policy.end_date,
    }))
}

// El adaptador traduce errores de dominio a HTTP; el dominio no sabe de status codes.
fn to_response(err: PolicyError) -> Response {
    let status = match err {
        PolicyError::InvalidTerm => StatusCode::BAD_REQUEST,
        PolicyError::Duplicated => StatusCode::CONFLICT,
        PolicyError::NotFound => StatusCode::NOT_FOUND,
    };
    message(status, &err.to_string())
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
