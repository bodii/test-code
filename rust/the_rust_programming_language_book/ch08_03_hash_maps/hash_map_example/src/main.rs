use std::collections::HashMap;

fn main() {
    // 新建一个哈希
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    println!("scores: {:?}", scores);

    // 使用vector的collect方法，创建同样效果的HashMap, 其中每个元组包含一个键值对
    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];
    let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();

    // 哈希map和所有权
    let field_name = String::from("Favorite color");
    let field_value = String::from("Blue");

    let mut map = HashMap::new();
    map.insert(field_name, field_value);
    // 这里field_name和field_value不再有效，
    // 学试使用它们会出现错误
    // error
    // println!("key: {}, value: {}", field_name, field_value);

    // 访问哈希map中的值
    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    let team_name = String::from("Blue");
    let score = scores.get(&team_name);
    println!("score: {:?}", score);

    // 遍历
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }

    // 更新哈希map
    // 覆盖一个值
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Blue"), 50);
    println!("{:?}", scores);

    // 只在键没有对应值时插入
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50);
    println!("{:?}", scores);

    // 根据旧值更新一个值
    let text = "hello world wonderful world";
    let mut map = HashMap::new();
    for word in text.split_whitespace() {
        // or_insert方法会返回这个键的值的一个可变引用(&mut v)
        let count = map.entry(word).or_insert(0);
        *count += 1;
        println!("word: {}, count: {}", word, &count);
    }
    println!("{:?}", map);

}
