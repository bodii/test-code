// 结构体定义中的泛型
struct Point<T> {
    x: T,
    y: T,
}

struct Point2<T, U> {
    x: T,
    y: U,
}

// 枚举定义中的泛型
enum Option2<T> {
    Some(T),
    None,
}

enum Result2<T, E> {
    Ok(T),
    Err(E),
}

// 方法定义中的泛型
struct Point3<T> {
    x: T,
    y: T,
}

impl<T> Point3<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

struct Point4<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point4<T, U> {
    fn mixup<V, W>(self, other: Point4<V, W>) -> Point4<T, W> {
        Point4 {
            x: self.x,
            y: other.y,
        }
    }
}

fn main() {
    let integer = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };

    let both_integer = Point2 { x: 5, y: 10 };
    let both_float = Point2 { x: 1.0, y: 4.0 };
    let integer_and_float = Point2 { x: 5, y: 4.0 };

    let p = Point3 { x: 5, y: 10 };
    println!("p.x = {}", p.x());

    let p1 = Point4 { x: 5, y: 10.4 };
    let p2 = Point4 { x: "Hello", y: 'c' };

    let p3 = p1.mixup(p2);
    println!("p3.x = {}, p3.y = {}", p3.x, p3.y);
}
