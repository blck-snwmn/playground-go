use std::{rc::Rc, sync::Arc};

fn main() {
    println!("Hello, world!");
    // let x = Arc::new("a");
    let x = Rc::new("a");
    let x_w_count = Rc::weak_count(&x);
    let x_s_count = Rc::strong_count(&x);
    println!("week={}", x_w_count);
    println!("strong={}", x_s_count);

    let xx = x.clone();
    let xx_w_count = Rc::weak_count(&xx);
    let xx_s_count = Rc::strong_count(&xx);
    println!("week={}", xx_w_count); // week=0
    println!("strong={}", xx_s_count); // strong=2

    let r = Rc::downgrade(&xx);
    println!("week={}", Rc::weak_count(&xx)); // week=1
    println!("strong={}", Rc::strong_count(&xx)); // strong=2
}
