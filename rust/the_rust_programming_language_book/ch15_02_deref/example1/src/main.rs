use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    // 定义了用于此trait的关联类型
    type Target = T;

    fn deref(&self) -> &T {
        &self.0
    }
}

fn main() {
    let x = 5;
    let y1 = &x;
    let y2 = Box::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *y1);
    assert_eq!(5, *y2);


    let x = 5;
    let y = MyBox::new(x);

    assert_eq!(5, x);
    // error: cannot be dereferenced
    // assert_eq!(5, *y);
    assert_eq!(5, *(y.deref()));

}
