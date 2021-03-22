fn main() {
    println!("Hello, world!");
    let mut v = Vec::new();
    // Go と同じでcapがうまると2倍確保する
    for i in 0..101 {
        println!("{:3}: cap={:3}, len={:3}", i, v.capacity(), v.len());
        v.push(i)
    }
}
