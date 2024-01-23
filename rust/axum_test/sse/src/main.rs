//! Run with
//!
//! ```not_rust
//! cargo run -p sse`
//! ```
//! Test with
//! ```not_rust
//! cargo test -p sse
//! ```

use std::{convert::Infallible, path::PathBuf, time::Duration};

use axum::{
    response::{sse::Event, Sse},
    routing::get,
    Router,
};
use axum_extra::{headers, TypedHeader};
use futures::{stream, Stream};
use tokio_stream::StreamExt as _;
use tower_http::{services::ServeDir, trace::TraceLayer};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt, EnvFilter};

#[tokio::main]
async fn main() {
    subscriber();

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app()).await.unwrap();
}

fn subscriber() {
    let filter =
        EnvFilter::try_from_default_env().unwrap_or_else(|_| "sse=debug,tower_http=debug".into());
    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    let assets_dir = PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("assets");
    let static_files_service = ServeDir::new(assets_dir).append_index_html_on_directories(true);
    Router::new()
        .fallback_service(static_files_service)
        .route("/sse", get(sse_handler))
        .layer(TraceLayer::new_for_http())
}

async fn sse_handler(
    TypedHeader(user_agent): TypedHeader<headers::UserAgent>,
) -> Sse<impl Stream<Item = Result<Event, Infallible>>> {
    println!("`{}` connected", user_agent.as_str());

    let stream = stream::repeat_with(|| Event::default().data("hi!"))
        .map(Ok)
        .throttle(Duration::from_secs(1));

    Sse::new(stream).keep_alive(
        axum::response::sse::KeepAlive::new()
            .interval(Duration::from_secs(1))
            .text("keep-alive-text"),
    )
}

#[cfg(test)]
mod tests {
    use eventsource_stream::Eventsource;
    use tokio::net::TcpListener;

    use super::*;

    #[tokio::test]
    async fn integration_test() {
        async fn spawn_app(host: impl Into<String>) -> String {
            let host = host.into();
            let listener = TcpListener::bind(format!("{}:0", host)).await.unwrap();
            let port = listener.local_addr().unwrap().port();
            tokio::spawn(async {
                axum::serve(listener, app()).await.unwrap();
            });

            format!("http://{}:{}", host, port)
        }

        let listening_url = spawn_app("127.0.0.1").await;

        let mut event_stream = reqwest::Client::new()
            .get(&format!("{}/sse", listening_url))
            .header("User-Agent", "integration_test")
            .send()
            .await
            .unwrap()
            .bytes_stream()
            .eventsource()
            .take(1);

        let mut event_data: Vec<String> = vec![];

        while let Some(event) = event_stream.next().await {
            match event {
                Ok(event) => {
                    if event.data == "[DONE]" {
                        break;
                    }

                    event_data.push(event.data);
                }
                Err(_) => {
                    println!("Error in event stream");
                }
            }
        }

        assert!(event_data[0] == "hi!");
    }
}
