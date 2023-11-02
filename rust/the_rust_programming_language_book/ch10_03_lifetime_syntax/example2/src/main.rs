fn main() {
    let string1 = String::from("long string is long");

    {
        let string2 = String::from("xyz");
        let result = longest(string1.as_str(), string2.as_str());
        println!("The longest string is {}", result);

        longest2(string1.as_str(), string2.as_str());

        // let result = longest3(string1.as_str(), string2.as_str());
        // println!("longest3 result: {}", result);

        let result = longest4(string1.as_str(), string2.as_str());
        println!("longest4 result: {}", result);
    }

}

fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn longest2(x: &str, y: &str) {
    println!("longest2 test: {}, {}", x, y);
}

fn _longest3<'a>(x: &'a str, y: &str) -> &'a str {
    // 悬垂引用
    // let result = String::from("really long string");
    // result.as_str()


    if x.len() > y.len() {
        println!("x > y");
    } else {
        println!("x < y");
    }

    x
}

fn longest4<'a>(x: &'a str, y: &'a str) -> String {
    let mut result = String::from("really long string");
    result.push_str(x);
    result.push_str(y);
    result
}
