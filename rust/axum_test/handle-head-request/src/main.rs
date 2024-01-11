use axum::response::{IntoResponse, Response};
use axum::{http::Method, routing::get, Router};

fn app() -> Router {
    Router::new().route("/get-head", get(get_head_handler))
}

// http head :3000/get-head
// http get :3000/get-head
async fn get_head_handler(method: Method) -> Response {
    if method == Method::HEAD {
        return ([("x-some-header", "header from HEAD")]).into_response();
    }

    do_some_computing_task();

    ([("x-some-header", "header from GET")], "body from GET").into_response()
}

fn do_some_computing_task() {
    // TODO
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();
    let info = format!("listening on {}", listener.local_addr().unwrap());
    tracing::info!(info);
    axum::serve(listener, app()).await.unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;
    use axum::body::Body;
    use axum::http::{Request, StatusCode};
    use http_body_util::BodyExt;
    use tower::ServiceExt;

    #[tokio::test]
    async fn test_get() {
        let app = app();

        let response = app
            .oneshot(Request::get("/get-head").body(Body::empty()).unwrap())
            .await
            .unwrap();

        assert_eq!(response.status(), StatusCode::OK);
        assert_eq!(response.headers()["x-some-header"], "header from GET");

        let body = response.collect().await.unwrap().to_bytes();
        assert_eq!(&body[..], b"body from GET");
    }

    #[tokio::test]
    async fn test_imlicit_head() {
        let app = app();

        let response = app
            .oneshot(Request::head("/get-head").body(Body::empty()).unwrap())
            .await
            .unwrap();
        assert_eq!(response.status(), StatusCode::OK);
        assert_eq!(response.headers()["x-some-header"], "header from HEAD");

        let body = response.collect().await.unwrap().to_bytes();
        assert!(body.is_empty());
    }
}
