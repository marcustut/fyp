use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Deserialize, Serialize, sqlx::FromRow)]
pub struct User {
    pub id: Uuid,
    pub username: String,
    pub email: String,
    #[serde(skip)]
    pub password_hash: String,
    pub full_name: Option<String>,
    pub bio: Option<String>,
    pub image: Option<String>,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

// TODO: Validate does not work currently
#[derive(Debug, Deserialize, Validate)]
pub struct NewUser {
    #[validate(length(min = 3))]
    pub username: String,
    #[validate(email)]
    pub email: String,
    #[validate(length(min = 3))]
    pub password: String,
}

// TODO: Validate does not work currently
#[derive(Debug, Deserialize, Validate)]
pub struct UpdateProfile {
    pub full_name: Option<String>,
    pub bio: Option<String>,
    #[validate(url)]
    pub image: Option<String>,
}
