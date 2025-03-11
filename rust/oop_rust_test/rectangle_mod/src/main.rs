mod geometry {
    pub struct Rectangle {
        pub width: u32,
        pub height: u32,
    }

    impl Rectangle {
        pub fn area(&self) -> u32 {
            self.width * self.height
        }
    }
}

fn main() {
    let rect = geometry::Rectangle {
        width: 10,
        height: 20,
    };

    println!("Rectangle area = {}.", rect.area());
}
