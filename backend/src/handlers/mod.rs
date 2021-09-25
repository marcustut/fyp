pub mod health;
pub mod test;
pub mod user;
pub mod graphql;

use actix_web::web::ServiceConfig;

pub fn app_config(config: &mut ServiceConfig) {
    config.service(health::health);
    config.service(user::find);
    config.service(user::create);
    config.service(user::delete);
    config.service(graphql::graphiql);
    config.service(graphql::graphql);
}
