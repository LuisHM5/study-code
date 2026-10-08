// Módulo de siniestros. Depende de policies SOLO a través de `PoliciesApi`.
mod app;
mod http;

use std::sync::Arc;

use axum::Router;

use crate::policies::PoliciesApi;

pub fn build(policies: Arc<dyn PoliciesApi>) -> Router {
    http::router(Arc::new(app::Claims::new(policies)))
}
