use async_session::{MemoryStore, SessionStore};
use axum::{
    async_trait,
    extract::{FromRef, FromRequestParts},
    RequestPartsExt,
};
use axum_extra::{headers, typed_header::TypedHeaderRejectionReason, TypedHeader};
use serde::{Deserialize, Serialize};

use http::{header, request::Parts};

use crate::{auth_redirect::AuthRedirect, COOKIE_NAME};

#[derive(Debug, Serialize, Deserialize)]
pub struct User {
    pub id: String,
    pub avatar: Option<String>,
    pub username: String,
    pub discriminator: String,
}

#[async_trait]
impl<S> FromRequestParts<S> for User
where
    MemoryStore: FromRef<S>,
    S: Send + Sync,
{
    type Rejection = AuthRedirect;

    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        let store = MemoryStore::from_ref(state);

        let cookies = parts
            .extract::<TypedHeader<headers::Cookie>>()
            .await
            .map_err(|e| match *e.name() {
                header::COOKIE => match e.reason() {
                    TypedHeaderRejectionReason::Missing => AuthRedirect,
                    _ => panic!("Unexpected error getting cookies: {e}"),
                },
                _ => panic!("Unexpected error getting cookies: {e}"),
            })?;

        let session_cookie = cookies.get(COOKIE_NAME).ok_or(AuthRedirect)?;

        let session = store
            .load_session(session_cookie.to_string())
            .await
            .unwrap()
            .ok_or(AuthRedirect)?;

        let user = session.get::<User>("user").ok_or(AuthRedirect)?;

        Ok(user)
    }
}

