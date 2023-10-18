fn main() {
    let rect1 = (30, 50);

    println!(
        "The area1 of the rectangle is {} square pixels.", 
        area_tuple(rect1)
    );

    let rect2 = Rectangle {
        width: 30,
        height: 50,
    };
    // 借用结构体而不是获取它的的所有权，
    // 这样在main函数就可以保持rect2的所有权并继续使用它
    // 所以这就是为什么在函数签名和调用的地方会有&
    println!(
        "The area2 of the rectangle is {} square pixels.", 
        area_struct(&rect2)
    );

    // error
    // println!("area2 is {}", rect2);
    //
    // ok
    println!("rect2 is {:?}", rect2);
    println!("rect2 is {:#?}", rect2);

    // 使用dbg!宏
    // dbg！宏接收一个表达式的所有权
    // 打印出代码中调用dbg!宏时所在的文件和行号，以及该表达式的结果值，
    // 并返回该值的所有权
    dbg!(&rect2);
    
    let scale = 2;
    let rect3 = Rectangle{
        // 返回表达式的值的所有权，所以width字段将获得相同的值，
        // 就像没有dbg!调用一样
        width: dbg!(30 * scale),
        height: 50,
    };
    dbg!(&rect3);

}

#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn area_tuple(dimensions: (u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}

fn area_struct(rectangle: &Rectangle) -> u32 {
    rectangle.width * rectangle.height
}
