use chrono::{Duration, Utc};

use crate::policies::infra::read_model::{PolicyView, ReadModel};

pub struct ListExpiring {
    pub days: i64,
}

pub fn handle(read_model: &ReadModel, query: ListExpiring) -> Vec<PolicyView> {
    let today = Utc::now().date_naive();
    read_model.expiring_between(today, today + Duration::days(query.days))
}
