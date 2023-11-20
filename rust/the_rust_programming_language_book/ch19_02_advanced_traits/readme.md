### 高级trait

#### 关联类型在trait定义中指定占位符类型
`关联类型(associated types)`是一个将类型占位符与trait相关联的方式，这样trait的方法签名中就可以使用这些占位符类型。trait的实现者会针对特定的实现在这个类型的位置指定相应的具体类型。如此可以定义一个使用多种类型的trait，直到实现比trait时都无需知道这些类型具体是什么

一个带有关联类型的trait的例子是标准库提供的Iterator trait。它有一个叫做Item的关联类型来替代遍历的值的类型。
```Rust
pub trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self:Item>;
}
```
Item是一个占位类型，同时next方法定义表明它返回Option<Self::Item>类型的值。这个trait的实现者会指定Item的具体类型，然而不管实现者指定何种类型，next方法都会返回一个包含了此具体类型值的Option。

```Rust
// 一个使用泛型的Iterator trait假想定义
pub trait Iterator<T> {
    fn next(&mut self) -> Option<T>;
}
```
使用泛型时，则不得不在每一个实现中标注类型。这是因为也可以实现为`Iterator<String> for Counter`，或任何其他类型，这样就可以有多个Counter的Iterator的实现。换句话说，当trait有泛型参数时，可以多次实现这个trait，每次需改变泛型类型的具体类型。接着当使用Counter的next方法时，必须提供类型标注来表明希望使用Iterator的哪一个实现。

通过关联类型，则无需标注类型，因为不能多次实现这个trait。只能选择一次Item会是什么类型，因为只能有一个`impl Iterator for Counter`。当调用Counter的next时不必每次指定我们需要u32值的迭代器。

#### 默认泛型类型参数和运算符重载
当使用泛型类型参数时，可以为泛型指定一个默认的具体类型。如果默认类型就足够的话，这消除了为具体类型实现trait的需要。
为泛型类型指定类型的语法是在声明泛型类型时使用:
`<PlaceholderType=ConcreteType>`。
`运算符重载(Operator overloading)`是指在特定情况下自定义运算符(如+)行为的操作。

Rust并不允许创建自定义运算符或重载任意运算符，不过std::ops中所列出的运算符和相应的trait可以通过实现运算符相关trait来重载。

add方法将两个Point实例的x值和y值分别相加来创建一个新的Point。Add trait有一个叫做Output的关联类型，它用来决定add方法的返回值类型。
```Rust
trait Add<RHS=Self> {
    type Output;

    fn add(self, rhs: RHS) -> Self::Output;
}
```
这是一个带有一个方法和一个关联类型的trait。尖括号中的`RHS=Self`: 这个语法叫做`默认类型参数(default type parameters)`。RHS是一个泛型类型参数("right hand side"的缩写)，它用于定义add方法中的rhs参数。如果实现Add trait时不的指定RHS的具体类型，RHS的类型将是默认的Self类型，也就是在其上实现Add的类型。

默认参数类型主要用于如下两个方面:
* 扩展类型而不破坏现有代码
* 在大部分用户都不需要的特定情况下进行自定义

#### 完全限定语法与消歧义：调用相同名称的方法
Rust既不以能避免一个trait与另一个trait拥有相同名称的方法，也不能阻止为同一类型同时实现这个两个trait。
不过，当调用这些同名方法时，需要告诉Rust希望使用哪一个

通常，完全限定语法定义为：
`<Type as Trait>::function(receiver_if_method, next_arg, ...);`
对于关联函数，其没有一个receiver，故只会有其他参数的列表。可以选择在任何函数或方法调用处使用完全限定语法。然而，允许省略任何Rust能够从程序中的其他信息中计算出的部分。只有当存在多个同名实现而Rust需要帮助以便知道调用哪个实现时，才需要使用这个较为冗长的语法。

#### 父trait用于在另一个trait中使用某个trait的功能
有时，可能会需要某个trait使用另一个trait的功能。在这种情况下，需要能够依赖相关的trait也被实现。这个所需的trait是实现的trait的**父(超)trait(supertrait)**。

#### newtype模式用于在外部类型上实现外部trait
**newtype模式(newtype pattern)**，它涉及到在一个元组结构体中创建一个新类型。这个元组结构体带有一个字段作为希望实现trait的类型的简单封装。接着这个封装类型对于crate是本地的，这样就可以在这个封装上实现trait。
Newtype是一个源自Haskell编程语言的概念。使用这个模式没有运行时性能消耗，这个封装类型在编译时就被省略。
例如，如果想要在Vec<T>上实现Display，而孤儿规则阻止我们直接这么做，因为Display trait和Vec<T>都定义于我们的crate之外。可以创建一个包含Vec<T>实例的Wrapper结构体，接着可以在Wrapper上实现Display并使用Vec<T>的值




