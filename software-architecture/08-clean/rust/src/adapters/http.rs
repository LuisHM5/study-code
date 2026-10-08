// Controllers: traducen HTTP -> Input del caso de uso, y delegan la salida al presenter.
use std::sync::Arc;

use axum::{
    extract::{rejection::JsonRejection, Path, State},
    http::StatusCode,
    routing::{get, post},
    Json, Router,
};
use chrono::NaiveDate;
use serde::Deserialize;
use uuid::Uuid;

use super::{
    console_notification_gateway::ConsoleNotificationGateway,
    memory_policy_gateway::MemoryPolicyGateway,
    presenter::{HttpError, IssuedView, PolicyView},
};
use crate::application::{
    get_policy::GetPolicy,
    issue_policy::{IssuePolicy, IssuePolicyInput},
};

pub struct UseCases {
    pub issue_policy: IssuePolicy<MemoryPolicyGateway, ConsoleNotificationGateway>,
    pub get_policy: GetPolicy<MemoryPolicyGateway>,
}

pub fn router(use_cases: Arc<UseCases>) -> Router {
    Router::new()
        .route("/policies", post(issue_policy))
        .route("/policies/{id}", get(get_policy))
        .with_state(use_cases)
}

// DTO de HTTP: propio del adaptador, con sus detalles de serialización.
#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
struct IssuePolicyRequest {
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

impl From<IssuePolicyRequest> for IssuePolicyInput {
    fn from(req: IssuePolicyRequest) -> Self {
        Self {
            number: req.number,
            start_date: req.start_date,
            end_date: req.end_date,
        }
    }
}

async fn issue_policy(
    State(uc): State<Arc<UseCases>>,
    body: Result<Json<IssuePolicyRequest>, JsonRejection>,
) -> Result<(StatusCode, Json<IssuedView>), HttpError> {
    let Json(req) = body.map_err(|_| HttpError::InvalidBody)?;
    let output = uc.issue_policy.execute(req.into()).await?;
    Ok((StatusCode::CREATED, Json(output.into())))
}

async fn get_policy(
    State(uc): State<Arc<UseCases>>,
    Path(id): Path<Uuid>,
) -> Result<Json<PolicyView>, HttpError> {
    let output = uc.get_policy.execute(id).await?;
    Ok(Json(output.into()))
}
