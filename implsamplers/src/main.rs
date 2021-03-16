trait Bar {
    fn zzzzz(&self) -> String;
}

impl<T> Bar for T {
    fn zzzzz(&self) -> String {
        "zzzzzz".to_string()
    }
}

struct Bob;

fn main() {
    println!("Hello, world!");
    let bob = Bob {};
    println!("{}", bob.zzzzz())
}
