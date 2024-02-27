use std::net::SocketAddr;

use axum::{
    extract::State,
    http::StatusCode,
    response::Result,
    routing::{get, post},
    Json, Router,
};
use diesel_migrations::{embed_migrations, EmbeddedMigrations, MigrationHarness};

use diesel::prelude::*;
use tracing_subscriber::{layer::SubscriberExt, util::SubscriberInitExt};

pub const MIGRATIONS: EmbeddedMigrations = embed_migrations!("migrations/");

table! {
    users (id) {
    id -> Integer,
    name -> Text,
    hair_color -> Nullable<Text>,
    }
}

#[derive(serde::Serialize, Selectable, Queryable)]
struct User {
    id: i32,
    name: String,
    hair_color: Option<String>,
}

#[derive(serde::Deserialize, Insertable)]
#[diesel(table_name = users)]
struct NewUser {
    name: String,
    hair_color: Option<String>,
}

fn subscriber() {
    let layer = tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| "diesel_postgres=debug".into());

    tracing_subscriber::registry()
        .with(layer)
        .with(tracing_subscriber::fmt::layer())
        .init();
}

const DEFAULT_DB_URL: &str = "postgres://postgres:123456@localhost/test";

#[tokio::main]
async fn main() {
    subscriber();

    let db_url = std::env::var("DATABASE_URL").unwrap_or_else(|_| DEFAULT_DB_URL.into());

    // set up connection pool
    let manager = deadpool_diesel::postgres::Manager::new(db_url, deadpool_diesel::Runtime::Tokio1);
    let pool = deadpool_diesel::postgres::Pool::builder(manager)
        .build()
        .unwrap();

    {
        let conn = pool.get().await.unwrap();
        conn.interact(|conn| conn.run_pending_migrations(MIGRATIONS).map(|_| ()))
            .await
            .unwrap()
            .unwrap();
    }

    let app = Router::new()
        .route("/user/list", get(list_users))
        .route("/user/create", post(create_user))
        .with_state(pool);

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));

    tracing::debug!("listening on {addr}");

    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();

    axum::serve(listener, app).await.unwrap();
}

async fn create_user(
    State(pool): State<deadpool_diesel::postgres::Pool>,
    Json(new_user): Json<NewUser>,
) -> Result<Json<User>, (StatusCode, String)> {
    let conn = pool.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| {
            diesel::insert_into(users::table)
                .values(new_user)
                .returning(User::as_returning())
                .get_result(conn)
        })
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}

async fn list_users(
    State(pool): State<deadpool_diesel::postgres::Pool>,
) -> Result<Json<Vec<User>>, (StatusCode, String)> {
    let conn = pool.get().await.map_err(internal_error)?;
    let res = conn
        .interact(|conn| users::table.select(User::as_select()).load(conn))
        .await
        .map_err(internal_error)?
        .map_err(internal_error)?;

    Ok(Json(res))
}

fn internal_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::INTERNAL_SERVER_ERROR, err.to_string())
}
