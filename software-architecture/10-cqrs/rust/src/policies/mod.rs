mod commands;
mod domain;
mod http;
mod infra;
mod queries;

use std::sync::Arc;

use axum::Router;

use commands::issue_policy::IssuePolicyHandler;
use infra::{read_model::ReadModel, write_repo::WriteRepository};

pub fn router() -> Router {
    let read_model = Arc::new(ReadModel::default());

    http::router(Arc::new(http::App {
        issue_policy: IssuePolicyHandler {
            policies: WriteRepository::default(),
            read_model: read_model.clone(),
        },
        read_model,
    }))
}
