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
