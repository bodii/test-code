use anyhow::{Context, Result};
use async_session::{MemoryStore, Session, SessionStore};
use axum::{
    extract::{Query, State},
    http::{header::SET_COOKIE, HeaderMap},
    response::{IntoResponse, Redirect},
};
use axum_extra::{headers, TypedHeader};
use oauth2::{
    basic::BasicClient, reqwest::async_http_client, AuthorizationCode, CsrfToken, Scope,
    TokenResponse,
};

use crate::{app_error::AppError, auth_request::AuthRequest, user::User, COOKIE_NAME};
// GET /
pub async fn index(user: Option<User>) -> impl IntoResponse {
    let redirect_url = match user {
        Some(_) => "/protected",
        None => "/auth/discord",
    };

    Redirect::to(redirect_url)
}

// GET /auth/discord
pub async fn discord_auth(State(client): State<BasicClient>) -> impl IntoResponse {
    let (auth_url, _csrf_token) = client
        .authorize_url(CsrfToken::new_random)
        .add_scope(Scope::new("identify".to_string()))
        .url();
    tracing::debug!("{}", auth_url);

    Redirect::to(auth_url.as_ref())
}

// GET /logout
pub async fn logout(
    State(store): State<MemoryStore>,
    TypedHeader(cookies): TypedHeader<headers::Cookie>,
) -> Result<impl IntoResponse, AppError> {
    let cookie = cookies
        .get(COOKIE_NAME)
        .context("unexpected error getting cookie name")?;
    let session = match store
        .load_session(cookie.to_string())
        .await
        .context("failed to load session")?
    {
        Some(s) => s,
        None => return Ok(Redirect::to("/")),
    };

    store
        .destroy_session(session)
        .await
        .context("failed to destroy session")?;

    Ok(Redirect::to("/"))
}

// GET /protected
pub async fn protected(user: User) -> impl IntoResponse {
    format!("Welcome to the protected area :)\nHere's your info:\n{user:?}")
}

pub async fn login_authorized(
    Query(query): Query<AuthRequest>,
    State(store): State<MemoryStore>,
    State(oauth_client): State<BasicClient>,
) -> Result<impl IntoResponse, AppError> {
    let token = oauth_client
        .exchange_code(AuthorizationCode::new(query.code.clone()))
        .request_async(async_http_client)
        .await
        .context("failed in sending request request to authorization server")?;

    let client = reqwest::Client::new();
    let user_data = client
        .get("http://127.0.0.1:14000/info")
        .bearer_auth(token.access_token().secret())
        .send()
        .await
        .context("failed in sending request to target Url")?
        .json::<User>()
        .await
        .context("failed to deserialize response as JSON")?;

    let mut session = Session::new();
    session
        .insert("user", &user_data)
        .context("failed in inserting serialized value into session")?;

    let cookie = store
        .store_session(session)
        .await
        .context("failed to store session")?
        .context("unexpected error retrieving cookie value")?;

    let cookie = format!("{}={}; SameSite=Lax; Path=/", COOKIE_NAME, cookie);

    let mut headers = HeaderMap::new();
    headers.insert(
        SET_COOKIE,
        cookie.parse().context("failed to parse cookie value")?,
    );

    Ok((headers, Redirect::to("/")))
}
