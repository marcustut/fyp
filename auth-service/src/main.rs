#[macro_use]
extern crate validator_derive;
#[macro_use]
extern crate lazy_static;

mod config;
mod db;
mod handlers;
mod models;

use crate::config::Config;
use crate::db::user::UserRepository;
use actix_web::{middleware::Logger, web, App, HttpServer};
use color_eyre::Result;
use handlers::app_config;
use std::sync::Arc;
use tracing::info;

// This is an immutable application state
struct AppState {
    app_name: String,
}

#[actix_rt::main]
async fn main() -> Result<()> {
    let config: Config = Config::from_env(true).expect("Server configuration");

    let pool = config.db_pool().await.expect("Database configuration");

    let crypto_service = config.crypto_service();

    info!("Starting server at http://{}:{}", config.host, config.port);
    HttpServer::new(move || {
        App::new()
            .wrap(Logger::default())
            .app_data(web::Data::new(AppState {
                app_name: String::from("Auth"),
            }))
            .app_data(web::Data::new(pool.clone()))
            .app_data(web::Data::new(crypto_service.clone()))
            .app_data(web::Data::new(UserRepository::new(Arc::new(pool.clone()))))
            .configure(app_config)
    })
    .bind(format!("{}:{}", config.host, config.port))?
    .run()
    .await?;

    Ok(())
}
