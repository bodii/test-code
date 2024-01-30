use std::net::SocketAddr;

use axum::{
    extract::ConnectInfo,
    routing::{get, post},
    Json, Router,
};
use tower_http::trace::TraceLayer;
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
        .unwrap_or_else(|_| "testing=debug,tower_http=debug".into());

    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn app() -> Router {
    Router::new()
        .route("/", get(|| async { "Hello, World!" }))
        .route(
            "/json",
            post(|payload: Json<serde_json::Value>| async move {
                Json(serde_json::json!({ "data": payload.0 }))
            }),
        )
        .route(
            "/requires-connect-into",
            get(|ConnectInfo(addr): ConnectInfo<SocketAddr>| async move { format!("Hi {addr}") }),
        )
        .layer(TraceLayer::new_for_http())
}

#[cfg(test)]
mod tests {
    use axum::{
        body::Body,
        extract::{connect_info::MockConnectInfo, Request},
        http::{self, StatusCode},
    };
    use http_body_util::BodyExt;
    use serde_json::{json, Value};
    use tokio::net::TcpListener;
    use tower::{Service, ServiceExt};

    use super::*;

    #[tokio::test]
    async fn hello_world() {
        let app = app();

        let response = app
            .oneshot(Request::builder().uri("/").body(Body::empty()).unwrap())
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::OK);

        let body = response.into_body().collect().await.unwrap().to_bytes();
        assert_eq!(&body[..], b"Hello, World!");
    }

    #[tokio::test]
    async fn json() {
        let app = app();

        let response = app
            .oneshot(
                Request::builder()
                    .method(http::Method::POST)
                    .uri("/json")
                    .header(http::header::CONTENT_TYPE, mime::APPLICATION_JSON.as_ref())
                    .body(Body::from(
                        serde_json::to_vec(&json!([1, 2, 3, 4])).unwrap(),
                    ))
                    .unwrap(),
            )
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::OK);

        let body = response.into_body().collect().await.unwrap().to_bytes();
        let body: Value = serde_json::from_slice(&body).unwrap();
        assert_eq!(body, json!({ "data": [1, 2, 3, 4] }));
    }

    #[tokio::test]
    async fn not_found() {
        let app = app();

        let response = app
            .oneshot(
                Request::builder()
                    .uri("/does-not-exist")
                    .body(Body::empty())
                    .unwrap(),
            )
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::NOT_FOUND);

        let body = response.into_body().collect().await.unwrap().to_bytes();
        assert!(body.is_empty());
    }

    #[tokio::test]
    async fn the_real_deal() {
        let listener = TcpListener::bind("0.0.0.0:0").await.unwrap();
        let addr = listener.local_addr().unwrap();

        tokio::spawn(async move {
            axum::serve(listener, app()).await.unwrap();
        });

        let client =
            hyper_util::client::legacy::Client::builder(hyper_util::rt::TokioExecutor::new())
                .build_http();

        let response = client
            .request(
                Request::builder()
                    .uri(format!("http://{addr}"))
                    .header("Host", "localhost")
                    .body(Body::empty())
                    .unwrap(),
            )
            .await
            .unwrap();

        let body = response.into_body().collect().await.unwrap().to_bytes();
        assert_eq!(&body[..], b"Hello, World!");
    }

    #[tokio::test]
    async fn multiple_request() {
        let mut app = app().into_service();

        let request = Request::builder().uri("/").body(Body::empty()).unwrap();
        let response = ServiceExt::<Request<Body>>::ready(&mut app)
            .await
            .unwrap()
            .call(request)
            .await
            .unwrap();
        assert_eq!(response.status(), StatusCode::OK);
    }

    #[tokio::test]
    async fn with_into_make_service_with_connect_info() {
        let mut app = app()
            .layer(MockConnectInfo(SocketAddr::from(([0, 0, 0, 0], 3000))))
            .into_service();

        let request = Request::builder()
            .uri("/requires-connect-into")
            .body(Body::empty())
            .unwrap();

        let response = app.ready().await.unwrap().call(request).await.unwrap();

        assert_eq!(response.status(), StatusCode::OK);
    }
}
