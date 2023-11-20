trait Animal {
    fn baby_name() -> String;
}

struct Dog;

impl Dog {
    fn baby_name() -> String {
        String::from("Spot")
    }
}

impl Animal for Dog {
    fn baby_name() -> String {
        String::from("puppy")
    }
}

fn main() {
    println!("A baby dog is called a {}", Dog::baby_name());

    // error: Animal::baby_name是关联函数而不是方法，因此它没有self参数
    // Rust无法计算出所需的是哪一个Animal::baby_name实现
    // println!("A baby dog is called a {}", Animal::baby_name());

    // ok
    // 需要使用`完全限定语法`
    println!("A baby dog is called a {}", <Dog as Animal>::baby_name());
}
