use nanoid::nanoid;
use serde::{Deserialize, Serialize};
use serde_json::json;
use worker::*;

mod utils;

fn log_request(req: &Request) {
    console_log!(
        "{} - [{}], located at: {:?}, within: {}",
        Date::now().to_string(),
        req.path(),
        req.cf().coordinates().unwrap_or_default(),
        req.cf().region().unwrap_or("unknown region".into())
    );
}

#[derive(Debug, Deserialize, Serialize)]
struct ShortenURLBody {
    url: String,
}

#[event(fetch)]
pub async fn main(req: Request, env: Env) -> Result<Response> {
    log_request(&req);

    // Optionally, get more helpful error messages written to the console in the case of a panic.
    utils::set_panic_hook();

    // Optionally, use the Router to handle matching endpoints, use ":name" placeholders, or "*name"
    // catch-alls to match on specific patterns. Alternatively, use `Router::with_data(D)` to
    // provide arbitrary data that will be accessible in each route via the `ctx.data()` method.
    let router = Router::new();

    // Add as many routes as your Worker needs! Each route will get a `Request` for handling HTTP
    // functionality and a `RouteContext` which you can use to and get route parameters and
    // Environment bindings like KV Stores, Durable Objects, Secrets, and Variables.
    router
        .get("/", |_, _| {
            Response::from_html(
                r#"
                    <!DOCTYPE html>
                    <html lang="en">
                    <head>
                        <meta charset="UTF-8">
                        <meta name="viewport" content="width=device-width, initial-scale=1.0">
                        <meta http-equiv="X-UA-Compatible" content="ie=edge">
                        <title>URL Shortener</title>
                    </head>
                    <body>
                        <div style="height: 100vh; display: flex; flex-direction: column; justify-content: center; align-items: center;">
                            <h2>URL Shortener</h2>
                            <form onsubmit="return handleSubmit(event)">
                                <input
                                type="text"
                                id="input-url"
                                name="url"
                                size="40"
                                placeholder="https://google.com"
                                required
                                />
                                <input type="submit" id="input-submit" value="Shorten" />
                            </form>

                            <p id="shortened-url"></p>
                        </div>
                        <script>
                            function handleSubmit(event) {
                                event.preventDefault();
                                fetch('/links', {
                                    method: 'POST',
                                    headers: {
                                        'Content-Type': 'application/json',
                                    },
                                    body: JSON.stringify({
                                        url: document.querySelector('#input-url').value,
                                    }),
                                    })
                                    .then((response) => {
                                        if (response.status == 200) {
                                            return response.json();
                                        } else {
                                            throw new Error('Issue saving URL');
                                        }
                                    })
                                    .then((data) => {
                                        document.querySelector('#shortened-url').innerHTML = data.shortened_url;
                                    })
                                    .catch((error) => {
                                        document.querySelector('#shortened-url').innerHTML = 'An error occured';
                                    });
                                return false;
                            }
                        </script>
                    </body>
                    </html>
                "#,
            )
        })
        .post_async("/links", |mut req, ctx| async move {
            let kv = ctx.kv("SHORTEN")?;

            let request: Result<ShortenURLBody> = req.json().await;
            match request {
                Ok(body) => {
                    let parsed_url = Url::parse(body.url.as_str());
                    match parsed_url {
                        Ok(url) => {
                            let slug = nanoid!(10);

                            kv.put(slug.as_str(), url.as_str())?
                                .expiration_ttl(86400)
                                .execute()
                                .await?;

                            let shortened_url =
                                format!("{}/{}", req.url()?.origin().ascii_serialization(), slug);

                            Response::from_json(
                                &json!({ "slug": slug, "shortened_url": shortened_url }),
                            )
                        }
                        Err(e) => Response::error(format!("Not a valid URL: {}", e), 400),
                    }
                }
                Err(e) => Response::error(format!("Request body invalid: {}", e), 400),
            }
        })
        .get_async("/:slug", |mut _req, ctx| async move {
            if let Some(slug) = ctx.param("slug") {
                let kv = ctx.kv("SHORTEN")?;

                match kv.get(slug).await {
                    Ok(res) => {
                        let res = res.unwrap();
                        let url = Url::parse(res.as_string().as_str())?;

                        return Response::redirect_with_status(url, 301);
                    }
                    Err(_e) => return Response::error(format!("Key not found: {}", slug), 404),
                }
            }
            Response::error("Bad Request", 400)
        })
        .get("/worker-version", |_, ctx| {
            let version = ctx.var("WORKERS_RS_VERSION")?.to_string();
            Response::ok(version)
        })
        .run(req, env)
        .await
}
