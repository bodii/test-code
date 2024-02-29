//! Uses `axum::extract::FromRequest` to wrap another extractor and customize the
//! rejection
//!
//! + Easy learning curve: Deriving `FromRequest` generates a `FromRequest`
//!   implementation for your type using another extractor. You only need
//!   to provide a `From` impl between the original rejection type and the
//!   target rejection. Crates like [`thiserror`] can provide such conversion
//!   using derive macros.
//! - Boilerplate: Requires deriving `FromRequest` for every custom rejection
//! - There are some known limitations: [FromRequest#known-limitations]
//!
//! [`thiserror`]: https://crates.io/crates/thiserror
//! [FromRequest#known-limitations]: https://docs.rs/axum-macros/*/axum_macros/derive.FromRequest.html#known-limitations

use axum::{
    extract::{rejection::JsonRejection, FromRequest},
    http::StatusCode,
    response::{IntoResponse, Response},
};
use serde::Serialize;
use serde_json::{json, Value};

pub async fn handler(Json(value): Json<Value>) -> impl IntoResponse {
    Json(dbg!(value))
}

#[derive(Debug)]
pub struct ApiError {
    status: StatusCode,
    message: String,
}

#[derive(FromRequest)]
#[from_request(via(axum::Json), rejection(ApiError))]
pub struct Json<T>(T);

impl<T: Serialize> IntoResponse for Json<T> {
    fn into_response(self) -> Response {
        let Self(value) = self;
        axum::Json(value).into_response()
    }
}

impl From<JsonRejection> for ApiError {
    fn from(rejection: JsonRejection) -> Self {
        Self {
            status: rejection.status(),
            message: rejection.body_text(),
        }
    }
}

impl IntoResponse for ApiError {
    fn into_response(self) -> Response {
        let payload = json!({
        "message": self.message,
        "origin": "derive_from_request"
        });

        (self.status, axum::Json(payload)).into_response()
    }
}
