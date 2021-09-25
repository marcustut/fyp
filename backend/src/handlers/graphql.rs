use std::sync::Arc;
use actix_web::{delete, get, post, web, Error, HttpResponse, Responder};
use juniper::http::graphiql::graphiql_source;
use juniper::http::GraphQLRequest;
use crate::models::user::Schema;

#[get("/graphiql")]
pub(crate) async fn graphiql() -> impl Responder {
    let html = graphiql_source("http://localhost:4000/graphql", None);
    HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(html)
}
#[post("/graphql")]
pub(crate) async fn graphql(st: web::Data<Arc<Schema>>, data: web::Json<GraphQLRequest>) -> impl Responder {
    let user = web::block(move || {
       let res = data.execute_sync(&st, &());
        Ok::<_, serde_json::error::Error>(serde_json::to_string(&res)?)
    }).await?;
    Ok(HttpResponse::Ok()
        .content_type("application/json")
        .body(user))
}