use async_graphql::{http::GraphiQLSource, EmptySubscription, Schema};
use async_graphql_axum::GraphQL;
use axum::{
    http::Method,
    response::{Html, IntoResponse},
    routing::get,
    Router,
};
use files::{MutationRoot, QueryRoot, Storage};
use tokio::net::TcpListener;
use tower_http::cors::{AllowOrigin, CorsLayer};

#[tokio::main]
async fn main() {
    let schema = Schema::build(QueryRoot, MutationRoot, EmptySubscription)
        .data(Storage::default())
        .finish();

    println!("GraphiQL IDE: http://localhost:3000");

    let cors_layer = CorsLayer::new()
        .allow_origin(AllowOrigin::predicate(|_, _| true))
        .allow_methods([Method::GET, Method::POST]);

    let app = Router::new()
        .route("/", get(graphiql).post_service(GraphQL::new(schema)))
        .layer(cors_layer);

    axum::serve(TcpListener::bind("127.0.0.1:3000").await.unwrap(), app)
        .await
        .unwrap();
}
async fn graphiql() -> impl IntoResponse {
    Html(GraphiQLSource::build().endpoint("/").finish())
}
