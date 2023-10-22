use rand::Rng;
use std::collections::HashMap;
// use std::{cmp::Ordering, io};

// 嵌套路径
// use std::io;
// use std::io::Write;
// 相同
// use std::io::{self, Write};

// 通过glob运算符将所有公有定义引入作用域
// 可以指定路径后跟*,glob运算符
// glob运算符经常用于测试模块tests中，这时会将所有内容引入作用域
// use std::collections::*;

fn main() {
    let mut map = HashMap::new();
    map.insert(1, 2);

    let secret_number = rand::thread_rng().gen_range(1..=100);
    println!("{}", secret_number);
}
