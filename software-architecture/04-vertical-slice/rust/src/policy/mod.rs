pub mod domain;
mod get_policy;
mod issue_policy;

use axum::{
    routing::{get, post},
    Router,
};

use crate::app::AppContext;

// Añadir un caso de uso = añadir un archivo y una línea aquí.
pub fn router() -> Router<AppContext> {
    Router::new()
        .route("/policies", post(issue_policy::handler))
        .route("/policies/{id}", get(get_policy::handler))
}
