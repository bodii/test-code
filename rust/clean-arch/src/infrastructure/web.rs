use actix_web::{middleware::Logger, web, App, HttpServer};
use log::info;

use crate::presentation::routes;

use super::repositories::postgres_user_repository::PostgresUserRepository;

pub async fn run() -> std::io::Result<()> {
    let repo = PostgresUserRepository::new();
    let app_data = web::Data::new(repo);

    info!("Starrting server...!");

    HttpServer::new(move || {
        App::new()
            .app_data(app_data.clone())
            .wrap(Logger::default())
            .configure(routes::user_routes::routes)
    })
    .bind("127.0.0.1:8080")
    .unwrap()
    .run()
    .await
}
