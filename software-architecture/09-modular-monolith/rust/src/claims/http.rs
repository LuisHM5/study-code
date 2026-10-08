use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    response::Response,
    routing::{get, post},
    Json, Router,
};
use serde::Deserialize;
use serde_json::{json, Value};
use uuid::Uuid;

use super::app::{Claim, Claims};
use crate::platform::error;

pub fn router(claims: Arc<Claims>) -> Router {
    Router::new()
        .route("/claims", post(report))
        .route("/claims/{id}", get(get_by_id))
        .with_state(claims)
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct ReportRequest {
    policy_id: Uuid,
    description: String,
}

async fn report(
    State(claims): State<Arc<Claims>>,
    body: Result<Json<ReportRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<Value>), Response> {
    let Json(req) = body.map_err(|_| error(StatusCode::BAD_REQUEST, "cuerpo inválido"))?;
    let claim = claims
        .report(req.policy_id, req.description)
        .map_err(|_| error(StatusCode::UNPROCESSABLE_ENTITY, "la póliza no existe"))?;
    Ok((StatusCode::CREATED, Json(json!({ "id": claim.id }))))
}

async fn get_by_id(
    State(claims): State<Arc<Claims>>,
    Path(id): Path<Uuid>,
) -> Result<Json<Claim>, Response> {
    claims
        .get(id)
        .map(Json)
        .ok_or_else(|| error(StatusCode::NOT_FOUND, "siniestro no encontrado"))
}
