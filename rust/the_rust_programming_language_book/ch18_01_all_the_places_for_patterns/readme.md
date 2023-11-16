### 所有可能会用到模式的位置

#### match分支
```Rust
match VALUE {
    PATTERN => EXPRESSION,
    PATTERN => EXPRESSION,
    PATTERN => EXPRESSION,
}
```
match表达式必须是`穷尽(exhausive)`的，意为match表达式所有可能的值都必须被考虑到。一个确保覆盖每个可能值的方法是在最后一个分支使用捕获所有的模式：比如，一个匹配任何值的名称永远也不会失败，因此可以覆盖所有匹配剩下的情况。

有一个特定的模式`_`可以匹配所有情况，不过它从不绑定任何变量。这在例如希望忽略任何未指定值的情况很有用。

### if let 条件表达式
if let 可以对应一个可选的带有代码的else在if let 中的模式不匹配时运行。

if let 表达式的缺点在于其穷尽性没有为编译器所检查，而match表达式则检查了。

#### while let 条件循环

#### for 循环

#### let 语句
let语句更为正式的样子如下:
```Rust
let PATTERN = EXPRESSION;
```
```Rust
let (x, y, z) = (1, 2, 3);
```
如果希望忽略元组中一个或多个值，也可以使用`_`或`..`

#### 函数参数
```Rust
fn foo(x: i32) {

}
```
x 部分就是一个模式！类似于之前对let所做的，可以在函数参数中匹配元组。
```Rust
fn print_coordinates(&(x, y): &(i32, i32)) {
    println!("Current location: ({}, {})", x, y);
}

fn main() {
    let point = (3, 5);
    print_coordinates(&point);
}
```
