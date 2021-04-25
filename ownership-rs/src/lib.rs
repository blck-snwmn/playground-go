#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn sample1() {
        // 値が共有されている間、変更は許可されない
        let mut v = vec![1, 2, 3, 4];
        let e = v.get(1).unwrap(); // 値を共有
        v.push(10); // 変更する

        // すでに変更をしようとしているので、
        // 以下ををuncomment するとコンパイルエラー
        // println!("{}", e);
        println!("{:?}", v);
    }
}
