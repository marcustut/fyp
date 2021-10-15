use crate::{
    config::crypto::CryptoService,
    models::user::{NewUser, User},
};
use eyre::Result;
use sqlx::PgPool;
use std::sync::Arc;

#[derive(Clone)]
pub struct UserRepository {
    pool: Arc<PgPool>,
}

impl UserRepository {
    pub fn new(pool: Arc<PgPool>) -> Self {
        Self { pool }
    }

    pub async fn find(&self, username: String) -> Result<User> {
        let user = sqlx::query_as!(
            User,
            "select * from users where username = $1 limit 1",
            username
        )
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }

    pub async fn create(&self, new_user: NewUser, crypto_service: &CryptoService) -> Result<User> {
        let password_hash: String = crypto_service.hash_password(new_user.password).await?;

        let user = sqlx::query_as!(
            User,
            "insert into users (username, email, password_hash) values ($1, $2, $3) returning *",
            new_user.username,
            new_user.email,
            password_hash
        )
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }

    pub async fn delete(&self, username: String) -> Result<User> {
        let user = sqlx::query_as!(
            User,
            "delete from users where username = $1 returning *",
            username
        )
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }
}
