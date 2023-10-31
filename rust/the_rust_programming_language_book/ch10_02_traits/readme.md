### trait: 定义共享的行为
trait 告诉Rust编译器某个特定类型拥有可能与其他类型共享的功能。可以通过trait以一种抽象的方式定义共享的行为。可以使用trait bounds指定泛型是任何拥有特定行为的类型。

> trait 类型于其他语言中常被称为`接口(interface)`的功能，虽然有一些不同。

### 定义trait
```Rust
// str/lib.rs
pub trait Summary {
    fn summarize(&self) -> String;
}
```
trait体中可以有多个方法：一行一个方法签名且都以分号结尾。

### 为类型实现trait
标准库并不位于crate本地作用域中，这个限制是被称为`相干性(coherence)`的程序属性的一部分，或者更具体的说是`孤儿规则(orphanrule)`，其得名于不存在父类型。这条规则确保了其他编写的代码不会破坏你的代码，反之亦然。
没有这条规则的话，两个crate可以分别对相同类型实现相同的trait，而Rust将无从得知应该使用哪一个实现。

### 默认实现
有时为trait中的某些或全部方法提供默认的行为，而不是在每个类型的每个实现中都定义自己的行为是很有用的。这样当为某个特定类型实现trait时，可以选择保留或重载每个方法的默认行为。
```Rust
// src/lib.rs
// Summary trait的定义，带有一个summarize方法的默认实现
pub trait Summary {
  fn summarize(&self) -> String {
    String::from("(Read more...)")
  }
}
```
```Rust
pub trait Summary {
  fn summarize_author(&self) -> String;
  
  fn summarize(&self) -> String {
    format!("(Read more from {}...)", self.summarize_author())
  }
}
```

### trait作为参数
```Rust
pub fn notify(item: impl Summary) {
  println!("Breaking news! {}", item.summarize());
}
```

### trait bound语法
impl trait 语法适用于直观的例子，它实际上是一种较长形式语法的语法糖。我们称为trait bound:
```Rust
pub fn notify<T: Summary>(item: T) {
  println!("Breaking news! {}", item.summarize());
}
```
impl trait 很方便，适用于短小的例子。trait bound则适用于更复杂的场景。
使用impl trait的语法看起来像这样:
```Rust
pub fn notify(item1: impl Summary, item2: impl Summary) {

}
```
不过如果你希望强制它们都是相同类型，这只有在使用trait bound时才有可能:
```Rust
pub fn notify<T: Summary>(item1: T, item2: T) {

}
```

### 通过`+` 指定多个trait bound
```Rust
pub fn notify(item: impl Summary + Display) {}
```
`+`语法也适用于泛型的trait bound:
```Rust
pub fn notify<T: Summary + Display> (item: T) {}
```
通过指定这两个trait bound, notify 的函数体可以调用summarize并使用{}来格式化item。

### 通过where简化trait bound 
使用过多的trait bound 也有缺点。每个泛型有其自已的trait bound，所以有多个泛型参数的函数在名称和参数列表之间会有很长的trait bound信息，这使得函数签名难以阅读。为此，Rust有另一个在函数签名之后的where从名中指定trait bound的语法：
```Rust
fn some_function<T: Display + Clone, u: Clone + Debug>(t: T, u: U) -> i32 {}
```
还可以这样使用where从名：
```Rust
fn some_function<T, U>(t: T, u: U) -> i32
  where T: Display + Clone,
        u: Clone + Debug
{}
```
这个函数签名就显得不那么杂乱，函数名、参数列表和返回值类型都离得很近，看起来跟没有那么多trait bounds的函数很像。


### 返回实现了trait的类型
```Rust
fn returns_summarizable() -> impl Summary {
  Tweet {
    username: String::from("horse_ebooks"),
    content: String::from("of course, as you probably already know, people"),
    reply: false,
    retweet: false,
  }
}
```
返回一个只是指定了需要实现的trait的类型的能力在闭包和迭代器场景十分的有用。
闭包和迭代器创建只有编译知道的类型，或者是非常非常长的类型。impl trait 允许你简单的指定函数返回一个Iterator而无需写出实际的冗长的类型。

### 使用trait bound 有条件地实现方法
通过使用带trait bound的泛型参数的impl块，可以有条件地只为那些实现了特定trait的类型实现方法。
```Rust
use std::fmt::Display;

struct Pair<T> {
  x: T,
  y: T,
}

impl<T> Pair<T> {
  fn new(x: T, y: T) -> Self {
    Self {
      x, 
      y,
    }
  }
}

impl<T: Display + PartialOrd> Pair<T> {
  fn cmp_display(&self) {
    if self.x > self.y {
      println!("The largest member is x = {}", self.x);
    } else {
      println!("The largest member is x = {}", self.x);
    }
  }
}
```

也可以对任何实现了特定trait的类型有条件地实现trait。对任何满足特定trait bound的类型实现trait被称为`blanket implementations`, 他们被广泛的用于Rust标准库中。
例如:
```Rust
impl<T: Display> ToString for T {

}
```
因为标准库有了这些blanket implementations, 可以对任何实现了Display trait的类型调用由ToString定义的to_string方法。例如:
```Rust
let s = 3.to_string();
```

还有一种泛型，一直在使用它甚至都没有察觉它的存在，这就是`生命周期(lifetimes)`。不同于其它泛型帮助我们确保类型拥有期望的行为，生命周期则有助于确保引用在需要的时候一直有效。
