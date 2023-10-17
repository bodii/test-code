fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);
    println!("The length of '{}' is {}.", s1, len);

    let mut s = String::from("hello");
    change(&mut s);
    let len = calculate_length(&s);
    println!("The length of '{}' is {}.", s, len);
}

fn calculate_length(s: &String) -> usize { // s是对String的引用
    s.len()
} // 这里，s离开了作用域。但因为它并不拥有引用值的所有权，
// 所以什么也不会发生

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
