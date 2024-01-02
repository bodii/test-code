trait Shape {
    fn draw(&self);
}

struct Circle;

impl Shape for Circle {
    fn draw(&self) {
        println!("Drawing a circle");
    }
}

struct Rectangle;

impl Shape for Rectangle {
    fn draw(&self) {
        println!("Drawing a rectangle");
    }
}

trait ShapeFactory {
    fn create_circle(&self) -> Box<dyn Shape>;
    fn create_rectangle(&self) -> Box<dyn Shape>;
}

struct ConcreteShapeFactory;

impl ShapeFactory for ConcreteShapeFactory {
    fn create_circle(&self) -> Box<dyn Shape> {
        Box::new(Circle)
    }

    fn create_rectangle(&self) -> Box<dyn Shape> {
        Box::new(Rectangle)
    }
}

fn main() {
    let factory = ConcreteShapeFactory;
    let circle = factory.create_circle();
    let rectangle = factory.create_rectangle();

    circle.draw();
    rectangle.draw();
}
