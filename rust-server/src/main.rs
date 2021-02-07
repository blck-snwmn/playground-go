use tide::prelude::*;
use tide::{utils::After, utils::Before, Body, Request, Response};

#[derive(Deserialize, Serialize, Debug)]
struct Customer {
    id: u32,
    name: String,
}

#[async_std::main]
async fn main() -> tide::Result<()> {
    let mut app = tide::new();
    app.with(After(|resp: Response| async move {
        println!("{:?}", resp);
        match resp.status() {
            tide::StatusCode::NotFound => println!("{}", tide::StatusCode::NotFound),
            tide::StatusCode::BadRequest => println!("{}", tide::StatusCode::BadRequest),
            _ => println!("ok"),
        }
        Ok(resp)
    }));
    app.with(Before(|req: Request<_>| async {
        println!("{:?}", req);
        req
    }));
    app.at("/hello").all(greet);
    app.at("/parent").nest({
        let mut api = tide::new();
        api.at("/hello").get(|_| async { Ok("Hello parent\n") });
        api
    });
    app.at("/call").post(|mut req: Request<()>| async move {
        let c: Customer = req.body_json().await?;
        Ok(Body::from_json(&c)?)
    });
    app.at("/group").get(|_: Request<()>| async {
        Ok(json!({
            "group_count": 2,
            "group": [
                { "type": "admin", "name": "bbbb" },
                { "type": "costomer", "name": "aaaa" }
            ]
        }))
    });
    app.listen("127.0.0.1:8080").await?;
    Ok(())
}

async fn greet(_: Request<()>) -> tide::Result {
    Ok("Hello!\n".into())
}
