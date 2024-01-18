use axum::{response::Html, routing::get, Router};
use listenfd::ListenFd;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    let app = Router::new().route("/", get(handler));

    let mut listenfd = ListenFd::from_env();
    let listener = match listenfd.take_tcp_listener(0).unwrap() {
        Some(listener) => TcpListener::from_std(listener).unwrap(),
        None => TcpListener::bind("127.0.0.1:3000").await.unwrap(),
    };

    println!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap();
}

async fn handler() -> Html<&'static str> {
    Html("<h1>Hello, World!</h1>")
}
