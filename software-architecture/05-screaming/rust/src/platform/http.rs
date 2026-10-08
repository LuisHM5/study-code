use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    response::{IntoResponse, Response},
    Json,
};
use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};
use uuid::Uuid;

use crate::policies::{
    issuing::issue_policy as issue, lookup::find_policy as find, PolicyBook, PolicyError,
    PolicyNumber, PolicyTerm,
};

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct IssuePolicyRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct PolicyView {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

pub async fn issue_policy(
    State(book): State<Arc<PolicyBook>>,
    body: Result<Json<IssuePolicyRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| message(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;

    let number = PolicyNumber::parse(&req.number).map_err(to_response)?;
    let term = PolicyTerm::new(req.start_date, req.end_date).map_err(to_response)?;
    let policy = issue(&book, number, term).map_err(to_response)?;

    Ok((StatusCode::CREATED, Json(json!({ "id": policy.id }))))
}

pub async fn find_policy(
    State(book): State<Arc<PolicyBook>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyView>, Response> {
    let policy = find(&book, id).map_err(to_response)?;
    Ok(Json(PolicyView {
        id: policy.id,
        number: policy.number.as_str().to_owned(),
        start_date: policy.term.start(),
        end_date: policy.term.end(),
    }))
}

fn to_response(err: PolicyError) -> Response {
    let status = match err {
        PolicyError::EmptyNumber | PolicyError::InvalidTerm => StatusCode::BAD_REQUEST,
        PolicyError::Duplicated => StatusCode::CONFLICT,
        PolicyError::NotFound => StatusCode::NOT_FOUND,
    };
    message(status, &err.to_string())
}

fn message(status: StatusCode, text: &str) -> Response {
    (status, Json(json!({ "message": text }))).into_response()
}
