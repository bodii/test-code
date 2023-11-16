### 模式语法
```Rust
let x = 1;

match x {
    1 => println!("one"),
    2 => println!("two"),
    3 => println!("three"),
    _ => println!("anything"),
}
```

#### 匹配命名变量
```Rust
fn main() {
    let x = Some(5);
    let y = 10;

    match x {
        Some(50) => println!("Got 50"),
        // ok
        Some(y) => println!("matched, y = {:?}", y),
        _ => println("Default case, x = {:?}", x),
    }
}
```

#### 多个模式
在match表达式中，可以使用`|`语法匹配多个模式，它代表`或(or)`的意思。

#### 通过..=匹配值的范围
```Rust
let x = 5;

match x {
    1..=5 => println!("one through five"), 
    _ => println!("something else"),
}
```
如果x 是1、2、3、4或5，第一个分支就会匹配。

#### 解构并分解值
可以使用模式来解构结构体、枚举、元组和引用，以便使用这些值的不同部分。

#### 解构枚举

#### 解构结构体和元组
```Rust
let ((feet, inches), Point {x, y}) = ((3, 10), Point { x: 3, y: -10 });
```

#### 忽略模式中值
`_`模式作为match表达式最后的分支特别有用，也可以将其用于任意模式，包括函数参数:
```Rust
fn foo(_: i32, y: i32) {
    println!("This code only uses the y parameter: {}", y);
}

fn main() {
    foo(3, 4);
}
```

#### 在名字前以一个下划线开头来忽略未使用的变量
```Rust
fn main() {
    let _x = 5;
    let y = 10;
}
```

#### 使用..忽略剩余值

#### 匹配守卫提供的额外条件
`匹配守卫(match guard)`是一个指定于match分支模式之后的额外if条件，它也必须被满足才能选择此分支。
```Rust
let num = Some(4);

match num {
    Some(x) if x < 5 => println!("less than five: {}", x),
    Some(x) => println!("{}", x),
    None => (),
}
```

#### @绑定
at运算符(@)允许我们在创建一个存放值的变量的同时测试其值是否匹配模式

