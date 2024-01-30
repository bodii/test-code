use axum::{
    response::{Html, Response},
    routing::get,
    Router,
};
use http::{Request, StatusCode};
use tower_service::Service;

fn main() {
    let request: Request<String> = Request::builder()
        .uri("https://serverless.example/api/")
        .body("Some Body Data".into())
        .unwrap();

    let response: Response = futures_executor::block_on(app(request));
    assert_eq!(StatusCode::OK, response.status());
}

#[allow(clippy::let_and_return)]
async fn app(request: Request<String>) -> Response {
    let mut router = Router::new().route("/api/", get(index));
    let response = router.call(request).await.unwrap();

    response
}

async fn index() -> Html<&'static str> {
    Html("<h1>Hello, World!</h1>")
}
