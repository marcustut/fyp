use async_graphql::{InputObject, SimpleObject};
use chrono::{DateTime, Utc};
use regex::Regex;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

lazy_static! {
    static ref RE_HAS_ONE_ALPHABET: Regex = Regex::new(r"[A-Za-z]+").unwrap();
    static ref RE_HAS_ONE_NUMBER: Regex = Regex::new(r"[0-9]+").unwrap();
}

#[derive(Debug, Deserialize, Serialize, sqlx::FromRow, SimpleObject)]
pub struct User {
    pub id: i64,
    pub username: String,
    pub email: String,
    #[serde(skip)]
    #[graphql(skip)]
    pub password_hash: String,
    pub full_name: Option<String>,
    pub bio: Option<String>,
    pub avatar_url: Option<String>,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

#[derive(Debug, Deserialize, Serialize, SimpleObject)]
pub struct UserWithAuth {
    pub access_token: String,
    pub expired_at: DateTime<Utc>,
    pub user: User,
}

#[derive(Clone, Debug, Deserialize, Validate, InputObject)]
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
    #[validate(length(min = 3, message = "must be at least 3 characters"))]
    pub full_name: Option<String>,
    #[validate(length(min = 5, message = "must be at least 5 characters"))]
    pub bio: Option<String>,
    #[validate(url(message = "avatar_url must be a valid URL"))]
    pub avatar_url: Option<String>,
}

#[derive(Debug, Deserialize, Validate, InputObject)]
pub struct UpdateProfile {
    #[validate(length(min = 3, message = "must be at least 3 characters"))]
    pub full_name: Option<String>,
    #[validate(length(min = 5, message = "must be at least 5 characters"))]
    pub bio: Option<String>,
    #[validate(url(message = "avatar_url must be a valid URL"))]
    pub avatar_url: Option<String>,
}

impl UpdateProfile {
    pub fn as_hashmap(&self) -> HashMap<String, String> {
        let mut hashmap = HashMap::new();

        match self.full_name.as_ref() {
            Some(x) => {
                hashmap.insert("full_name".to_string(), x.clone());
            }
            None => {
                hashmap.insert("full_name".to_string(), "".to_string());
            }
        };
        match self.bio.as_ref() {
            Some(x) => {
                hashmap.insert("bio".to_string(), x.clone());
            }
            None => {
                hashmap.insert("bio".to_string(), "".to_string());
            }
        };
        match self.avatar_url.as_ref() {
            Some(x) => {
                hashmap.insert("avatar_url".to_string(), x.clone());
            }
            None => {
                hashmap.insert("avatar_url".to_string(), "".to_string());
            }
        };

        hashmap
    }
}

#[derive(Clone, Debug, Deserialize, Validate, InputObject)]
pub struct SignInWithEmail {
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

#[derive(Clone, Debug, Deserialize, Validate, InputObject)]
pub struct SignInWithUsername {
    #[validate(length(min = 3, message = "must be at least 3 characters"))]
    pub username: String,
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
