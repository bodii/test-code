#[cfg(not(unix))]
fn main() {
    println!("This example requires unix");
}

#[cfg(unix)]
#[tokio::main]
async fn main() {
    unix::server().await;
}

#[cfg(unix)]
mod unix {
    use std::{convert::Infallible, path::PathBuf, sync::Arc};

    use axum::{
        body::Body,
        extract::{connect_info, ConnectInfo},
        routing::get,
        Router,
    };
    use http_body_util::BodyExt;
    use hyper::{body::Incoming, Method, Request, StatusCode};
    use hyper_util::{
        rt::{TokioExecutor, TokioIo},
        server,
    };
    use tokio::net::{unix::UCred, UnixListener, UnixStream};
    use tower::Service;
    use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

    pub async fn server() {
        subscriber();

        let path = PathBuf::from("/tmp/axum/helloworld");

        let _ = tokio::fs::remove_file(&path).await;
        tokio::fs::create_dir_all(path.parent().unwrap())
            .await
            .unwrap();

        let uds = UnixListener::bind(path.clone()).unwrap();
        tokio::spawn(async move {
            let app = Router::new().route("/", get(handle));

            let mut make_service = app.into_make_service_with_connect_info::<UdsConnectInfo>();

            loop {
                let (socket, _remote_addr) = uds.accept().await.unwrap();

                let tower_service = unwrap_infallible(make_service.call(&socket).await);

                tokio::spawn(async move {
                    let socket = TokioIo::new(socket);

                    let hyper_service =
                        hyper::service::service_fn(move |request: Request<Incoming>| {
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
        });

        let stream = TokioIo::new(UnixStream::connect(path).await.unwrap());
        let (mut sender, conn) = hyper::client::conn::http1::handshake(stream).await.unwrap();
        tokio::task::spawn(async move {
            if let Err(err) = conn.await {
                println!("connection error: {err:#}");
            }
        });

        let request = Request::builder()
            .method(Method::GET)
            .uri("http://uri-doesnt-matter.com")
            .body(Body::empty())
            .unwrap();

        let response = sender.send_request(request).await.unwrap();

        assert_eq!(response.status(), StatusCode::OK);

        let body = response.collect().await.unwrap().to_bytes();
        let body = String::from_utf8(body.to_vec()).unwrap();

        assert_eq!(body, "Hello, World!");
    }

    fn subscriber() {
        let layer = tracing_subscriber::EnvFilter::try_from_default_env()
            .unwrap_or_else(|_| "unix_domain_socket=debug".into());

        tracing_subscriber::registry()
            .with(layer)
            .with(tracing_subscriber::fmt::layer())
            .init();
    }

    async fn handle(ConnectInfo(info): ConnectInfo<UdsConnectInfo>) -> &'static str {
        println!("new connection from `{:?}`", info);

        "Hello, World!"
    }

    #[derive(Clone, Debug)]
    #[allow(dead_code)]
    struct UdsConnectInfo {
        peer_addr: Arc<tokio::net::unix::SocketAddr>,
        peert_cred: UCred,
    }

    impl connect_info::Connected<&UnixStream> for UdsConnectInfo {
        fn connect_info(target: &UnixStream) -> Self {
            let peer_addr = target.peer_addr().unwrap();
            let peert_cred = target.peer_cred().unwrap();

            Self {
                peer_addr: Arc::new(peer_addr),
                peert_cred,
            }
        }
    }

    fn unwrap_infallible<T>(result: Result<T, Infallible>) -> T {
        match result {
            Ok(value) => value,
            Err(err) => match err {},
        }
    }
}
