use axum::{response::Html, routing::get, Form, Router};
use serde::Deserialize;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

#[derive(Deserialize, Debug)]
#[allow(dead_code)]
struct Input {
    name: String,
    email: String,
}

#[tokio::main]
async fn main() {
    let filter = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "example_form=debug".into());
    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();

    let app = Router::new().route("/", get(show_form).post(accept_form));

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", listener.local_addr().unwrap());

    axum::serve(listener, app).await.unwrap();
}

// http :3000
async fn show_form() -> Html<&'static str> {
    let html = r#"
    <!doctype html>
    <html>
        <head></head>
        <body>
            <form action="/" method="post">
                <label for="name">
                    Enter your name:
                    <input type="text" name="name">
                </label>

                <label for="email">
                    Enter your email:
                    <input type="text" name="email">
                </label>

                <input type="submit" value="Subscribe!">
            </form>
        </body>
    </html>
    "#;

    Html(html)
}

//  http --form post :3000 name="Joy" email="joy@example.com"
//  http -f post :3000 name="Joy" email="joy@example.com"
//
//  http -f post :3000 \
// name=aa \
// email=bb

async fn accept_form(Form(input): Form<Input>) {
    dbg!(&input);
}
