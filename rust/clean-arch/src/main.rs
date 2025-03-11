use crate::infrastructure::web::run;

mod application;
mod domain;
mod infrastructure;
mod presentation;
mod schema;
use dotenv::dotenv;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();

    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();

    run().await
}
