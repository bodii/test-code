use axum::{
    body::Body,
    extract::{Request, State},
    response::{IntoResponse, Response},
    routing::get,
    Router,
};
use hyper::{StatusCode, Uri};
use hyper_util::{
    client::legacy::{connect::HttpConnector, Client as hyperClient},
    rt::TokioExecutor,
};

type Client = hyperClient<HttpConnector, Body>;

#[tokio::main]
async fn main() {
    tokio::spawn(server());

    let client: Client =
        hyperClient::<(), ()>::builder(TokioExecutor::new()).build(HttpConnector::new());

    let app = Router::new().route("/", get(handler)).with_state(client);

    let listener = tokio::net::TcpListener::bind("127.0.0.1:4000")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap();
}

async fn server() {
    let app = Router::new().route("/", get(|| async { "Hello, world!" }));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    println!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap();
}

async fn handler(State(client): State<Client>, mut req: Request) -> Result<Response, StatusCode> {
    let path = req.uri().path();
    let path_query = req
        .uri()
        .path_and_query()
        .map(|v| v.as_str())
        .unwrap_or(path);

    let uri = format!("http://127.0.0.1:3000{}", path_query);

    *req.uri_mut() = Uri::try_from(uri).unwrap();
    Ok(client
        .request(req)
        .await
        .map_err(|_| StatusCode::BAD_REQUEST)?
        .into_response())
}
