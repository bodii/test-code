fn main() {
    let number = 3;
    if number < 5 {
        println!("condition was ture");
    } else {
        println!("condition was false");
    }

    // error 
    /* 
    if number {
        println!("number was three");
    }
    */

    if number != 0 {
        println!("number was something other than zero");
    }

    let number = 6;
    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }

    let condition = true;
    let number = if condition { 5 } else { 6 };
    println!("The value of number is: {number}");
    // error: `if` and `else` have incompatible types
    // let number = if condition { 5 } else { "six" };
    // println!("The value of number is: {number}");

    // ---------------------------
    // ---------------------------
    // 使用循环重复执行
    // Rust有三种循环: loop、while 和 for.
    let mut counter = 0;
    let result = loop {
        counter += 1;
       if counter == 10 {
           break counter * 2;
       } 
    };

    println!("The result is {result}");

    // 循环标签：在我个循环之间消除歧义
    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                // 语句将退出外层循环
                break 'counting_up
            }
            remaining -= 1;
        }
        count += 1;
    }
    println!("End count = {count}");

    // while     
    let mut number = 3;
    while number != 0 {
        println!("{number}!");
        number -= 1;
    }
    println!("LIFTOFF!!!");

    let a: [i32; 5] = [10, 20, 30, 40, 50];
    let mut index = 0;
    while index < 5 {
        println!("the value is: {}", a[index]);

        index += 1;
    }

    // for
    let a: [i32; 5] = [10, 20, 30, 40, 50];
    for element in a {
        println!("the value is: {element}");
    }

    // range rev
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
