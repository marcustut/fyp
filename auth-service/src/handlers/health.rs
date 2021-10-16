use crate::AppState;
use actix_web::{get, web, HttpResponse, Responder};

#[get("/health")]
pub(crate) async fn health(data: web::Data<AppState>) -> impl Responder {
    HttpResponse::Ok().body(format!("{} is healthy!", data.app_name))
}
