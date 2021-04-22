use std::ops::Deref;

struct Foo {}

impl<'a> Foo {
    fn greet(&self) -> String {
        "hello deref sample".to_string()
    }

    fn gen() -> &'a Bar {
        &Bar { inner: Foo {} }
    }
}

struct Bar {
    inner: Foo,
}

impl Deref for Bar {
    type Target = Foo;

    fn deref(&self) -> &Self::Target {
        &self.inner
    }
}
struct Baz {
    inner: Bar,
}
// これ有効なんだ。。。
impl<'a> Baz {
    fn hello() -> &'a str {
        "hello baz"
    }
}

impl Deref for Baz {
    type Target = Bar;

    fn deref(&self) -> &Self::Target {
        &self.inner
    }
}

fn main() {
    println!("Hello, world!");
    {
        let b = &Baz {
            inner: Bar { inner: Foo {} },
        };
        // Deref を実装しているので、Fooに定義しているgreetが呼び出せる
        println!("{}", b.greet());
        // もちろん型強制できる
        let b: &Bar = b;
        let b: &Foo = b;
        println!("{}", b.greet());
    }

    {
        let b2 = &Baz {
            inner: Bar { inner: Foo {} },
        };
        // Baz -> Bar -> Foo と推移的に
        let b2: &Foo = b2;
        println!("{}", b2.greet());
    }
    {
        println!("{}", Baz::hello())
    }
}
