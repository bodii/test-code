use std::time::Duration;

use axum::{
    body::Bytes,
    extract::{MatchedPath, Request},
    http::HeaderMap,
    response::{Html, Response},
    routing::get,
    Router,
};
use tokio::net::TcpListener;
use tower_http::{classify::ServerErrorsFailureClass, trace::TraceLayer};
use tracing::{info_span, Span};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt, EnvFilter};

#[tokio::main]
async fn main() {
    subscriber();

    let listener = TcpListener::bind("127.0.0.1:3000").await.unwrap();
    let local_addr = listener.local_addr().unwrap();
    tracing::debug!("listening on {}", local_addr);

    axum::serve(listener, app()).await.unwrap();
}

fn subscriber() {
    let filter = EnvFilter::try_from_default_env().unwrap_or_else(|_| {
        "tracing_aka_logging=debug,tower_http=debug,axum::rejection=trace".into()
    });

    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    Router::new().route("/", get(handler)).layer(
        // https://docs.rs/tower-http/latest/tower_http/trace/index.html
        TraceLayer::new_for_http()
            .make_span_with(|request: &Request<_>| {
                let matched_path = request
                    .extensions()
                    .get::<MatchedPath>()
                    .map(MatchedPath::as_str);
                info_span!(
                    "http_request",
                    method = ?request.method(),
                    matched_path,
                    some_other_field = tracing::field::Empty,
                )
            })
            .on_request(|request: &Request<_>, _span: &Span| {
                tracing::debug!("started {} {}", request.method(), request.uri().path())
            })
            .on_response(|_response: &Response, latency: Duration, _span: &Span| {
                tracing::debug!("response generated in {:?}", latency)
            })
            .on_body_chunk(|chunk: &Bytes, _latency: Duration, _span: &Span| {
                tracing::debug!("sending {} bytes", chunk.len())
            })
            .on_eos(
                |_trailers: Option<&HeaderMap>, stream_duration: Duration, _span: &Span| {
                    tracing::debug!("stream closed after {:?}", stream_duration)
                },
            )
            .on_failure(
                |_error: ServerErrorsFailureClass, _latency: Duration, _span: &Span| {
                    tracing::debug!("something went wrong")
                },
            ),
    )
}

async fn handler() -> Html<&'static str> {
    Html("<h1>Hello, World!</h1>")
}
