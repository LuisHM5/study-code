//! Clean Architecture: 4 círculos y una regla, las dependencias apuntan hacia adentro.
//!   1. domain       -> entidades
//!   2. application  -> casos de uso + gateways (traits)
//!   3. adapters     -> controllers, presenters, implementaciones de gateways
//!   4. main.rs      -> framework y cableado: el único lugar que conoce todo

mod adapters;
mod application;
mod domain;

use std::sync::Arc;

use adapters::{
    console_notification_gateway::ConsoleNotificationGateway,
    http::{self, UseCases},
    memory_policy_gateway::MemoryPolicyGateway,
};
use application::{get_policy::GetPolicy, issue_policy::IssuePolicy};

#[tokio::main]
async fn main() {
    let policies = MemoryPolicyGateway::default();

    let app = http::router(Arc::new(UseCases {
        issue_policy: IssuePolicy {
            policies: policies.clone(),
            notifier: ConsoleNotificationGateway,
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
