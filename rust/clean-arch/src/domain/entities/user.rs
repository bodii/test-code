use diesel::prelude::*;
use serde::Serialize;

#[derive(Debug, Clone, Serialize, Queryable, Selectable)]
#[diesel(table_name=crate::schema::users)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct User {
    pub id: i32,
    pub name: String,
    pub email: String,
    pub phone: String,
    pub address: String,
}
