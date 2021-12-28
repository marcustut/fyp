use std::sync::Arc;

use eyre::Result;
use sqlx::MySqlPool;

use crate::{
    config::crypto::CryptoService,
    models::user::{NewUser, UpdateProfile, User},
};

#[derive(Clone)]
pub struct UserRepository {
    pool: Arc<MySqlPool>,
}

impl UserRepository {
    pub fn new(pool: Arc<MySqlPool>) -> Self {
        Self { pool }
    }

    pub async fn find_by_username(&self, username: String) -> Result<User> {
        let user = sqlx::query_as!(User, "SELECT * FROM users WHERE username = ?", username)
            .fetch_one(&*self.pool)
            .await?;

        Ok(user)
    }

    pub async fn find_by_email(&self, email: String) -> Result<User> {
        let user = sqlx::query_as!(User, "SELECT * FROM users WHERE email = ?", email)
            .fetch_one(&*self.pool)
            .await?;

        Ok(user)
    }

    pub async fn create(&self, new_user: NewUser, crypto_service: &CryptoService) -> Result<User> {
        let password_hash: String = crypto_service.hash_password(new_user.password).await?;

        let mut tx = self.pool.begin().await?;

        let user_id = sqlx::query!(
            r#"
                INSERT INTO users (
                    username, 
                    email, 
                    password_hash,
                    full_name,
                    avatar_url,
                    bio
                )
                VALUES (
                    ?, 
                    ?, 
                    ?,
                    ?,
                    ?,
                    ?
                ) 
            "#,
            new_user.username,
            new_user.email,
            password_hash,
            new_user.full_name,
            new_user.avatar_url,
            new_user.bio
        )
        .execute(&mut tx)
        .await?
        .last_insert_id();

        let user = sqlx::query_as!(User, "SELECT * FROM users WHERE id = ?", user_id)
            .fetch_one(&mut tx)
            .await?;

        tx.commit().await?;

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

        let mut tx = self.pool.begin().await?;

        sqlx::query(
            format!(
                r#"
                    UPDATE users
                    SET {}
                    WHERE username = ?
                "#,
                dynamic_update_statement
            )
            .as_str(),
        )
        .bind(username.clone())
        .execute(&mut tx)
        .await?;

        let user = sqlx::query_as!(User, "SELECT * FROM users WHERE username = ?", username)
            .fetch_one(&mut tx)
            .await?;

        tx.commit().await?;

        Ok(user)
    }

    pub async fn delete(&self, username: String) -> Result<User> {
        let mut tx = self.pool.begin().await?;

        let user = sqlx::query_as!(User, "SELECT * FROM users WHERE username = ?", username)
            .fetch_one(&mut tx)
            .await?;

        sqlx::query!(
            r#"
                DELETE FROM users 
                WHERE username = ? 
            "#,
            username
        )
        .execute(&mut tx)
        .await?;

        tx.commit().await?;

        Ok(user)
    }
}
