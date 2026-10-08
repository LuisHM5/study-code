//! Onion: anillos concéntricos. Cada anillo solo conoce a los de más adentro:
//! domain <- application <- infrastructure / presentation.

mod application;
mod domain;
mod infrastructure;
mod presentation;

use std::sync::Arc;

use application::{get_policy::GetPolicyService, issue_policy::IssuePolicyService};
use domain::PolicyRepository;
use infrastructure::{
    console_notifier::ConsoleNotifier, memory_repository::InMemoryPolicyRepository,
};

#[tokio::main]
async fn main() {
    let policies: Arc<dyn PolicyRepository> = Arc::new(InMemoryPolicyRepository::default());

    let app = presentation::http::router(Arc::new(presentation::http::Services {
        issue_policy: IssuePolicyService {
            policies: policies.clone(),
            notifier: Arc::new(ConsoleNotifier),
        },
        get_policy: GetPolicyService { policies },
    }));

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
