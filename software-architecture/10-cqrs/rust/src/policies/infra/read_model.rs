use std::sync::RwLock;

use chrono::NaiveDate;
use serde::Serialize;
use uuid::Uuid;

use crate::policies::domain::Policy;

// Vista de lectura: la forma exacta que necesita la pantalla, lista para serializar.
#[derive(Clone, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct PolicyView {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

// Almacén de lectura, separado del de escritura (equivale a una tabla desnormalizada).
#[derive(Default)]
pub struct ReadModel {
    rows: RwLock<Vec<PolicyView>>,
}

impl ReadModel {
    pub fn project(&self, policy: &Policy) {
        self.rows.write().unwrap().push(PolicyView {
            id: policy.id,
            number: policy.number.clone(),
            start_date: policy.start_date,
            end_date: policy.end_date,
        });
    }

    pub fn by_id(&self, id: Uuid) -> Option<PolicyView> {
        let rows = self.rows.read().unwrap();
        rows.iter().find(|row| row.id == id).cloned()
    }

    pub fn expiring_between(&self, from: NaiveDate, to: NaiveDate) -> Vec<PolicyView> {
        let rows = self.rows.read().unwrap();
        let mut expiring: Vec<PolicyView> = rows
            .iter()
            .filter(|row| row.end_date >= from && row.end_date <= to)
            .cloned()
            .collect();
        expiring.sort_by_key(|row| row.end_date);
        expiring
    }
}
