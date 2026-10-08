//! Hexagonal: el núcleo (domain + casos de uso) define puertos (traits);
//! los adaptadores los implementan. Aquí se conecta cada puerto con su adaptador.

mod policy;

use std::sync::Arc;

use policy::{
    adapters::{console::ConsoleNotifier, http, memory::InMemoryPolicyRepository},
    get_policy::GetPolicy,
    issue_policy::IssuePolicy,
    ports::{Notifier, PolicyRepository},
};

#[tokio::main]
async fn main() {
    // Adaptadores de salida. Cambiar de proveedor es cambiar estas dos líneas.
    let policies: Arc<dyn PolicyRepository> = Arc::new(InMemoryPolicyRepository::default());
    let notifier: Arc<dyn Notifier> = Arc::new(ConsoleNotifier);

    // Adaptador de entrada: HTTP.
    let app = http::router(Arc::new(http::UseCases {
        issue_policy: IssuePolicy {
            policies: policies.clone(),
            notifier,
        },
        get_policy: GetPolicy { policies },
    }));

    let port = std::env::var("PORT").unwrap_or_else(|_| "3000".into());
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{port}"))
        .await
        .unwrap();
    println!("escuchando en http://localhost:{port}");
    axum::serve(listener, app).await.unwrap();
}
