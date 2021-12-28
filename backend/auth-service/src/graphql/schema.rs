use crate::graphql::utils::get_user_repo_from_ctx;
use async_graphql::{Context, EmptySubscription, Error, ErrorExtensions, Object, Schema, Value};

use validator::Validate;

use crate::models::user::{NewUser, UpdateProfile, User};

use super::utils::get_crypto_service_from_ctx;

pub type AppSchema = Schema<Query, Mutation, EmptySubscription>;

pub struct Query;

#[Object]
impl Query {
    // async fn hash_password(&self, ctx: &Context<'_>, password: String) -> Result<String, Error> {
    //     let crypto_service = get_crypto_service_from_ctx(ctx);
    //     match crypto_service.hash_password(password).await {
    //         Ok(pwd) => Ok(pwd),
    //         Err(e) => Err(Error::new(format!("{:#}", e))),
    //     }
    // }

    async fn get_user(
        &self,
        ctx: &Context<'_>,
        #[graphql(validator(min_length = 3))] username: String,
    ) -> Result<User, Error> {
        // get user repository
        let user_repo = get_user_repo_from_ctx(ctx);

        // find the user with a username
        let result = user_repo.find(username.clone()).await;

        match result {
            Ok(user) => Ok(user),
            Err(report) => {
                tracing::error!("Failed to fetch user: {:?}", report);
                Err(Error::new(format!(
                    "user with username `{}` does not exist",
                    username
                )))
            }
        }
    }
}

// TODO: Add mutation for `sign_in`, `sign_up`, `refresh_token`
pub struct Mutation;

#[Object]
impl Mutation {
    async fn create_user(&self, ctx: &Context<'_>, new_user: NewUser) -> Result<User, Error> {
        let user_repo = get_user_repo_from_ctx(ctx);
        let crypto_service = get_crypto_service_from_ctx(ctx);

        // validate the input
        match new_user.validate() {
            Ok(_) => (),
            Err(err) => {
                return Err(Error::new("error validating input").extend_with(|_, e| {
                    e.set(
                        "validation_error",
                        Value::from_json(serde_json::to_value(err.clone()).unwrap()).unwrap(),
                    )
                }))
            }
        };

        // create the user
        let result = user_repo.create(new_user.clone(), &crypto_service).await;
        match result {
            Ok(user) => Ok(user),
            Err(report) => {
                tracing::error!("Failed to create user: {:?}", report);
                Err(Error::new(format!(
                    "user with username `{}` already exists",
                    new_user.username
                )))
            }
        }
    }

    async fn update_user(
        &self,
        ctx: &Context<'_>,
        update_profile: UpdateProfile,
        #[graphql(validator(min_length = 3))] username: String,
    ) -> Result<User, Error> {
        let user_repo = get_user_repo_from_ctx(ctx);
        let result = user_repo.update(update_profile, username.clone()).await;
        match result {
            Ok(user) => Ok(user),
            Err(report) => {
                tracing::error!("Failed to update user: {:?}", report);
                Err(Error::new(format!(
                    "unable to update user with username `{}`",
                    username
                )))
            }
        }
    }

    async fn delete_user(
        &self,
        ctx: &Context<'_>,
        #[graphql(validator(min_length = 3))] username: String,
    ) -> Result<User, Error> {
        let user_repo = get_user_repo_from_ctx(ctx);

        let result = user_repo.delete(username.clone()).await;
        match result {
            Ok(user) => Ok(user),
            Err(report) => {
                tracing::error!("Failed to delete user: {:?}", report);
                Err(Error::new(format!(
                    "user with username `{}` does not exist",
                    username
                )))
            }
        }
    }
}
