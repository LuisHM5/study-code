use uuid::Uuid;

use crate::policies::infra::read_model::{PolicyView, ReadModel};

pub struct GetPolicy {
    pub id: Uuid,
}

pub fn handle(read_model: &ReadModel, query: GetPolicy) -> Option<PolicyView> {
    read_model.by_id(query.id)
}
