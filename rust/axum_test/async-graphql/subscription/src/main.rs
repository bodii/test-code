use tokio::net::TcpListener;

use async_graphql::{http::GraphiQLSource, Schema};
use async_graphql_axum::{GraphQL, GraphQLSubscription};
use axum::{
    response::{self, IntoResponse},
    routing::get,
    Router,
};
use books::{MutationRoot, QueryRoot, Storage, SubscriptionRoot};

async fn graphiql() -> impl IntoResponse {
    response::Html(
        GraphiQLSource::build()
            .endpoint("/")
            .subscription_endpoint("/ws")
            .finish(),
    )
}

#[tokio::main]
async fn main() {
    let schema = Schema::build(QueryRoot, MutationRoot, SubscriptionRoot)
        .data(Storage::default())
        .finish();

    let app = Router::new()
        .route(
            "/",
            get(graphiql).post_service(GraphQL::new(schema.clone())),
        )
        .route_service("/ws", GraphQLSubscription::new(schema));

    println!("GraphiQL IDE: http://localhost:3000");

    axum::serve(TcpListener::bind("127.0.0.1:3000").await.unwrap(), app)
        .await
        .unwrap();
}
