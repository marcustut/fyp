use crate::config::crypto::CryptoService;
use crate::db::user::UserRepository;
use crate::models::user::NewUser;
use actix_web::{delete, get, post, web, HttpResponse, Responder};
use validator::Validate;

#[get("/user/{username}")]
pub(crate) async fn find(
    user_repo: web::Data<UserRepository>,
    username: web::Path<String>,
) -> impl Responder {
    let username = username.into_inner();
    let result = user_repo.get_ref().find_by_username(username).await;
    match result {
        Ok(user) => HttpResponse::Ok().json(user),
        _ => HttpResponse::BadRequest().body("Failed to find user"),
    }
}

#[post("/user")]
pub(crate) async fn create(
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

#[cfg(test)]
mod tests {
    use crate::handlers::app_config;
    use crate::models::user::User;
    use crate::Config;
    use crate::UserRepository;
    use actix_web::{test, web, App};
    use serde_json::json;
    use std::sync::Arc;

    #[actix_rt::test]
    async fn test_user() {
        let config: Config = Config::from_env(false).expect("Server configuration");

        let pool = config.new_db_pool().await.expect("Database configuration");

        let crypto_service = config.new_crypto_service();

        let mut app = test::init_service(
            App::new()
                .app_data(web::Data::new(pool.clone()))
                .app_data(web::Data::new(crypto_service.clone()))
                .app_data(web::Data::new(UserRepository::new(Arc::new(pool.clone()))))
                .configure(app_config),
        )
        .await;

        // the test user's username
        let test_username = String::from("test_user");

        // request body for creating a user
        let req_body = json!({
            "username": test_username,
            "email": "test@test.com",
            "password": "test1234"
        });

        // creating a user
        let resp = test::TestRequest::post()
            .uri("/user")
            .set_json(&req_body)
            .send_request(&mut app)
            .await;
        assert!(resp.status().is_success(), "Failed to create user");

        // creating an existing user
        let resp = test::TestRequest::post()
            .uri("/user")
            .set_json(&req_body)
            .send_request(&mut app)
            .await;
        assert!(
            resp.status().is_client_error(),
            "Should not be possible to create user with same username"
        );

        // finding a user
        let resp = test::TestRequest::get()
            .uri(&format!("/user/{}", test_username))
            .send_request(&mut app)
            .await;
        assert!(resp.status().is_success(), "Failed to find user");

        // parse the result as a user
        let test_user: User = test::read_body_json(resp).await;
        assert_eq!(test_user.username, test_username, "Found wrong user");

        // deleting a user
        let resp = test::TestRequest::delete()
            .uri(&format!("/user/{}", test_user.username))
            .send_request(&mut app)
            .await;
        assert!(resp.status().is_success(), "Failed to delete user");
    }
}
