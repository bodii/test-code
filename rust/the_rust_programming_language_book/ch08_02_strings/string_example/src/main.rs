fn main() {
    let mut s1 = String::from("foo");
    let s2 = "bar";
    s1.push_str(s2);
    println!("s1 is {}", s1);
    println!("s2 is {}", s2);


    let mut s = String::from("lo");
    s.push('l');
    println!("s is {}", s);

    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // 后面s1不能再继续使用了
    println!("s3 is {}", s3);


    let s1 = String::from("tic");
    let s2 = String::from("tac");
    let s3 = String::from("toe");
    // let s = s1 + "-" + &s2 + "-" + &s3;
    // or
    let s = format!("{}-{}-{}", s1, s2, s3);
    println!("{s}");

    let len = String::from("Hola").len();
    // 4 个字节数
    println!("Hola is length: {len}");

    let len = String::from("Здравствуйте").len();
    // 24 个字节数, 这里每个Unicode标量值需要2个字节存储
    println!("Здравствуйте is length: {len}");

    // 遍历字符串
    println!("遍历字符串नमस्ते，char结果：");
    for c in "नमस्ते".chars() {
        println!("{}", c);
    }

    // 使用字节的方式遍历
    println!("遍历字符串नमस्ते，字节结果：");
    for b in "नमस्ते".bytes() {
        println!("{}", b);
    }
}
