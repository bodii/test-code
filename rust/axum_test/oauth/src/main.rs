//! ```not_rust
//! # start oauth server
//! cd oauth-server
//! go run server.go
//!
//! # start reqeust api
//! CLIENT_ID=000000 CLIENT_SECRET=999999 cargo run
//!
//! # visit
//! request in the browser: http://127.0.0.1:3000
//! ```
//!
mod api;
mod app_error;
mod app_state;
mod auth_redirect;
mod auth_request;
mod user;

use std::env;

use anyhow::{Context, Result};
use api::{discord_auth, index, login_authorized, logout, protected};
use app_state::AppState;
use async_session::MemoryStore;
use axum::{routing::get, Router};
use oauth2::{basic::BasicClient, AuthUrl, ClientId, ClientSecret, RedirectUrl, TokenUrl};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

static COOKIE_NAME: &str = "SESSION";

#[tokio::main]
async fn main() {
    subscriber();

    let store = MemoryStore::new();
    let oauth_client = oauth_client().unwrap();

    let app_state = AppState {
        store,
        oauth_client,
    };

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .context("failed to build TcpListener")
        .unwrap();

    tracing::debug!(
        "listening on http://{}",
        listener
            .local_addr()
            .context("failed to return local address")
            .unwrap()
    );

    axum::serve(listener, app(app_state)).await.unwrap();
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "oauth=debug".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app(state: AppState) -> Router {
    Router::new()
        .route("/", get(index))
        .route("/auth/discord", get(discord_auth))
        .route("/auth/authorized", get(login_authorized))
        .route("/protected", get(protected))
        .route("/logout", get(logout))
        .with_state(state)
}

fn oauth_client() -> Result<BasicClient, app_error::AppError> {
    let client_id = env::var("CLIENT_ID").context("Missing CLIENT_ID!")?;
    // let client_id = env::var("CLIENT_ID").unwrap_or_else(|_| "client_id".to_string());

    let client_secret = env::var("CLIENT_SECRET").context("Missing CLIENT_SECRET")?;
    // let client_secret = env::var("CLIENT_SECRET").unwrap_or_else(|_| "client_secret".to_string());

    let redirect_url = env::var("REDIRECT_URL")
        .unwrap_or_else(|_| "http://127.0.0.1:3000/auth/authorized".to_string());
    // see https://github.com/RangelReale/osin
    // or see https://pkg.go.dev/github.com/RangelReale/osin
    let auth_url =
        env::var("AUTH_URL").unwrap_or_else(|_| "http://127.0.0.1:14000/authorize".to_string());

    let token_url =
        env::var("TOKEN_URL").unwrap_or_else(|_| "http://127.0.0.1:14000/token".to_string());

    Ok(BasicClient::new(
        ClientId::new(client_id),
        Some(ClientSecret::new(client_secret)),
        AuthUrl::new(auth_url).context("failed to create new authorization server URL")?,
        Some(TokenUrl::new(token_url).context("failed to create new token endpoint URL")?),
    )
    .set_redirect_uri(
        RedirectUrl::new(redirect_url).context("failed to create new redirection URL")?,
    ))
}
