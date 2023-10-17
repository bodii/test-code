### 切片Slice类型
slice 允计引用集合中一段连续的元素序列，而不用引用整个集合。

### 字符串slice
字符串slice(string slice) 是String中一部分值的引用：
```Rust
let s = String::from("hello world");

let hello = &s[0..5];
let world = &s[6..11];
```
可以使用一个由中括号中的[starting_index..ending_index]指定的range创建一个slice,其中starting_index是slice的第一个位置，ending_index则是slice最后一个位置的后一个值。

> 由于Rust的..range语法,如果想从索引0开始，可以不写两个点号之前的值。
```Rust
let s = String::from("hello");

let slice = &s[0..2];
let slice = &s[..2];
```
> 依此类推，如果slice包含String的最后一个字节，也可以舍弃尾部的数字。
```Rust
let s = String::from("hello");
let len = s.len();

let slice = &s[3..len];
let slice = &s[3..];
```
> 也可以同时舍弃这两个值来获取整个字符串的slice
```Rust
let s = String::from("hello");

let len = s.len();
let slice = &s[0..len];
let slice = &s[..];
```

```Rust
let s = "Hello, world!";
```
这里`s`的类型是&str: 它是一个指向二进制程序特定位置的slice。
这也就是为什么字符串字面量是不可变的；&str是一个不可变引用。

### 字符串slice作为参数
```Rust
fn first_word(s: &String) -> &str {}
// = 等同于
fn first_word(s: &str) -> &str {}
```
> 如是有一个字符串slice，可以直接传递它。如果有一个String，则可以传递整个String的Slice或对String的引用。这种灵活性利用了`deref coercions`的优势。
```Rust
fn main() {
  let my_string = String::from("hello world");
  
  // `first_word` 接受 `String`的切片，无论是部分还是全部
  let word = first_word(&my_string[0..6]);
  let word = first_word(&my_string[..]);

  // `first_word`也接受`String`的引用
  // 这等同于`String`的全部切片
  let word = first_word(&my_string);

  let my_string_literal = "hello world";
  // `first_word`接受字符串字面量的切片，无论是部分还是全部
  let word = first_word(&my_string[0..6]);
  let word = first_word(&my_string[..]);
  
  // 因为字符串字面值就是字符串slice
  let word = first_word(my_string_literal)
}
```

### 其他类型的Slice
```Rust
let a = [1, 2, 3, 4, 5];
let slice = &a[1..3];
```
这个slice的类型是&[i32]。

### 总结
所有权、借用和slice这些概念让Rust程序在编译时确保内存安全。Rust语言提供了跟其他系统编程语言相同的方式来控制你使用的内存，但拥有数据所有者在离开作用域后自动清除其数据的功能意味着你无须额外编写和调试相关控制代码。
