use std::fmt;

macro_rules! show {
    ($s:literal, $i:ident) => {
        println!("[{0:5}]{1:p}: {1:?}", $s, $i)
    };
}

#[derive(Debug)]
struct SampleMove {
    id: u8,
}
impl fmt::Pointer for SampleMove {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        // use `as` to convert to a `*const T`, which implements Pointer, which we can use

        let ptr = self as *const Self;
        fmt::Pointer::fmt(&ptr, f)
    }
}
#[derive(Debug, Clone, Copy)]
struct SampleCopy {
    id: u8,
}
impl fmt::Pointer for SampleCopy {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        // use `as` to convert to a `*const T`, which implements Pointer, which we can use

        let ptr = self as *const Self;
        fmt::Pointer::fmt(&ptr, f)
    }
}
fn main() {
    {
        // move semantics
        let s1 = SampleMove { id: 10 };
        show!("s1", s1);
        let s2 = s1;
        // show!(s1); // already moved!
        show!("s2", s2); // change adress
    }
    {
        // copy semantics
        let s1 = SampleCopy { id: 10 };
        show!("s1", s1);
        let s2 = s1;
        show!("s2", s2);
        show!("s1", s1); // no move. no change adress
    }
}
