use async_graphql::{InputObject, SimpleObject};
use chrono::NaiveDateTime;
use regex::Regex;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

lazy_static! {
    static ref RE_HAS_ONE_ALPHABET: Regex = Regex::new(r"[A-Za-z]+").unwrap();
    static ref RE_HAS_ONE_NUMBER: Regex = Regex::new(r"[0-9]+").unwrap();
}

#[derive(Debug, Deserialize, Serialize, sqlx::FromRow, SimpleObject)]
pub struct User {
    pub id: Uuid,
    pub username: String,
    pub email: String,
    #[serde(skip)]
    #[graphql(skip)]
    pub password_hash: String,
    pub full_name: Option<String>,
    pub bio: Option<String>,
    pub image: Option<String>,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

#[derive(Debug, Deserialize, Validate, InputObject)]
pub struct NewUser {
    #[validate(length(min = 3, message = "must be at least 3 characters"))]
    pub username: String,
    #[validate(email(message = "must be a valid email"))]
    pub email: String,
    #[validate(length(min = 8, message = " Password must have at least 8 characters"))]
    #[validate(regex(
        path = "RE_HAS_ONE_ALPHABET",
        message = "Password must have at least one alphabet"
    ))]
    #[validate(regex(
        path = "RE_HAS_ONE_NUMBER",
        message = "Password must have at least one number"
    ))]
    pub password: String,
}

#[derive(Debug, Deserialize, Validate)]
pub struct UpdateProfile {
    #[validate(length(min = 3, message = "must be at least 3 characters"))]
    pub full_name: Option<String>,
    #[validate(length(min = 5, message = "must be at least 5 characters"))]
    pub bio: Option<String>,
    #[validate(url(message = "image must be a URL"))]
    pub image: Option<String>,
}
