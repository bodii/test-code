trait Pilot {
    fn fly(&self);
}

trait Wizard {
    fn fly(&self);
}

struct Human;

impl Pilot for Human {
    fn fly(&self) {
        println!("This is your captain seaking.");
    }
}

impl Wizard for Human {
    fn fly(&self) {
        println!("Up!");
    }
}

impl Human {
    fn fly(&self) {
        println!("*waving arms furiously*");
    }
}

fn main() {
    // Rust调用了直接实现在Human上的fly方法
    let person = Human;
    person.fly();

    let person = Human;
    // 在方法名前指定trait名向Rust澄清了希望调用哪个fly实现
    Pilot::fly(&person);
    Wizard::fly(&person);
    person.fly();
}
