//! Manual implementation of `FromRequest` that wraps another extractor
//!
//! + Powerful API: Implementing `FromRequest` grants access to `RequestParts`
//!   and `async/await`. This means that you can create more powerful rejections
//! - Boilerplate: Requires creating a new extractor for every custom rejection
//! - Complexity: Manually implementing `FromRequest` results on more complex code

use axum::{
    async_trait,
    extract::{rejection::JsonRejection, FromRequest, MatchedPath, Request},
    http::StatusCode,
    response::IntoResponse,
    RequestPartsExt,
};
use serde_json::{json, Value};

pub async fn handler(Json(value): Json<Value>) -> impl IntoResponse {
    Json(dbg!(value));
}

pub struct Json<T>(pub T);

#[async_trait]
impl<S, T> FromRequest<S> for Json<T>
where
    axum::Json<T>: FromRequest<S, Rejection = JsonRejection>,
    S: Send + Sync,
{
    type Rejection = (StatusCode, axum::Json<Value>);

    async fn from_request(req: Request, state: &S) -> Result<Self, Self::Rejection> {
        let (mut parts, body) = req.into_parts();

        let path = parts
            .extract::<MatchedPath>()
            .await
            .map(|path| path.as_str().to_owned())
            .ok();

        let req = Request::from_parts(parts, body);

        match axum::Json::<T>::from_request(req, state).await {
            Ok(value) => Ok(Self(value.0)),
            Err(rejection) => {
                let payload = json!({
                               "message": rejection.body_text(),
                               "origin": "custom_extractor",
                               "path": path,
                });

                Err((rejection.status(), axum::Json(payload)))
            }
        }
    }
}
