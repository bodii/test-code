use std::net::{Ipv4Addr, Ipv6Addr, SocketAddr};

use axum::{extract::Request, routing::get, Router};
use hyper::body::Incoming;
use hyper_util::{
    rt::{TokioExecutor, TokioIo},
    server,
};
use tokio::net::TcpListener;
use tower::Service;

#[tokio::main]
async fn main() {
    let app = app();

    let localhost_v4 = SocketAddr::new(Ipv4Addr::LOCALHOST.into(), 8080);
    let listener_v4 = TcpListener::bind(&localhost_v4).await.unwrap();

    let localhost_v6 = SocketAddr::new(Ipv6Addr::LOCALHOST.into(), 8080);
    let listener_v6 = TcpListener::bind(&localhost_v6).await.unwrap();

    loop {
        // Accept connections from `listener_v4` and `listener_v6` at the same time
        let (socket, _remote_addr) = tokio::select! {
            result = listener_v4.accept() => {
                result.unwrap()
            }
            result = listener_v6.accept() => {
                result.unwrap()
            }
        };

        let tower_service = app.clone();

        tokio::spawn(async move {
            let socket = TokioIo::new(socket);

            let hyper_service = hyper::service::service_fn(move |request: Request<Incoming>| {
                tower_service.clone().call(request)
            });

            if let Err(err) = server::conn::auto::Builder::new(TokioExecutor::new())
                .serve_connection_with_upgrades(socket, hyper_service)
                .await
            {
                eprintln!("failed to serve connection: {err:#}");
            }
        });
    }
}

fn app() -> Router {
    Router::new().route("/", get(|| async { "Hello, World!" }))
}
