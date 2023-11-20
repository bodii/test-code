### 高级类型

#### 为了类型安全和抽象而使用newtype模式
newtype模式可以用于一些其他功能，包括静态的确保某值不被混淆，和用来表示一个值的单元。
另一个newtype模式的应用在于抽象掉一些类型的实现细节：例如，封装类型可以暴露出与直接使用其内部私有类型时所不同的所有API，以便限制其功能。
newtype也可以隐藏其内部的泛型类型。

#### 类型别名用来创建类型同义词
连同newtype模式，Rust还提供了声明`类型别名(type alias)`的能力，使用type关键字来给予现有类型另一个名字。
```Rust
type Kilometers = i32;
```
这意味着Kilometers是i32的`同义词(synonym)`;
```Rust
type Kilometers = i32;

let x: i32 = 5;
let y: Kilometers = 5;

println!("x + y = {}", x + y);
```
因为Kilometers是i32的别名，他们是同一类型，可以将i32与Kilometers相加，也可以将Kilometers传递给获取i32参数的函数。

类型别名的主要用途是减少重复。例如，可能会有这样很长的类型：
```Rust
Box<dyn Fn() + Send + 'static>
```
在函数签名或类型标注中每次都写这个类型将是枯燥且易于出错的。
```Rust
let f: Box<dyn Fn() + Send + 'static> = Box::new(|| println!("hi"));

fn takes_long_type(f: Box<dyn Fn() + Send + 'static>) {}

fn returns_long_type() -> Box<dyn Fn() + Send + 'static> {}
```
类型别名通过减少项目中重复代码的数量来使其更加易于控制。
```Rust
type Thunk = Box<dyn Fn() + Send + 'static>;
let f: Thunk = Box::new(|| println!("hi"));

fn takes_long_type(f: Thunk) {}

fn returns_long_type() -> Thunk {}
```
类型别名也经常与Result<T, E>结合使用来减少重复。考虑一下标准库中的std::io模块。I/O操作通常会返回一个Result<T，E>，因为这些操作可能会失败。标准库中的std::io::Error结构体代表了所有可能的I/O错误。std::io中大部分函数会返回Result<T, E>，其中E是std::io::Error。

#### 从不返回的never type 
Rust有一个叫做`!`的特殊类型。在类型理论术语中，它被称为`empty type`，因为它没有值。
never type: 在函数从不返回的时候充当返回值:
```Rust
fn bar() -> ! {}
```
"函数bar从不返回"，而从不返回的函数被称为**发散函数(diverging functions)**。不能创建!类型的值，所以bar也不可能返回值。

```Rust
let guess: u32 = match guess.trim().parse() {
    Ok(num) => num,
    Err(_) => continue,
};
```
continue的值是`!`。
```Rust
impl<T> Option<T> {
    pub fn unwrap(self) -> T {
        match self {
            Some(val) => val,
            None => panic!("called `Option::unwrap()` on a `None` value"),
        }
    }
}
```
panic!是`!类型`

```Rust
print!("forever ");

loop {
    print!("and ever ");
}
```
这里循环永远也不结束，所以此表达式的值是!。如果引入break这就不为真了，因为循环在执行到break后就会终止。

#### 动态大小类型和Sized trait
因为Rust需要知道例如应该为特定类型的值分配多少空间这样的信息其类型系统的一个特定的角落可能令人迷惑: 这就是`动态大小类型(dynamically sized types)`的概念。这有时被称为"DST"或"unsized types"，这些类型允许我们处理只有在运行时才知道大小的类型。

动态大小类型的细节： str。没错，不是&str，而是str本身。str是一个DST；直到运行时都不知道字符串有多长。因为直到运行时都不知道其大小，也就意味着不能创建str类型的变量，也不能获取格拉类型的参数
```Rust
// error
let s1: str = "Hello there!";
// error
let s2: str = "How's it going?";
```
Rust需要知道应该为特定类型的值分配多少内存，同时所有同一类型的值使用相同数量的内存。
如果Rust允许编写上面这样的代码，也就意味着这两个str需要占用完全相同大小的空间，不过它们有着不同的长度。这就是为什么不能创建一个存放动态大小类型的变量的原因
正确应该是：s1和s2的类型是&str而不是str。

所以虽然&T是一个储存了T所在的内存位置的单个值，&str则是`两个值`：str的地址和其长度。
这样，&str就有了一个在编译时可以知道的大小：它是usize长度的两位。

> 动态大小类型的黄金规则: 必须将动态大小类型的值置于某种指针之后。

可以将str与所有类型的指针结合：如何Box<str>或Rc<str>。
另一个动态大小类型：trait
每一个trait都是一个可以通过trait名称来引用的动态大小类型。为了将trait用于trait对象，必须将他们放入指针之后，比如`&dyn Trait`或Box<dyn Trait> (Rc<dyn Trait> 也可以)

为了处理DST,Rust有一个特定的trait来确定一个类型的大小是否在编译时可知：这就是Sized trait。这个trait自动为编译器在编译时就知道其大小的类型实现。另外，Rust隐式的为每一个泛型函数增加了Sized bound。
```Rust
fn generic<T>(t: T) {}
// 实际上被当作如下处理：
fn generic<T: Sized>(t: T) {}
```
泛型函数默认只能用于编译时已知大小的类型。然而可以使用如下特殊语法来放宽这个限制:
```Rust
fn generic<T: ?Size>(t: &T) {}
```
?Sized trait bound 与Sized相对；它可以读作`"T可能是也可能不是Sized的"`。这个语法可能用于Sized，而不能用于其他trait。
另外注意我们将t参数的类型从T 变为&T: 因为其类型可能不是Sized的，所以需要将其置于某种指针之后
