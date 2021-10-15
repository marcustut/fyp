use crate::config::crypto::CryptoService;
use crate::db::user::UserRepository;
use crate::models::user::NewUser;
use actix_web::{delete, get, post, web, HttpResponse, Responder};
use validator::Validate;

#[get("/user/{username}")]
async fn find(user_repo: web::Data<UserRepository>, username: web::Path<String>) -> impl Responder {
    let username = username.into_inner();
    let result = user_repo.get_ref().find(username).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to find user"),
    }
}

#[post("/user")]
async fn create(
    user_repo: web::Data<UserRepository>,
    crypto_service: web::Data<CryptoService>,
    new_user: web::Json<NewUser>,
) -> impl Responder {
    // unwrap the param
    let new_user = new_user.into_inner();
    // validate the new_user input
    match new_user.validate() {
        Ok(_) => (),
        Err(e) => return HttpResponse::BadRequest().json(e),
    };
    // create the user
    let result = user_repo.get_ref().create(new_user, &crypto_service).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to create user"),
    }
}

#[delete("/user/{username}")]
pub(crate) async fn delete(
    user_repo: web::Data<UserRepository>,
    username: web::Path<String>,
) -> impl Responder {
    let username = username.into_inner();
    // delete the user
    let result = user_repo.get_ref().delete(username).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to delete user"),
    }
}
