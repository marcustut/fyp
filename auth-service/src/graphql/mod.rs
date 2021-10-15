pub mod schema;

use async_graphql::*;

pub fn new_schema() -> schema::AppSchema {
    Schema::new(schema::Query, EmptyMutation, EmptySubscription)
}
