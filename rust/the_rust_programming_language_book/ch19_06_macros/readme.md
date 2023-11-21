### 宏
**宏（Macro）**指的是Rust中一系列的功能：使用macro_rules!的`声明（Declarative)宏`，和三种`过程(Procedural)宏`:
* 自定义#[derive]宏在结构体和枚举上指定通过derive属性添加的代码
* 类属性(Attribute-like)宏定义可用于任意项的自定义属性
* 类函数宏看起来像函数不过作用于作为参数传递的token

#### 宏和函数的区别
宏是一种为写其他代码而写代码的方式，即所谓的`元编程(metaprogramming)`。

元编程对于减少大量编写和维护的代码是非常有用的，它也扮演了函数扮演的角色。但宏有一些函数所没有的附加能力。

一个函数标签必须声明函数参数个数和类型。相比之下，宏能够接受不同数量的参数：同一个参数调用println!("hello")或用两个参数调用println!("hello {}", name)。而且，宏可以在编译器翻译代码前展开，例如，宏可以在一个给定类型上实现trait。而函数则不行，因为函数是在运行时被调用，同时trait需要在编译时实现。

宏定义通常要比函数定义更难阅读、理解以及维护。

宏和函数的最后一个重要的区别是：在一个文件里调用宏之前必须定义它，或将其引入作用域，而函数则可以在任何地方定义和调用。

#### 使用macro_rules!的声明宏用于通用元编程
Rust最常用的宏形式是`声明宏(declarative macros)`。它们有时也被称为"macros by example"、"macro_rules!宏"或者就是"macros"。其核心概念是，声明宏允许我们编写一些类似Rust match表达式的代码。
match表达式是控制结构，其接收一个表达式，与表达式的结果进行模式匹配，然后根据模式匹配执行相关代码。宏也将一个值和包含相关代码的模式进行比较；此种情况下，该值是传递给宏的Rust源代友字面量，模式用于和传递给宏的源代码进行比较，同时每个模式的相关代码则用于替换传递给宏的代码。所有这一切都发生于编译时。

可以使用macro_rules!来定义宏
```Rust
#[macro_export]
macro_rules! vec {
    ( $( $x:expr ),* ) => {
        {
            let mut temp_vec = Vec::new();
            $(
                temp_vec.push($x);
            )*
            temp_vec
        };

    }
}
```
#[macro_export]标注说明，只要将定义了宏的crate引入作用域，宏就应当是可用的。如果没有该标注，这个宏就是不能被引入作用域。

接着使用macro_rules!和宏名称开始宏定义，且所定义的宏并不带感叹号。名字后跟大括号表示宏定义体，在该例中宏名称是vec。

vec!宏的结构和match表达式的结构类似。此处有一个单边模式( $( $x:expr ),* )，后跟=>以及和模式相关的代码块。如果模式匹配，该相关代码块将被执行。假设这是这个宏中唯一的模式，则只有这一种有效匹配，其他任何匹配都是错误的

调用该宏时，替换该宏调用所生成的代码：
```Rust
let mut temp_vec = Vec::new();
temp_vec.push(1);
temp_vec.push(2);
temp_vec.push(3);
temp_vec
```
我们已经定义了一个宏，其可以接收任意数量和类型的参数，同时可以生成能够创建包含指定元素的vector的代码。

#### 用于从属性生成代码的过程宏
第二种形式的宏被称为`过程宏(procedural macros)`，因为它们更像函数(一种过程类型)。过程宏接收Rust代码作为输入，在这些代码上进行操作，然后产生另一些代码作为输出，而非像声明式宏那样匹配对应模式然后以另一部分代码替换当前代码。


第三种类型的过程宏(自定义派生(derive)，类属性和类函数)，不过它们的工作方式都类似。

创建过程宏时，其定义必须驻留在它们自己的具有特殊crate类型的crate中。
文件名:src/lib.rs
```Rust
use proc_macro;

#[some_attribute]
pub fn some_name(input: TokenStream) -> TokenStream {

}
```
一个使用过程宏的例子


#### 如何编写自定义derive宏

#### 类属性宏
类属性宏与自定义派生宏相似，不同于为derive属性生成代码，它们允许你创建新的属性。
```Rust
#[route(GET, "/")]
fn index() {}
```
#[route]属性将由框架本身定义为一个过程宏：
```Rust
#[proc_macro_attribute]
pub fn route(attr: TokenStream, item: TokenStream) -> TokenStream {}
```
类属性宏与自定义派生宏工作方式一致

#### 类函数宏
类函数宏定义看起来像函数调用的宏。类似于macro_rules!，它们比函数更灵活；例如，可以接受未知数量的参数。
类函数宏获取TokenStream参数，其定义使用Rust代码操纵TokenStream，就像另两种过程宏一样：
```Rust
let sql = sql!(SELECT * FROM posts WHERE id=1);
```
定义:
```Rust
#[proc_macro]
pub fn sql(input: TokenStream) -> TokenStream {}
```
