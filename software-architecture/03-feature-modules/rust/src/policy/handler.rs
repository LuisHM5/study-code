use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    Json,
};
use serde_json::{json, Value};
use uuid::Uuid;

use super::{
    model::{NewPolicy, Policy},
    service::PolicyService,
};
use crate::platform::error::AppError;

pub async fn issue(
    State(service): State<Arc<PolicyService>>,
    body: Result<Json<NewPolicy>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), AppError> {
    let Json(input) = body.map_err(|_| AppError::InvalidBody)?;
    let policy = service.issue(input)?;
    Ok((StatusCode::CREATED, Json(json!({ "id": policy.id }))))
}

pub async fn get_by_id(
    State(service): State<Arc<PolicyService>>,
    Path(id): Path<Uuid>,
) -> Result<Json<Policy>, AppError> {
    Ok(Json(service.get(id)?))
}
