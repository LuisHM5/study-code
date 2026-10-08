use axum::{
    extract::{Path, State},
    Json,
};
use chrono::NaiveDate;
use serde::Serialize;
use uuid::Uuid;

use crate::app::{AppContext, AppError};

#[derive(Serialize)]
#[serde(rename_all = "camelCase")]
pub struct Response {
    id: Uuid,
    number: String,
    start_date: NaiveDate,
    end_date: NaiveDate,
}

// Un slice de lectura simple va directo a los datos, sin capas intermedias.
pub async fn handler(
    State(ctx): State<AppContext>,
    Path(id): Path<Uuid>,
) -> Result<Json<Response>, AppError> {
    let policies = ctx.policies.lock().unwrap();
    let policy = policies.get(&id).ok_or(AppError::NotFound)?;
    Ok(Json(Response {
        id: policy.id,
        number: policy.number.clone(),
        start_date: policy.start_date,
        end_date: policy.end_date,
    }))
}
