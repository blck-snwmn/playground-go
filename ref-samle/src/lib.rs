use std::{
    borrow::Borrow,
    cell::{Ref, RefCell},
};

struct Sample {
    innner: RefCell<Test>,
}

impl Sample {
    fn new() -> Self {
        Self {
            innner: RefCell::new(Test::new()),
        }
    }

    fn return_inner<'a>(&'a self) -> Ref<'a, String> {
        // let x = self.innner. ;
        // x.return_inner()
        // let x = Ref::map(self.innner.borrow(), |x| x.return_inner().borrow());
    }
}

struct Test {
    innner: RefCell<String>,
}

impl Test {
    fn new() -> Self {
        Self {
            innner: RefCell::new("hello world".to_string()),
        }
    }
    fn return_inner<'a>(&'a self) -> Ref<'a, String> {
        self.innner.borrow()
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {}
}
