//! Uses `axum_extra::extract::WithRejection` to transform one rejection into
//! another
//!
//! + Easy learning curve: `WithRejection` acts as a wrapper for another
//!   already existing extractor. You only need to provide a `From` impl
//!   between the original rejection type and the target rejection. Crates like
//!   `thiserror` can provide such conversion using derive macros. See
//!   [`thiserror`]
//! - Verbose types: types become much larger, which makes them difficult to
//!   read. Current limitations on type aliasing makes impossible to destructure
//!   a type alias. See [#1116]
//!   
//! [`thiserror`]: https://crates.io/crates/thiserror
//! [#1116]: https://github.com/tokio-rs/axum/issues/1116#issuecomment-1186197684
//!
use axum::{
    extract::rejection::JsonRejection,
    response::{IntoResponse, Response},
    Json,
};
use axum_extra::extract::WithRejection;
use serde_json::{json, Value};
use thiserror::Error;

pub async fn handler(
    WithRejection(Json(value), _): WithRejection<Json<Value>, ApiError>,
) -> impl IntoResponse {
    Json(dbg!(value))
}

#[derive(Debug, Error)]
pub enum ApiError {
    #[error(transparent)]
    JsonExtractorRejection(#[from] JsonRejection),
}

impl IntoResponse for ApiError {
    fn into_response(self) -> Response {
        let (status, message) = match self {
            ApiError::JsonExtractorRejection(json_rejection) => {
                (json_rejection.status(), json_rejection.body_text())
            }
        };

        let payload = json!({"message": message, "origin": "with_rejection"});

        (status, Json(payload)).into_response()
    }
}
