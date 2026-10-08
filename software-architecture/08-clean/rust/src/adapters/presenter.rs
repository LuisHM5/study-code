// Presenter: da forma HTTP a los Output y a los errores de los casos de uso.
use axum::{
    http::StatusCode,
    response::{IntoResponse, Response},
    Json,
};
use chrono::NaiveDate;
use serde::Serialize;
use serde_json::json;
use uuid::Uuid;

use crate::{
    application::{get_policy::GetPolicyOutput, issue_policy::IssuePolicyOutput, AppError},
    domain::DomainError,
};

#[derive(Serialize)]
pub struct IssuedView {
    id: Uuid,
}

impl From<IssuePolicyOutput> for IssuedView {
    fn from(output: IssuePolicyOutput) -> Self {
        Self { id: output.id }
    }
}

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct PolicyView {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

impl From<GetPolicyOutput> for PolicyView {
    fn from(output: GetPolicyOutput) -> Self {
        Self {
            id: output.id,
            number: output.number,
            start_date: output.start_date,
            end_date: output.end_date,
        }
    }
}

pub enum HttpError {
    InvalidBody,
    App(AppError),
}

impl From<AppError> for HttpError {
    fn from(err: AppError) -> Self {
        Self::App(err)
    }
}

impl IntoResponse for HttpError {
    fn into_response(self) -> Response {
        let (status, message) = match self {
            HttpError::InvalidBody => (StatusCode::BAD_REQUEST, "cuerpo inválido".to_owned()),
            HttpError::App(err) => {
                let status = match err {
                    AppError::Domain(DomainError::InvalidTerm) => StatusCode::BAD_REQUEST,
                    AppError::Duplicated => StatusCode::CONFLICT,
                    AppError::NotFound => StatusCode::NOT_FOUND,
                };
                (status, err.to_string())
            }
        };
        (status, Json(json!({ "message": message }))).into_response()
    }
}
