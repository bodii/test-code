fn main() {
    let s = String::from("Hello World");
    let _word = first_word(&s);

    // s.clear(); // 清空了字符串，使用其等于""
    // word 在此处的值仍然是5，
    // 但是没有更多的字符串让我们可以有效地应用数值5。word的值现在完全无效
    
    let s = String::from("hello world");
    // 字符串slice(String slice)是String中一部分值的引用
    let _hello = &s[0..5];
    let _world = &s[6..11];

    let my_string_literal = "hello world";
    let word = second_word(&my_string_literal[0..6]);
    println!("1: {}", word);

    // 这样写也可以，即不使用slice语法
    let word = second_word(my_string_literal);
    println!("2: {}", word);
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();
    
    // iter 方法返回集合中的每一个元素
    // 而enumerate包装了iter的结果，将这些元素作为元组的一部分来返回。
    // enumerate方法返回的元组，可以使用模式来解构
    // 从.iter().enumerate()中获取了集合元素的引用，所以模式中使用了&。
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}

fn second_word(s: &String) -> &str { // 函数返回一个slice
   let bytes = s.as_bytes(); 

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[..i];
        }
    }

    &s[..]
}
