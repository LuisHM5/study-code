use uuid::Uuid;

use super::{Policy, PolicyBook, PolicyError, PolicyNumber, PolicyTerm};
use crate::platform::mail;

/// Emisión de una póliza.
pub fn issue_policy(
    book: &PolicyBook,
    number: PolicyNumber,
    term: PolicyTerm,
) -> Result<Policy, PolicyError> {
    let mut policies = book.policies.lock().unwrap();
    if policies.values().any(|p| p.number == number) {
        return Err(PolicyError::Duplicated);
    }

    let policy = Policy {
        id: Uuid::new_v4(),
        number,
        term,
    };
    policies.insert(policy.id, policy.clone());

    mail::notify_policy_issued(policy.number.as_str());
    Ok(policy)
}
