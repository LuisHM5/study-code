use std::sync::Arc;

use chrono::NaiveDate;
use uuid::Uuid;

use crate::policies::{
    domain::{Policy, PolicyError},
    infra::{read_model::ReadModel, write_repo::WriteRepository},
};

pub struct IssuePolicy {
    pub number: String,
    pub start_date: NaiveDate,
    pub end_date: NaiveDate,
}

pub struct IssuePolicyHandler {
    pub policies: WriteRepository,
    pub read_model: Arc<ReadModel>,
}

impl IssuePolicyHandler {
    // Un comando no devuelve datos de lectura: solo el id.
    pub fn handle(&self, cmd: IssuePolicy) -> Result<Uuid, PolicyError> {
        if self.policies.exists_by_number(&cmd.number) {
            return Err(PolicyError::Duplicated);
        }
        let policy = Policy::issue(cmd.number, cmd.start_date, cmd.end_date)?;
        self.policies.save(&policy);

        // Proyección: actualiza el modelo de lectura. Aquí es síncrona;
        // con almacenes separados sería asíncrona, a partir de un evento.
        self.read_model.project(&policy);

        println!("[notificación] póliza emitida {}", policy.number);
        Ok(policy.id)
    }
}
