### 引用与借用
```Rust
fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
```
这此`&`符号就是`引用`，它们允计你使用值但不获取其所有权。

将创建一个引用的行为称为`借用(borrowing)`。
正如现实生活中，如果一个人拥有某样东西，你可以从他那里借来。当你使用完毕，必须还回去。

```Rust
fn main() {
    let mut s = String::from("hello");
    change(&mut s);
    let len = calculate_length(&s);
    println!("The length of '{}' is {}.", s, len);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```
> 不过可变引用有一个很大的限制：在同一时间，只能有一个对某一特定数据的可变引用。尝试创建两个可变引用的代码将会失败。

```Rust
let mut = String::from("hello");

let r1 = &mut s;
// error
let r2 = &mut s;

println("{}, {}", r1, r2);
```
这个报错说这段代码是无效的，因为我们不能在同一时间多次将s作为可变变量借用。第一个可变的借入在r1中，并且必须持续到在println!中使用它。

防止同一时间对同一数据进行多个可变引用的限制允许可变性，不过是以一种受限制的方式允许。
这个限制的好处是Rust可以在编译时就避免数据竞争。
`数据竞争(data race)`类似于竞态条件，它由这三个行为造成：
* 两个或更多指针同时访问同一数据。
* 至少有一个指针被用来写入数据。
* 没有同步数据访问的机制。
以上三个行为同时发生才会造成数据竞争，而不是单一行为。

可以使用大括号来创建一个新的作用域，以允许拥有多个可变引用，只是不能`同时`拥有:
```Rust
let mut s = String::from("hello");
{
  let r1 = &mut s;
} // r1在这里离开了作用域，所以我们完全可以创建一个新的引用
let r2 = &mut s;
```

类似的规则也存在于同时使用可变与不可变引用中。

这些代码会导致一个错误:
```Rust
let mut s = String::from("hello");

let r1 = &s; // 没问题
let r2 = &s; // 没问题
let r3 = &mut s; // 有问题

println("{}, {}, and {}", r1, r2, r3);
```
> 不能在拥有不可变引用的同时拥有可变引用。

* 一个引用的作用域从声明的地方开始一直持续到最后一次使用为止。
```Rust
let mut s = String::from("hello");

let r1 = &s; // 没问题
let r2 = &s; // 没问题
println!("{} and {}", r1, r2);
// 此位置之后r1和r2不再使用

let r3 = &mut s; // 没问题
println!("{}", r3);
```

### 悬垂引用(Dangling References)
在具有指针的语言中，很容易通过释放内存时保留指向它的指针而错误地生成一个`悬垂指针（Dangling pointer）`，所谓悬垂指针是其指向的内存可能已经被分配给其它持有者。相比之下，在Rust中编译器确保引用永远也不会变成悬垂状态：当你拥有一些数据的引用，编译器确保数据不会在其引用之前离开作用域。
```Rust
fn main() {
  let reference_to_nothing = dangle();
}

fn dangle() -> &String { // dangle 返回一个字符串的引用
  let s = String::from("hello"); // s是一个新字符串
  &s // 返回字符串s 的引用
} // 这里s离开作用域并被丢弃。其内存被释放。
// 危险!
```
因为s是在dangle函数内创建的，当dangle的代码执行完毕后，s将被释放。不过我们尝试返回它的引用。这意味着这个引用会指向一个无效的String。
```Rust
// 正确的方法
fn no_dangle() -> String {
  let s = String::from("hello");

  s
} // 所有权被移出，所以没有值被释放
```
### 引用的规则
* 在任意给定时间，要么只能有一个可变引用，要么只能有多个不可变引用。
* 引用必须总是有效的。
