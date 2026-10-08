// Utilidades técnicas compartidas por todos los módulos.
use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
    Json,
};
use serde_json::json;

pub fn error(status: StatusCode, message: &str) -> Response {
    (status, Json(json!({ "message": message }))).into_response()
}
