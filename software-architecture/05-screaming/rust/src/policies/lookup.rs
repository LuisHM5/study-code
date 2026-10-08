use uuid::Uuid;

use super::{Policy, PolicyBook, PolicyError};

/// Consulta de una póliza emitida.
pub fn find_policy(book: &PolicyBook, id: Uuid) -> Result<Policy, PolicyError> {
    book.policies
        .lock()
        .unwrap()
        .get(&id)
        .cloned()
        .ok_or(PolicyError::NotFound)
}
