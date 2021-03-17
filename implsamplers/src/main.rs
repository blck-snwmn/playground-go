use std::fmt::Display;

trait Bar {
    fn zzzzz(&self) -> String;
}

impl<T> Bar for T {
    fn zzzzz(&self) -> String {
        "zzzzzz".to_string()
    }
}

trait BazSelf {
    fn to_self(&self) -> Self;
}

impl<T: Display> BazSelf for &T {
    fn to_self(&self) -> Self {
        self
    }
}
struct Bob;

impl Display for Bob {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "hello, bob!")
    }
}

fn main() {
    println!("Hello, world!");
    let bob = &Bob {};
    println!("{}", bob.zzzzz());
    let b = bob.to_self();
    println!("{}", bob);
    println!("{}", bob.to_self());
    println!("{}", b);
}
