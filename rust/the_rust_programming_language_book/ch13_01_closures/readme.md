### 闭包：可以捕获环境的匿名函数
Rust的闭包(closures)是可以保存进变量或作为参数传递给其他函数的匿名函数。可以在一个地方创建闭包，然后在不同的上下文中执行闭包运算。不同于函数，闭包允许捕获调用者作用域中的值。

#### 使用闭包创建行为的抽象

### 闭包类型推断和标注
闭包不要求像fn函数那样在参数和返回值上注明类型。函数中需要类型标注是因为他们是暴露给用户的显式接口的一部分。严格的定义这些接口对于保证所有人都认同函数使用和返回值的类型来说是很重要的。但是闭包并不用于这样暴露在外的接口：他们储存在变量中并被使用，不用命名他们或暴露给库的用户调用。

如果相比严格的必要性增加明确性并变得更啰嗦，可以选择增加类型标注；
```Rust
let expensive_closure = |num: u32| -> u32 {
    println!("calculating slowly...");
    thread::sleep(Duration::from_secs(2));
    num
};
```

例：
```Rust
// 一个函数定义
fn  add_one_v1   (x: u32) -> u32 { x + 1 }
// 一个完整标注的闭包定义
let add_one_v2 = |x: u32| -> u32 { x + 1 }; 
// 省略了类型标注的闭包定义
let add_one_v3 = |x|             { x + 1 };
// 去掉大括号只有一行的闭包定义
let add_one_v4 = |x|               x + 1 ;
```

闭包定义会为每个参数和返回值推断一个具体类型。
注意其定义并没有增加任何类型标注：如果尝试调用闭包两次，第一次使用String类型作为参数而第二次使用u32,则会得到一个错误:
```Rust
let example_closure = |x| x;

let s = example_closure(String::from("hello"));
// error: mismatched types
let n = example_closure(5);
```

### 使用带有泛型和fn trait的闭包
Fn系列trait由标准库提供。所有的闭包都实现了trait Fn、FnMut或FnOnce中的一个。
```Rust
struct Cacher<T>
    where T: Fn(u32) -> u32
{
    calculation: T,
    value: Option<u32>,
}
```

### 闭包会捕获其环境
闭包可以捕获其环境并访问其被定义的作用域的变量。
```Rust
let x = 4;
let equal_to_x = |z| z == x;

let y = 4;

assert!(equal_to_x(y));
```
这里，即使x并不是equal_to_x的一个参数，equal_to_x闭包也被允许使用变量x，因为它与equal_to_x定义于相同的作用域。
函数则不能做到同样的事，如果尝试如下例子，它并不能编译：
```Rust
let x = 4;
// error
fn equal_to_x(z: i32) -> bool { z == x }
let y = 4;
assert!(equal_to_x(y));
```
当闭包从环境中捕获一个值，闭包会在闭包体中储存这个值以供使用。
这会使用内存并产生额外的开销，在更一般的场景中，当我们不需要闭包来捕获环境时，我们不希望产生这些开销。因为函数从未允许捕获环境，定义和使用函数也就从不会有这些额外的开销。

闭包可以通过三种方式捕获其环境，直接对应函数的三种获取参数的方式：
> 获取所有权，可变借用和不可变借用。
这三种捕获值的方式被编码为如下三个Fn trait:
* FnOnce 消费从周围作用域捕获的变量，闭包周围的作用域被称为其`环境`, environment。为了消费捕获到的变量，闭包必须获取其所有权并在定义闭包时将其移动进闭包。其名称的Once部分代表了闭包不能多次获取相同变量的所有权的事实，所以它只能被调用一次。
* FnMut 获取可变的借用值所以可以改变其环境
* Fn 从其环境获取不可变的借用值

当创建一个闭包时，Rust根据其如何使用环境中变量来推断我们希望如何引用环境。由于所有闭包都可以被调用至少一次，所以所有闭包都实现了FnOnce。那些并没有移动被捕获变量的所有权到闭包的闭包也实现了FnMut，而不需要对被捕获的变量进行可变访问的闭包则也实现了Fn。

> 如果希望强制闭包获取其使用的环境值的所有权，可以在参数列表前使用move关键字。这个技巧在将闭包传递给新线程以便将数据移动到新线程中时最为实用。
```Rust
let x = vec![1, 2, 3];

let equal_to_x = move |z| z == x;

// error: use of moved value: `x`
println!("can't use x here: {:?}", x);

let y = vec![1, 2, 3];

assert!(equal_to_x(y));
```
