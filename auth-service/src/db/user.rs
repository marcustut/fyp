use crate::{
    config::crypto::CryptoService,
    models::user::{NewUser, UpdateProfile, User},
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
            r#"
                SELECT 
                    *
                FROM 
                    users 
                WHERE 
                    username = $1 
                LIMIT 1
            "#,
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
            r#"
                INSERT INTO users (
                    username, 
                    email, 
                    password_hash
                ) 
                VALUES (
                    $1, 
                    $2, 
                    $3
                ) 
                RETURNING *
            "#,
            new_user.username,
            new_user.email,
            password_hash
        )
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }

    pub async fn update(&self, update_profile: UpdateProfile, username: String) -> Result<User> {
        let dynamic_update_statement: Vec<String> = update_profile
            .as_hashmap()
            .into_iter()
            .map(|kv| -> String { format!("{} = nullif('{}', ''), ", kv.0, kv.1) })
            .collect();
        let dynamic_update_statement = dynamic_update_statement.join(" ");
        let dynamic_update_statement: String = dynamic_update_statement
            .chars()
            .take(dynamic_update_statement.len() - 2)
            .collect();

        println!("{}\n", dynamic_update_statement);

        let user = sqlx::query_as(
            format!(
                r#"
                UPDATE users
                SET {}
                WHERE username = $1
                RETURNING *
            "#,
                dynamic_update_statement
            )
            .as_str(),
        )
        .bind(username)
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }

    pub async fn delete(&self, username: String) -> Result<User> {
        let user = sqlx::query_as!(
            User,
            r#"
                DELETE FROM users 
                WHERE username = $1 
                RETURNING *
            "#,
            username
        )
        .fetch_one(&*self.pool)
        .await?;

        Ok(user)
    }
}
