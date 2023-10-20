### 定义模块来控制作用域与私有性

`模块`让我们可以将一个crate中的代码进行分组，以提高可读性与重用性。模块还可以控制项的`私有性`,即项是可以被外部代码使用的(public),还是作为一个内部实现的内容，不能被外部代码使用(private)。

> 通过执行`cargo new --lib {name}`，来创建一个新的库

### 定义
```Rust
mod front_fo_house {
    mod hosting {
        fn add_to_waitlist() {}
        fn seat_at_table() {}
    }

    mod serving {
        fn take_order() {}
        fn serve_order() {}
        fn take_payment() {}
    }
}
```
模块中可以包含其他项，比如结构体、枚举、常量、trait，或者包含函数。

对应的模块树:
```
crate
 └── front_of_house
     ├── hosting
     │   ├── add_to_waitlist
     │   └── seat_at_table
     └── serving
         ├── take_order
         ├── serve_order
         └── take_payment
```
src/main.rs和src/lib.rs被称为crate根。
这两个文件中任意一个内容会构成名为crate的模块，且该模块位于crate的被称为`模块树`的模块结构的根部（“at the root of the crate's module structure”）