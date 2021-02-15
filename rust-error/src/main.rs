use std::fs::File;

use anyhow::{anyhow, bail};
use anyhow::{Context, Result};

use thiserror::Error;

fn main() {
    println!("Hello, world!");
    match can_open() {
        Ok(_) => println!("success"),
        Err(_) => println!("failed"),
    }
    println!("========");
    if let Err(e) = can_open_outer() {
        println!("error! {:?}", e)
    }

    println!("========");
    // create error
    let e = anyhow!("error! {}:{}", 10, "test");
    println!("{:?}", e);

    println!("========");
    let e = return_error();
    println!("{:?}", e);

    println!("========");
    println!("result={:?}", to_even_outer(10));
    println!("result={:?}", to_even_outer(11));

    println!("========");
    // SomethingError::Foo に変換されている
    println!("result={:?}", switch_error_outer("foo"));
    println!("result={:?}", switch_error_outer("bar"));
    println!("result={:?}", switch_error_outer("baz"));
    println!("result={:?}", switch_error_outer("_"));
}

fn can_open_outer() -> Result<()> {
    can_open().with_context(|| format!("test"))
}

// using anyhow::Result
fn can_open() -> Result<()> {
    let _ = File::open("main.rs")?;
    Ok(())
}

fn return_error() -> Result<String> {
    bail!("error {}:{}", "a", "xx");

    // because bail is early return,
    // this expression is unreachable!!
    Ok("success".to_string())
}

#[derive(Error, Debug)]
#[error("input value isn't even. got={input}")]
struct OddError {
    input: u64,
}

fn to_even_outer(v: u64) -> Result<u64> {
    let r = to_even(v)?;
    Ok(r)
}

fn to_even(v: u64) -> std::result::Result<u64, OddError> {
    if v % 2 == 0 {
        Ok(v % 2)
    } else {
        Err(OddError { input: v })
    }
}

fn switch_error_outer(s: &str) -> Result<()> {
    switch_error(s)?;
    Ok(())
}
fn switch_error(s: &str) -> std::result::Result<(), SomethingError> {
    let e = match s {
        // #[from] atribute によって、このfromが使えている
        // #[from] がないとコンパイルエラー
        "foo" => SomethingError::from(to_even(11).unwrap_err()),
        "bar" => SomethingError::Bar(100),
        "baz" => SomethingError::Baz {
            baz1: "xxx".to_string(),
            baz2: 15,
        },
        _ => SomethingError::No,
    };
    Err(e)
}

#[derive(Error, Debug)]
enum SomethingError {
    #[error("error! foo")]
    Foo(#[from] OddError),
    // #[from] は他のVariantが同じtypeのみから生成される場合
    // コンパイルエラーになる
    // #[error("error! foo2")]
    // Foo2(#[from] OddError),
    #[error("err:{0}")]
    Bar(u32),
    #[error("{baz1}-{baz2}")]
    Baz { baz1: String, baz2: u32 },
    #[error("err no")]
    No,
}
