use core::fmt;
use std::str::FromStr;

use axum::{extract::Query, routing::get, Router};
use serde::{de, Deserialize, Deserializer};
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
    let filter = EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "query_params_with_empty_strings=debug".into());

    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    Router::new().route("/", get(handler))
}

#[derive(Debug, Deserialize)]
#[allow(dead_code)]
struct Params {
    #[serde(default, deserialize_with = "empty_string_as_none")]
    foo: Option<i32>,
    bar: Option<String>,
}

fn empty_string_as_none<'de, D, T>(de: D) -> Result<Option<T>, D::Error>
where
    D: Deserializer<'de>,
    T: FromStr,
    T::Err: fmt::Display,
{
    let opt = Option::<String>::deserialize(de)?;
    match opt.as_deref() {
        None | Some("") => Ok(None),
        Some(s) => FromStr::from_str(s).map_err(de::Error::custom).map(Some),
    }
}

async fn handler(Query(params): Query<Params>) -> String {
    format!("{params:?}")
}

#[cfg(test)]
mod tests {
    use super::*;
    use axum::{body::Body, http::Request};
    use http_body_util::BodyExt;
    use tower::ServiceExt;

    #[tokio::test]
    async fn test_something() {
        assert_eq!(
            send_request_get_body("foo=1&&bar=bar").await,
            r#"Params { foo: Some(1), bar: Some("bar") }"#
        );

        assert_eq!(
            send_request_get_body("foo=1&&bar=bar").await,
            r#"Params { foo: Some(1), bar: Some("bar") }"#
        );

        assert_eq!(
            send_request_get_body("foo=&&bar=").await,
            r#"Params { foo: None, bar: Some("") }"#
        );

        assert_eq!(
            send_request_get_body("foo=1").await,
            r#"Params { foo: Some(1), bar: None }"#
        );

        assert_eq!(
            send_request_get_body("bar=bar").await,
            r#"Params { foo: None, bar: Some("bar") }"#
        );

        assert_eq!(
            send_request_get_body("foo=").await,
            r#"Params { foo: None, bar: None }"#
        );

        assert_eq!(
            send_request_get_body("bar=").await,
            r#"Params { foo: None, bar: Some("") }"#
        );

        assert_eq!(
            send_request_get_body("").await,
            r#"Params { foo: None, bar: None }"#
        );
    }

    async fn send_request_get_body(query: &str) -> String {
        let body = app()
            .oneshot(
                Request::builder()
                    .uri(format!("/?{query}"))
                    .body(Body::empty())
                    .unwrap(),
            )
            .await
            .unwrap()
            .into_body();

        let bytes = body.collect().await.unwrap().to_bytes();
        String::from_utf8(bytes.to_vec()).unwrap()
    }
}
