#[cfg(test)]
mod tests {
    use crate::handlers::app_config;
    use crate::models::user::User;
    use crate::AppState;
    use crate::Config;
    use crate::UserRepository;
    use actix_web::{test, web, App};
    use serde_json::json;
    use std::sync::Arc;

    #[actix_rt::test]
    async fn test_health_ok() {
        let mut app = test::init_service(
            App::new()
                .app_data(web::Data::new(AppState {
                    app_name: String::from("Auth"),
                }))
                .configure(app_config),
        )
        .await;

        let req = test::TestRequest::get().uri("/health").to_request();
        let resp = test::call_service(&mut app, req).await;
        assert!(resp.status().is_success(), "Healthcheck");
    }

    #[actix_rt::test]
    async fn test_user() {
        let config: Config = Config::from_env(false).expect("Server configuration");

        let pool = config.db_pool().await.expect("Database configuration");

        let crypto_service = config.crypto_service();

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
