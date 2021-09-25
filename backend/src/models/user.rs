use chrono::{NaiveDateTime, NaiveDate};
use serde::{Deserialize, Serialize};
use uuid::Uuid;
use juniper::{GraphQLObject, GraphQLInputObject, FieldResult, EmptySubscription, RootNode};

#[derive(Debug, Deserialize, Serialize, sqlx::FromRow, GraphQLObject)]
#[graphql(description = "A user of the system")]
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

// TODO: Validate does not work currently
#[derive(Debug, Deserialize, Validate, GraphQLInputObject)]
#[graphql(description = "A user of the system")]
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

pub struct QueryRoot;

#[juniper::graphql_object]
impl QueryRoot {
    fn user(_id: String) -> FieldResult<User> {
        OK(User{
            id: Uuid::new_v4(),
            username: "test_user".to_string(),
            email: "test@test.com".to_string(),
            password_hash: "".to_string(),
            full_name: Option::from("Test User".to_string()),
            bio: Option::from("simple bio".to_string()),
            image: Option::from("https://www.google.com".to_string()),
            created_at: NaiveDate::from_ymd(2016, 7, 8).and_hms(9, 10, 11),
            updated_at: NaiveDate::from_ymd(2016, 7, 8).and_hms(9, 10, 11),
        })
    }
}

pub struct MutationRoot;

#[juniper::graphql_object]
impl MutationRoot {
    fn create_user(new_user: NewUser) -> FieldResult<User> {
        Ok(User{
            id: Uuid::new_v4(),
            username: new_user.username,
            email: new_user.email,
            password_hash: new_user.password,
            full_name: Option::from("Test User".to_string()),
            bio: Option::from("simple bio".to_string()),
            image: Option::from("https://www.google.com".to_string()),
            created_at: NaiveDate::from_ymd(2016, 7, 8).and_hms(9, 10, 11),
            updated_at: NaiveDate::from_ymd(2016, 7, 8).and_hms(9, 10, 11),
        })
    }
}

pub type Schema = RootNode<'static, QueryRoot, MutationRoot, EmptySubscription>;

pub fn create_schema() -> Schema {
    Schema::new(QueryRoot{}, MutationRoot{}, EmptySubscription::new())
}