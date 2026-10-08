use std::time::Duration;

use serde_json::json;

use crate::policy::Policy;

// Cliente del otro servicio. Lo que antes era una llamada a función ahora es una
// llamada de red: puede tardar o fallar, así que siempre lleva timeout.
pub struct NotificationsClient {
    http: reqwest::Client,
    base_url: String,
}

impl NotificationsClient {
    pub fn new(base_url: String) -> Self {
        let http = reqwest::Client::builder()
            .timeout(Duration::from_secs(2))
            .build()
            .expect("cliente HTTP");
        Self { http, base_url }
    }

    pub async fn policy_issued(&self, policy: &Policy) {
        let result = self
            .http
            .post(format!("{}/notifications", self.base_url))
            .json(&json!({ "policyId": policy.id, "number": policy.number }))
            .send()
            .await
            .and_then(|response| response.error_for_status());

        // La póliza ya está emitida: un fallo al notificar solo se registra.
        if let Err(err) = result {
            eprintln!(
                "[policies] no se pudo notificar la póliza {}: {err}",
                policy.number
            );
        }
    }
}
