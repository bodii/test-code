### 通过Deref trait 将智能指针当作常规引用处理
实现Deref trait允许我们重载`解引用运算符(derference operator) ` `*`(与乘法运算符或通配符相区别)。通过这种方式实现Deref trait的智能指针可以被当作常规引用来对待，可以编写操作引用的代码并用于智能指针。

### 通过解引用运算符追踪指针的值
```Rust
fn main() {
    let x = 5;
    let y1 = &x;

    assert_eq!(5, x);
    assert_eq!(5, *y1);
}
```
变量x存放了一个i32值5。y1等于x的一个引用。可以断言x等于5。然而y1的值做出断言，必须使用*y1来解出引用所指向的值(也就是 `解引用`)。一旦解引用了y1，就可以访问y1所指向的整型值并可以与5做比较。


### 像引用一样使用Box<T>
```Rust
fn main() {
    let x = 5;
    let y1 = &x;
    let y2 = Box::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *y1);
    assert_eq!(5, *y2);
}
```
y2设置为一个指向x值的box实例，而不是指向x值的引用。在最后的断言中，可以使用解引用运算符以y2为引用相同的方式追踪box的指针。

### 自定义智能指针
从根本上说，Box<T>被定义为包含一个元素的元组结构体。
```Rust
struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}
```
这里定义了一个结构体MyBox并声明一个泛型参数T，因为可以存放任何类型的值。
MyBox是一个包含T类型元素的元组结构体。MyBox::new函数获取一个T类型的参数并返回一个存放传入值的MyBox实例。

### 通过实现Deref trait将某类型像引用一样处理
```Rust
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
    let y = MyBox::new(x);

    assert_eq!(5, x);
    assert_eq!(5, *(y.deref()));
}
```
deref方法向编译器提供了一种能力：能够获取任何实现了Deref trait的类型值，并且可以通过调用这个类型的deref方法来获取一个解引用方法已知的&引用. 

### 函数和方法的隐式解引用强制转换
`解引用强制转换(deref coercions)`是Rust在函数或方法传参上的一种便利。解引用强制转换只能工作在实现了Deref trait的类型上。解引用强制转换将一种类型(A)隐式转换为另外一种类型(B)的引用，因为A类型实现了Deref trait，并且其关联类型是B类型。比如，解引用强制转换可以将&String转换为&str，因为类型String实现了Deref trait并且关联类型是str。
```Rust
#[stable(feature = "rust1", since = "1.0.0")]
impl ops::Deref for String {
    type Target = str;

    #[inline]
    fn deref(&self) -> &str {
        unsafe { str::from_utf8_unchecked(&self.vec) }
    }
}

fn hello(name: &str) {
    println!("Hello, {}!", name);
}

fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&m);
}
```

如果Rust没有实现解引用强制转换，为了使用&MyBox<String>类型的值调用Hello:
```Rust
fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&(*m)[..]);
}
```
(*m)将MyBox<String>解引用为String。接着&和[..]获取了整个String的字符串slice来匹配hello的签名。

### 解引用强制转换如何与可变性交互
类似于使用Deref trait重载不可变引用的*运算符，Rust提供了DerefMut trait用于重载可变引用的*运算符。

Rust 在发现类型和trait的实现满足以下三种情况时会进行解引用强制转换：
* 当 T: Deref<Target=U> : 从 &T 到 &U  
* 当 T: DerefMut<Target=U> : 从 &mut T 到 &mut U
* 当 T: Deref<Target=U> ： 从&mut T 到 &U

