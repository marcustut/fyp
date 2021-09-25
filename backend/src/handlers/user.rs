use crate::config::crypto::CryptoService;
use crate::db::user::UserRepository;
use crate::models::user::NewUser;
use actix_web::{delete, get, post, web, HttpResponse, Responder};

#[get("/user/{username}")]
async fn find(username: web::Path<String>, user_repo: web::Data<UserRepository>) -> impl Responder {
    let username = username.into_inner();
    let result = user_repo.get_ref().find(username).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to find user"),
    }
}

#[post("/user")]
async fn create(
    new_user: web::Json<NewUser>,
    user_repo: web::Data<UserRepository>,
    crypto_service: web::Data<CryptoService>,
) -> impl Responder {
    // create the user
    let result = user_repo
        .get_ref()
        .create(new_user.into_inner(), &crypto_service)
        .await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to create user"),
    }
}

#[delete("/user/{username}")]
async fn delete(
    username: web::Path<String>,
    user_repo: web::Data<UserRepository>,
) -> impl Responder {
    let username = username.into_inner();
    // delete the user
    let result = user_repo.get_ref().delete(username).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to delete user"),
    }
}
