use chrono::NaiveDate;
use serde::Serialize;
use uuid::Uuid;

// Los hechos que le pueden ocurrir a una póliza. Son inmutables y son la fuente de verdad.
#[derive(Clone, Debug, Serialize)]
#[serde(tag = "type", rename_all_fields = "camelCase")]
pub enum PolicyEvent {
    PolicyIssued {
        number: String,
        start_date: NaiveDate,
        end_date: NaiveDate,
    },
    PolicyRenewed {
        new_end_date: NaiveDate,
    },
}

// Sobre con el que un evento viaja por el bus.
#[derive(Clone, Debug)]
pub struct Published {
    pub policy_id: Uuid,
    pub event: PolicyEvent,
}
