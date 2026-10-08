use axum::{
    extract::{rejection::JsonRejection, State},
    http::StatusCode,
    Json,
};
use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

use super::domain::{validate_term, Policy};
use crate::app::{AppContext, AppError};

// Request y Response son privados del slice: cambiarlos no afecta a los demás.
#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Request {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

#[derive(Serialize)]
pub struct Response {
    id: Uuid,
}

pub async fn handler(
    State(ctx): State<AppContext>,
    body: Result<Json<Request>, JsonRejection>,
) -> Result<(StatusCode, Json<Response>), AppError> {
    let Json(req) = body.map_err(|_| AppError::InvalidBody)?;
    validate_term(req.start_date, req.end_date)?;

    let mut policies = ctx.policies.lock().unwrap();
    if policies.values().any(|p| p.number == req.number) {
        return Err(AppError::Duplicated);
    }

    let policy = Policy {
        id: Uuid::new_v4(),
        number: req.number,
        start_date: req.start_date,
        end_date: req.end_date,
    };
    println!("[notificación] póliza emitida {}", policy.number);
    let id = policy.id;
    policies.insert(id, policy);

    Ok((StatusCode::CREATED, Json(Response { id })))
}
