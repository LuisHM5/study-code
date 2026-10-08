use std::{
    collections::HashMap,
    sync::{Arc, Mutex},
};

use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
    Json,
};
use serde_json::json;
use uuid::Uuid;

use crate::policy::domain::Policy;

// Dependencias compartidas. Cada slice toma de aquí solo lo que usa.
#[derive(Clone, Default)]
pub struct AppContext {
    pub policies: Arc<Mutex<HashMap<Uuid, Policy>>>,
}

#[derive(Debug, thiserror::Error)]
pub enum AppError {
    #[error("cuerpo inválido")]
    InvalidBody,
    #[error("vigencia inválida")]
    InvalidTerm,
    #[error("número duplicado")]
    Duplicated,
    #[error("póliza no encontrada")]
    NotFound,
}

impl IntoResponse for AppError {
    fn into_response(self) -> Response {
        let status = match self {
            AppError::InvalidBody | AppError::InvalidTerm => StatusCode::BAD_REQUEST,
            AppError::Duplicated => StatusCode::CONFLICT,
            AppError::NotFound => StatusCode::NOT_FOUND,
        };
        (status, Json(json!({ "message": self.to_string() }))).into_response()
    }
}
