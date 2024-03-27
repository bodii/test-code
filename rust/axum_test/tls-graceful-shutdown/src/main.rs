//! Run with
//! ```not_rust
//! cargo run
//!
//! curl -L -k http://127.0.0.1:7878
//! # or
//! http --follow --verify=no  :7878
//! ```
//!
use std::{future::Future, net::SocketAddr, path::PathBuf, time::Duration};

use axum::{
    extract::Host, handler::HandlerWithoutStateExt, response::Redirect, routing::get, BoxError,
    Router,
};
use axum_server::tls_rustls::RustlsConfig;
use hyper::{StatusCode, Uri};
use tokio::signal;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

#[derive(Debug, Clone, Copy)]
struct Ports {
    http: u16,
    https: u16,
}

#[tokio::main]
async fn main() {
    subscriber();

    let handle = axum_server::Handle::new();
    let shutdown_future = shutdown_signal(handle.clone());

    let ports = Ports {
        http: 7878,
        https: 3000,
    };

    tokio::spawn(redirect_http_to_https(ports, shutdown_future));

    let config = get_certificate_config().await;

    let handler = || async { "Hello, World!" };
    let app = Router::new().route("/", get(handler));

    let addr = SocketAddr::from(([127, 0, 0, 1], ports.https));
    tracing::debug!("listening on {addr}");

    axum_server::bind_rustls(addr, config)
        .handle(handle)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "tls_graceful_shutdown=debug".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

async fn get_certificate_config() -> RustlsConfig {
    let self_signed_certs_path: PathBuf =
        PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("self_signed_certs");
    let cert_pem_path = self_signed_certs_path.join("cert.pem");
    let key_pem_path = self_signed_certs_path.join("key.pem");

    let config = RustlsConfig::from_pem_file(cert_pem_path, key_pem_path)
        .await
        .unwrap();

    config
}

async fn shutdown_signal(handle: axum_server::Handle) {
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }

    tracing::info!("Received termination signal shutting down");
    handle.graceful_shutdown(Some(Duration::from_secs(10)));
}

async fn redirect_http_to_https<F>(ports: Ports, signal: F)
where
    F: Future<Output = ()> + Send + 'static,
{
    fn make_https(host: String, uri: Uri, ports: Ports) -> Result<Uri, BoxError> {
        let mut parts = uri.into_parts();

        parts.scheme = Some(axum::http::uri::Scheme::HTTPS);

        if parts.path_and_query.is_none() {
            parts.path_and_query = Some("/".parse().unwrap());
        }

        let https_host = host.replace(&ports.http.to_string(), &ports.https.to_string());
        parts.authority = Some(https_host.parse()?);

        Ok(Uri::from_parts(parts)?)
    }

    let redirect = move |Host(host): Host, uri: Uri| async move {
        match make_https(host, uri, ports) {
            Ok(uri) => Ok(Redirect::permanent(&uri.to_string())),
            Err(error) => {
                tracing::warn!(%error, "failed to convert URI to HTTPS");
                Err(StatusCode::BAD_REQUEST)
            }
        }
    };

    let addr = SocketAddr::from(([127, 0, 0, 1], ports.http));
    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();
    tracing::debug!("listening on {addr}");
    axum::serve(listener, redirect.into_make_service())
        .with_graceful_shutdown(signal)
        .await
        .unwrap();
}
