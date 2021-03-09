fn main() {
    let s = Some("t");
    println!("{:?}", type_of(&s));
    let s: Result<&str, ()> = Ok("t");
    println!("{:?}", type_of(&s));
}

fn type_of<T>(_: &T) -> &str {
    std::any::type_name::<T>()
}
