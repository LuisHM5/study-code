use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, Query, State},
    http::StatusCode,
    response::{IntoResponse, Response},
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::Deserialize;
use serde_json::{json, Value};
use uuid::Uuid;

use super::{
    commands::issue_policy::{IssuePolicy, IssuePolicyHandler},
    domain::PolicyError,
    infra::read_model::{PolicyView, ReadModel},
    queries::{
        get_policy::{self, GetPolicy},
        list_expiring::{self, ListExpiring},
    },
};

pub struct App {
    pub issue_policy: IssuePolicyHandler,
    pub read_model: Arc<ReadModel>,
}

pub fn router(app: Arc<App>) -> Router {
    Router::new()
        .route("/policies", post(issue))
        .route("/policies/expiring", get(expiring))
        .route("/policies/{id}", get(get_by_id))
        .with_state(app)
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssueRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Deserialize)]
struct ExpiringParams {
    days: Option<i64>,
}

// Comando
async fn issue(
    State(app): State<Arc<App>>,
    body: Result<Json<IssueRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let id = app
        .issue_policy
        .handle(IssuePolicy {
            number: req.number,
            start_date: req.start_date,
            end_date: req.end_date,
        })
        .map_err(|err| {
            let status = match err {
                PolicyError::InvalidTerm => StatusCode::BAD_REQUEST,
                PolicyError::Duplicated => StatusCode::CONFLICT,
            };
            message(status, &err.to_string())
        })?;
    Ok((StatusCode::CREATED, Json(json!({ "id": id }))))
}

// Consultas
async fn get_by_id(
    State(app): State<Arc<App>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyView>, Response> {
    get_policy::handle(&app.read_model, GetPolicy { id })
        .map(Json)
        .ok_or_else(|| message(StatusCode::NOT_FOUND, "póliza no encontrada"))
}

async fn expiring(
    State(app): State<Arc<App>>,
    Query(params): Query<ExpiringParams>,
) -> Json<Vec<PolicyView>> {
    let days = params.days.unwrap_or(30);
    Json(list_expiring::handle(
        &app.read_model,
        ListExpiring { days },
    ))
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
