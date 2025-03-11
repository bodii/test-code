use std::sync::Arc;

use async_trait::async_trait;
use diesel::prelude::*;
use diesel::{self, RunQueryDsl};

use crate::domain::{entities::user::User, repositories::user_repository::UserRepository};
use crate::infrastructure::db::connection::{establish_connection, DBPool};
use crate::presentation::handlers::user_handler::NewUser;
use crate::schema;
use crate::schema::users::dsl::users;
use crate::schema::users::email;

#[derive(Clone)]
pub struct PostgresUserRepository {
    pool: DBPool,
}

impl PostgresUserRepository {
    pub fn new() -> Self {
        // println!("url: {}", std::env::var("DATABASE_URL").unwrap());
        let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
        Self {
            pool: establish_connection(&database_url),
        }
    }
}

#[async_trait]
impl UserRepository for Arc<PostgresUserRepository> {
    async fn find_by_email(&self, input_email: String) -> Option<User> {
        let mut conn = self.pool.get().unwrap();
        users
            .filter(email.eq(input_email))
            .first::<User>(&mut conn)
            .optional()
            .expect("Error loading user")
    }

    async fn save(&self, user: &NewUser) -> Result<(), diesel::result::Error> {
        diesel::insert_into(schema::users::table)
            .values(user)
            .execute(&mut self.pool.get().unwrap())
            .unwrap();

        Ok(())
    }
}
