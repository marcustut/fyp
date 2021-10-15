use actix_web::{get, post, web, HttpResponse, Responder};
use async_graphql::http::{playground_source, GraphQLPlaygroundConfig};
use async_graphql_actix_web::Request;

use crate::graphql::schema::AppSchema;

#[post("/graphql")]
async fn graphql(schema: web::Data<AppSchema>, request: Request) -> impl Responder {
    let schema = schema.into_inner();
    let res = schema.execute(request.into_inner()).await;
    HttpResponse::Ok().json(res)
}

#[get("/playground")]
async fn playground() -> impl Responder {
    HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(playground_source(GraphQLPlaygroundConfig::new("/graphql")))
}
