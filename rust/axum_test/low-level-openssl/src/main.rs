use std::{path::PathBuf, pin::Pin};

use axum::{http::Request, routing::get, Router};
use hyper::body::Incoming;
use hyper_util::rt::{TokioExecutor, TokioIo};
use openssl::ssl::{Ssl, SslAcceptor, SslFiletype, SslMethod};
use tokio::net::TcpListener;
use tokio_openssl::SslStream;
use tower::Service;
use tracing::{error, info, warn};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

#[tokio::main]
async fn main() {
    subscriber();

    let signed_certs_path = PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("self_signed_certs");

    let mut tls_builder = SslAcceptor::mozilla_modern_v5(SslMethod::tls()).unwrap();
    tls_builder
        .set_certificate_file(signed_certs_path.join("cert.pem"), SslFiletype::PEM)
        .unwrap();

    tls_builder
        .set_private_key_file(signed_certs_path.join("key.pem"), SslFiletype::PEM)
        .unwrap();
    tls_builder.check_private_key().unwrap();

    let tls_acceptor = tls_builder.build();

    let bind = "[::1]:3000";
    let tcp_listener = TcpListener::bind(bind).await.unwrap();
    info!("HTTPS server listening on {bind}. To contact curl -k https://localhost:3000");

    let app = app();

    loop {
        let tower_service = app.clone();
        let tls_acceptor = tls_acceptor.clone();

        let (cnx, addr) = tcp_listener.accept().await.unwrap();

        tokio::spawn(async move {
            let ssl = Ssl::new(tls_acceptor.context()).unwrap();
            let mut tls_stream = SslStream::new(ssl, cnx).unwrap();
            if let Err(err) = SslStream::accept(Pin::new(&mut tls_stream)).await {
                error!(
                    "error during tls handshake connection from {}; {}",
                    addr, err
                );
                return;
            }

            let stream = TokioIo::new(tls_stream);

            let hyper_serivce = hyper::service::service_fn(move |request: Request<Incoming>| {
                tower_service.clone().call(request)
            });

            let ret = hyper_util::server::conn::auto::Builder::new(TokioExecutor::new())
                .serve_connection_with_upgrades(stream, hyper_serivce)
                .await;

            if let Err(err) = ret {
                warn!("error serving connection from {}: {}", addr, err);
            }
        });
    }
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "low_level_openssl=debug".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    Router::new().route("/", get(handler))
}

async fn handler() -> &'static str {
    "Hello, World"
}
