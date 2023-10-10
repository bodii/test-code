fn main() {
    println!("Hello, world!");

    another_function();
    another_function2(5);

    print_labeled_measurement(5, 'h');

    let y = {
        // Expressions
        let x = 3;
        x + 1
    };

    println!("The value of y is: {y}"); // 4

    let x = five();
    println!("The value of x is: {x}");

    let x = plus_one(5);
    println!("The value of x is: {x}");
}

fn another_function() {
    println!("Another function.");
}

fn another_function2(x: i32) {
    println!("The value of x is: {x}");
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}

// 表达式的结尾没有分号
// 如果在表达式的结尾加上分号，它就变成语句，而语句不会返回值.
fn five() -> i32 {
    5
}

fn plus_one(x: i32) -> i32 {
    x + 1
    // panic
    // x + 1;
}
