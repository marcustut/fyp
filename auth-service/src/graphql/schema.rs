use crate::graphql::utils::get_user_repo_from_ctx;
use async_graphql::validators::StringMinLength;
use async_graphql::{Context, EmptySubscription, Error, ErrorExtensions, Object, Schema};
use serde_json::json;
use validator::Validate;

use crate::models::user::{NewUser, User};

use super::utils::get_crypto_service_from_ctx;

pub type AppSchema = Schema<Query, Mutation, EmptySubscription>;

pub struct Query;

#[Object]
impl Query {
    async fn add(&self, a: i32, b: i32) -> i32 {
        a + b
    }

    async fn get_user(
        &self,
        ctx: &Context<'_>,
        #[graphql(validator(StringMinLength(length = "3")))] username: String,
    ) -> Result<User, Error> {
        // get user repository
        let user_repo = get_user_repo_from_ctx(ctx);

        // find the user with a username
        let result = user_repo.find(username.clone()).await;

        match result {
            Ok(user) => Ok(user),
            _ => Err(Error::new(format!(
                "Can't find user with username `{}`",
                username
            ))),
        }
    }
}

pub struct Mutation;

#[Object]
impl Mutation {
    async fn create_user(&self, ctx: &Context<'_>, new_user: NewUser) -> Result<User, Error> {
        let user_repo = get_user_repo_from_ctx(ctx);
        let crypto_service = get_crypto_service_from_ctx(ctx);

        // validate the input
        match new_user.validate() {
            Ok(_) => (),
            Err(e) => {
                return Err(Error::new("Failed validating NewUser")
                    .extend_with(|_, e| e.set("validation_error", "HAHHA")))
            }
        };

        // create the user
        let result = user_repo.create(new_user, &crypto_service).await;
        match result {
            Ok(user) => Ok(user),
            _ => Err(Error::new("unable to create user")),
        }
    }
}
