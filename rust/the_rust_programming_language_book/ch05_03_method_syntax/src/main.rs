#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // 如果想要在方法中改变调用方法的实例，需要将第一个参数改为`&mut self`
    // 直接使用self（不加&）获取实例的所有权，通常用在当方法将self转换成别
    // 的实例的时候，这时我们想要防止调用者在转换后使用原始的实例
    fn area(&self) -> u32 {
        self.width * self.height
    }

    // !! 可以选择将方法的名称和结构中的一个字段相同
    fn width(&self) -> bool {
        self.width > 0
    }
    
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

// 一个结构体多个impl块
impl Rectangle {
    fn square(size: u32) -> Rectangle {
        Rectangle {
            width: size,
            height: size,
        }
    }
}

fn main() {
    let rect1 = Rectangle{
        width: 30,
        height: 50,
    };

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );

    if rect1.width() {
        println!(
            "The rectangle has a nonzero width; it is {}",
            rect1.width
        );
    }

    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect2 hold rect3? {}", rect2.can_hold(&rect3));

    // 关联函数调用
    let sq = Rectangle::square(3);
    println!("create a square {:?}", sq);
    // 方法调用
    let rect4 = sq.area();
    println!(
        "The area of the rectangle4 is {} square pixels.",
        rect4
    );
}
