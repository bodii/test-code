use super::schema::users;
use diesel::prelude::*;

#[derive(serde::Serialize, Selectable, Queryable)]
pub struct User {
    pub id: i32,
    pub name: String,
    pub hair_color: Option<String>,
}

#[derive(serde::Deserialize, Insertable)]
#[diesel(table_name = users)]
pub struct NewUser {
    pub name: String,
    pub hair_color: Option<String>,
}
