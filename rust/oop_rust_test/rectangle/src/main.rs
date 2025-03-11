struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn scale(&mut self, scale_factor: u32) {
        self.width *= scale_factor;
        self.height *= scale_factor;
    }
}
fn main() {
    let mut rect = Rectangle {
        width: 10,
        height: 20,
    };

    rect.scale(2);

    println!("Rectangle area = {}.", rect.area());
}
