use axum::{
    async_trait,
    extract::{FromRef, FromRequestParts, State},
    http::{request::Parts, StatusCode},
    routing::get,
    Router,
};
use mysql_async::{prelude::*, Opts};
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt, EnvFilter};

#[tokio::main]
async fn main() {
    subscriber();

    let database_url = Opts::from_url("mysql://root:123456@localhost:3306/").unwrap();
    let pool = mysql_async::Pool::new(database_url);

    let app = Router::new()
        .route(
            "/",
            get(using_connection_pool_extractor).post(using_connection_extractor),
        )
        .with_state(pool);

    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();

    tracing::debug!("listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}

fn subscriber() {
    let filter = EnvFilter::try_from_default_env().unwrap_or_else(|_| "tokio_mysql=debug".into());
    tracing_subscriber::registry()
        .with(filter)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

fn internal_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::INTERNAL_SERVER_ERROR, err.to_string())
}

type ConnectionPool = mysql_async::Pool;
async fn using_connection_pool_extractor(
    State(pool): State<ConnectionPool>,
) -> Result<String, (StatusCode, String)> {
    let mut conn = pool.get_conn().await.map_err(internal_error)?;

    let row: Option<String> = conn
        .query_first("select 1 + 1")
        .await
        .map_err(internal_error)?;

    let two = row.unwrap_or_else(|| "0".into());
    Ok(two)
}

struct DatabaseConnection(mysql_async::Conn);

#[async_trait]
impl<S> FromRequestParts<S> for DatabaseConnection
where
    ConnectionPool: FromRef<S>,
    S: Send + Sync,
{
    type Rejection = (StatusCode, String);

    async fn from_request_parts(_parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        let pool = ConnectionPool::from_ref(state);

        let conn = pool.get_conn().await.map_err(internal_error)?;

        Ok(Self(conn))
    }
}

async fn using_connection_extractor(
    DatabaseConnection(mut conn): DatabaseConnection,
) -> Result<String, (StatusCode, String)> {
    let row: Option<String> = conn
        .query_first("select 1 + 2")
        .await
        .map_err(internal_error)?;
    let two = row.unwrap_or_else(|| "0".into());

    Ok(two)
}
