// use std::borrow::Cow;
// use std::collections::BTreeMap;

use crate::graphql::utils::get_user_repo_from_ctx;
use async_graphql::validators::StringMinLength;
// use async_graphql::Value::Null;
use async_graphql::{Context, EmptySubscription, Error, ErrorExtensions, Object, Schema, Value};
// use serde_json::json;

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
        let result = user_repo.create(new_user, &crypto_service).await;
        match result {
            Ok(user) => Ok(user),
            Err(report) => Err(Error::new(report.to_string())),
        }
    }

    async fn delete_user(
        &self,
        ctx: &Context<'_>,
        #[graphql(validator(StringMinLength(length = "3")))] username: String,
    ) -> Result<User, Error> {
        let user_repo = get_user_repo_from_ctx(ctx);

        let result = user_repo.delete(username).await;
        match result {
            Ok(user) => Ok(user),
            Err(report) => Err(Error::new(report.to_string())),
        }
    }
}
