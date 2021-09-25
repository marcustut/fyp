pub mod health;
pub mod test;
pub mod user;

use actix_web::web::ServiceConfig;

pub fn app_config(config: &mut ServiceConfig) {
    config.service(health::health);
    config.service(user::find);
    config.service(user::create);
    config.service(user::delete);
}
