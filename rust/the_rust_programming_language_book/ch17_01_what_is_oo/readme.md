### 面向对象语言的特征
> 面向对象的程序是由对象组成。一个对象包含数据和操作这些数据的过程。这些过程通常被称为方法或操作。

在这个定义下，Rust是面向对象的： 结构体和枚举包含数据而impl块提供了在结构体和枚举之上的方法。虽然带有方法的结构体和枚举并不被称为对象，但是他们提供了与对象相同的功能。

### 封装隐藏了实现细节
封装(encapsulation)的思想：对象的实现细节不能被使用对象的代码获取到。所以唯一与对象交互的方式是通过对象提供的公有API；使用对象的代码无法深入到对象内部并直接改变数据或者行为。封装使得改变和重构对象的内部时无需改变使用对象的代码。
可以使用pub关键字决定模块、类型、函数和方法是公有的，而默认情况下其他一切都是私有的。
```Rust
pub struct AveragedCollection {
    list: Vec<i32>,
    average: f64,
}
```

### 继承，作为类型系统与代码共享
继承（inheritance)是一个很多编程语言都提供的机制，一个对象可以定义为继承另一个对象的定义，这使其可以获得父对象的数据和行为，而无需重新定义。

如果一个语言必须有继承才能被称为面向对象语言的话，那么Rust就不是面向对象。

Rust选择了一个不同的途径，使用trait对象而不是继承。