use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Clone, Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Policy {
    pub id: Uuid,
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct NewPolicy {
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}
