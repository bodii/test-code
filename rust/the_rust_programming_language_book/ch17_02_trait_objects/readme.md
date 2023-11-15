### 为使用不同类型的值而设计的trait对象

#### 定义通用作为的trait
trait对象并不像其他语言中的对象那么通用：其(trait对象)具体的作用是允许对通用作为进行抽象。

当编写库的时候，我们不知道何人会在何进增加SelectBox类型，不过Screen的实现能够操作并绘制这个新类型，因为SelectBox实现了Draw trait，这意味着它实现了draw方法。

这个概念 -- 只关心值所反映的信息而不是其具体类型 -- 类似于动态类型语言中称为`鸭子类型(duck typing)`的概念：如果它走起来像一只鸭子，叫起来像一只鸭子，那么它就是一只鸭子！

使用trait对象和Rust类型系统来进行类似鸭子类型操作的优势是无需在运行时检查一个值是否实现了特定方法或者担心在调用时因为值没有实现方法而产生错误。如果值没实现trait对象所需的trait则Rust不会编译这些代码。

### trait对象执行动态分发
当对泛型使用trait bound时编译器所进行单态化处理：编译器为每一个被泛型类型参数代替的具体类型生成了非泛型的函数和方法实现。单态化所产生的代码进行`静态分发(static dispatch)`。静态分发发生于编译器在编译时就知晓调用了什么方法的时候。这与`动态分发(dynamic dispatch)`相对，这时编译器在编译时无法知晓调用了什么方法。在动态发发的情况下，编译器会生成在运行时确定调用了什么方法的代码。

当使用trait对象时，Rust必须使用动态分发。编译器无法知晓所有可能用于trait对象代码的类型，所以它也不知道应该调用哪个类型的哪个方法实现。为此，Rust在运行时使用trait对象中的指针来知晓需要调用哪个方法。动态分发也阻止编译器有选择的内联方法代码，这会相应的禁用一些优化

### trait对象要求对象安全
只有`对象安全(object safe)`的trait才可以组成trait对象。围绕所有使用trait对象安全的属性存在一些复杂的规则，不过在实践中，只涉及到两条规则。如果一个trait中所有方法有如下属性时，则该trait是对象安全的：
* 返回值类型不为Self
* 方法没有任何泛型类型参数

`Self`关键字是我们要实现trait或方法的类型的别名。对象安全对于trait对象是必须的，因为一旦有了trait对象，就不再知晓实现该trait的具体类型是什么了。如果trait方法返回具体的Self类型，但是trait对象忘记了其真正的类型，那么方法不可能使用已经忘却的原始具体类型。同理对于泛型类型参数来说，当使用trait时其会放入具体的类型参数：此具体类型变成了实现该trait的类型的一部分。当使用trait对象时其具体类型被抹去了，故无从得知放入泛型参数类型的类型是什么。

一个trait的方法不是对象安全的例子是标准库中的Clone trait。Clone trait的clone方法参数签名：
```Rust
pub trait Clone {
    fn clone(&self) -> Self;
}
```

String 实现了Clone trait，当在String实例中调用clone方法时会得到一个String实例。类似的，当调用Vec<T>实例的clone方法会得到一个Vec<T>实例。clone的签名需要知道什么类型会代替Self，因为这是它的返回值。

如果尝试做一些违反有关trait对象的对象安全规则的事情，编译器会提示你：
```Rust
pub struct Screen {
    // error: the trait `std::clone::Clone` cannot be made into an object
    pub components: Vec<Box<dyn Clone>>,
}
```

