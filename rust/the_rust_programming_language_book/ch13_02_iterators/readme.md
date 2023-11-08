### 使用迭代器处理无素序列
迭代器模式允许对一个序列的项进行某些处理。
`迭代器(iterator)`负责遍历序列中的每一项和决定序列何时结束的逻辑。当使用迭代器时，我们无需重新实现这些逻辑。

在Rust中，迭代器是`惰性的(lazy)`，这意味着在调用方法使用迭代器之前它都不会有效果。

创建一个迭代器:
```Rust
let v1 = vec![1, 2, 3];

let v1_iter = v1.iter();

for val in v1_iter {
    println!("Got: {}", val);
}
```

### Iterator trait 和next方法
迭代器都实现了一个叫做Iterator的定义标准库的trait:
```Rust
pub trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self::Item>;
}
```
`type Item` 和 `Self::Item`，他们定义了trait的`关联类型(associated type)`。
Item类型将是迭代器返回元素的类型。
next是Iterator实现者被要求定义的唯一方法。next一次返回迭代器的一个项，封装在Some中，当迭代器结束时，它返回None。

```Rust
#[test]
fn iterator_demonstration() {
    let v1 = vec![1, 2, 3];

    let mut v1_iter = v1.iter();

    assert_eq!(v1_iter.next(), Some(&1));
    assert_eq!(v1_iter.next(), Some(&2));
    assert_eq!(v1_iter.next(), Some(&3));
    assert_eq!(v1_iter.next(), None);
}
```
> 注意v1_iter需要是可变的：在迭代器上调用next方法改变了迭代器中用来记录序列位置的状态。
换句话说，代码`消费（consume)`了，或使用了迭代器。
另外需要注意到从next调用中得到的值是vector的不可变引用。iter方法生成一个不可变引用的迭代器。
如果希望迭代可变引用，则可以调用iter_mut而不是iter.
如果获取v1所有权并返回拥有所有权的迭代器，则可以调用into_iter 而不是iter.

### 消费迭代器的方法
一些方法在其定义中调用了next方法，这就是为什么在实现Iterator trait时要求实现next方法。
这些调用next方法的方法被称为`消费适配器(consuming adaptors)`，因为调用他们会消耗迭代器，因而会消费迭代器。当其遍历每一个项时，它将每一个项加总到一个总和并在迭代完成时返回总和。
```Rust
#[test]
fn iterator_sum() {
    let v1 = vec![1, 2, 3];

    let v1_iter = v1.iter();

    let total: i32 = v1_iter.sum();

    assert_eq!(total, 6);
}
```

### 产生其他迭代器的方法
Iterator trait中定义了另一类方法，被称为`迭代器适配器(iterator adaptors)`，他们允许将当前迭代器变为不同类型的迭代器。
可以链式调用多个迭代器适配器。不过因为所有的迭代器都是惰性的，必须调用一个消费适配器方法以便获取迭代器调用的结果。
```Rust
let v1: Vec<i32> = vec![1, 2, 3];
// warning: 迭代器适配器是惰性的，而这里需要消费迭代器
v1.iter().map(|x| x + 1);
```

```Rust
let v1: Vec<i32> = vec![1, 2, 3];

// 调用map方法创建一个新迭代器，接着调用collect方法消费迭代器并创建一个vector
let v2: Vec<_> = v1.iter().map(|x| x + 1).collect();

assert_eq!(v2, vec![2, 3, 4]);
```

### 使用闭包获取环境
