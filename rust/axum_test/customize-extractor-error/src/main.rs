//! Run with
//!
//! ```not_rust
//! cargo run -p custom-extractor-error
//! ```
use axum::{routing::post, Router};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

mod custom_extractor;
mod derive_from_request;
mod with_rejection;

#[tokio::main]
async fn main() {
    subscriber();

    let app = Router::new()
        .route("/with-rejection", post(with_rejection::handler))
        .route("/custom-extractor", post(custom_extractor::handler))
        .route("/derive-from-request", post(derive_from_request::handler));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap();
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "custom_extractor_error=trace".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}
