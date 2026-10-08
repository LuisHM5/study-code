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

use crate::{
    application::{
        get_policy::GetPolicyService,
        issue_policy::{IssuePolicyCommand, IssuePolicyService},
    },
    domain::DomainError,
};

pub struct Services {
    pub issue_policy: IssuePolicyService,
    pub get_policy: GetPolicyService,
}

pub fn router(services: Arc<Services>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/{id}", get(get_by_id))
        .with_state(services)
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
    State(services): State<Arc<Services>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let id = services
        .issue_policy
        .execute(IssuePolicyCommand {
            number: req.number,
            start_date: req.start_date,
            end_date: req.end_date,
        })
        .await
        .map_err(to_response)?;
    Ok((StatusCode::CREATED, Json(json!({ "id": id }))))
}

async fn get_by_id(
    State(services): State<Arc<Services>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyResponse>, Response> {
    let dto = services.get_policy.execute(id).await.map_err(to_response)?;
    Ok(Json(PolicyResponse {
        id: dto.id,
        number: dto.number,
        start_date: dto.start_date,
        end_date: dto.end_date,
    }))
}

fn to_response(err: DomainError) -> Response {
    let status = match err {
        DomainError::EmptyNumber | DomainError::InvalidTerm => StatusCode::BAD_REQUEST,
        DomainError::Duplicated => StatusCode::CONFLICT,
        DomainError::NotFound => StatusCode::NOT_FOUND,
    };
    message(status, &err.to_string())
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
