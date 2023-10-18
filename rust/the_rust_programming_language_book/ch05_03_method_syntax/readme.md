### 结构的方法
当使用object.something()调用方法时，Rust会自动为object 添加&、&mut或*以便使object与方法签名匹配。也就是说，这些代码是等价的:
```
p1.distance(&p2);
(&p1).distance(&p2);
```

### 关联函数
所有在impl块中定义的函数被称为`关联函数(associated function)`，因为它们与impl后面命名的类型相关。
可以定义不以self为第一参数的关联函数（因此不是方法），因为它们并不作用于一个结构体的实例。我们已经使用了一个这样的函数，String::from函数，它是在String类型上定义的。

关联函数经常被用作返回一个结构体新实例的构造函数。例如我们可以使用一个关联函数，它接受一个维度参数并且同时作为宽和高，这样可以更轻松的创建一个正方形Rectangle而不必指定两次同样的值:
```Rust
impl Rectangle {
  fn square(size: u32) -> Rectangle {
    Rectangle {
      width: size,
      height: size,
    }
  }
}
```

使用`结构体名和::`语法来调用这个关联函数: 比如 let sq = Rectangle::square(3);
这个方法位于结构的命名空间中： ::语法用于关联函数和模块创建的命名空间

### 多个impl块
每个结构体都允许拥有多个impl块。
```Rust
impl Rectangle {
  fn area(&self) -> u32 {
    self.width * self.height
  }
}

impl Rectangle {
  fn can_hold(&self, other: &Rectangle) -> bool {
    self.width > other.width && self.height > other.height
  }
}

```
