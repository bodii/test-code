fn main() {
    let v = vec![1, 2, 3, 4, 5];

    // panic
    // let does_not_exist = &v[100];
    // 返回None
    // warning
    let does_not_exist = v.get(100);

    let mut v = vec![1, 2, 3, 4, 5];
    // error, 当引用后，如果要修改v，则当前引用的v就会被内存释放，因而引发引用错误
    // let first = &v[0];
    v.push(6);
    // ok 
    let first = &v[0];
    println!("The first element is: {}", first);

    let v = vec![100, 32, 57];
    for i in &v {
        println!("{}", i);
    }

    // 遍历可变的vector
    let mut v = vec![100, 32, 57];
    for i in &mut v {
        *i += 50;
    }
    let first = &v[0];
    println!("first: {}", first);

    // 例如，假如我们想要从电子表格的一行中获取值，而这一行的有些列包含数字，
    // 有些包含浮点值，还有些是字符串。我们可以定义一个枚举，其成员会存放这些不同的值，
    // 同时所有这些枚举成员都会被当作相同类型，那个枚举的类型。
    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];
}
