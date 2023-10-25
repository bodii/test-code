### 什么是字符串
Rust的核心语言中只有一种字符串类型: str, 字符串slice,它通常以被借用的形式出现， &str。
字符串slice: 它们是一些储存在别处的UTF-8编码的字符串类型。

Rust标准库中还包含一系列其他字符串类型，比如OsString、OsStr、CString和CStr。相关库crate甚至会提供更多的储存字符串数据的选择。

### 新建字符串
```Rust
let mut s = String::new();
```
```Rust
let data = "initial contents";

let s = data.to_string();

// 该方法也可直接用于字符串字面量：
let s = "initial contents".to_string();

// 也可以使用String::from函数
let s = String::from("initial contents");
```

### 更新字符串
String的大小可以增加，其内容也可以改变，就像可以放入更多数据来改变Vec的内容一样。
另外，可以方便的使用 `+` 运算符或 `format!` 宏来拼接String值。

> 使用push_str和push附加字符串
#### 可以通过push_str方法来附加字符串slice，从而使String变长
```Rust
let mut s = String::from("foo");
s.push_str("bar");
// s= foobar
```
push_str方法采用字符串slice，因为我们并不需要获取参数的所有权。
```Rust
let mut s1 = String::from("foo");
let s2 = "bar";
s1.push_str(s2);
println!("s2 is {}", s2);
```
如果push_str方法获取了s2的所有权，就不能在最后一行打印出其值了。

#### push方法被定义为获取一个单独的字符作为参数，并附加到String中。
```Rust
let mut s = String::from("lo");
s.push('l');
// s = lol
```

#### 使用+运算符或format!宏拼接字符串
```Rust
let s1 = String::from("Hello, ");
let s2 = String::from("world!");
let s3 = s1 + &s2; // 注意s1被移动了，不能继续使用
```
s1在相加后不再有效的原因，和使用s2的引用的原因，与使用`+`运算符时调用的函数签名有关。
+运算符使用了add函数：
```Rust
// 看起来像这样：
fn add(self, s: &str) -> String {}
```
首先，s2使用了&，意味着我们使用第二个字符串的`引用`与第一个字符串相加。这是因为add函数的s参数：只能将&str和String相加，不能将两个String值相加。&s2的类型是&String而不是&str。
之所以能够在add调用中使用&s2是因为&String可以被`强制(coerced)`成&str。当add函数被调用时，Rust使用了一个被称为`解引用强制转换(deref coercion)`的技术，可以将其理解为它把&s2变成了&2[..]。因为add没有获取参数的所有权，所以s2在这个操作后仍然是有效的String。

其次，可以发现签名中add获取了self的所有权，因为self没有使用&。

如果想要级联多个字符串， +的行为显得笨重了:
```Rust
let s1 = String::from("tic");
let s2 = String::from("tac");
let s3 = String::from("toe");
let s = format!("{}-{}-{}", s1, s2, s3);
```

### 索引字符串
Rust的字符串不支持索引。
```Rust
let len = String::from("Hola").len();
// 4 个字节数
println!("Hola is length: {len}");

let len = String::from("Здравствуйте").len();
// 24 个字节数, 这里每个Unicode标量值需要2个字节存储
println!("Здравствуйте is length: {len}");
```
> 因此一个字符串字节值的索引并不总是对应一个有效的Unicode标量值

从Rust的角度来讲，事实上有三种相关方式可以理解字符串：
* 字节
* 标量值
* 字形簇（最接近人们眼中`字母`的概念)

最后一个Rust不允许使用索引获取String字符的原因是，索引操作预期总是需要常数时间(O(1))。但对于String不可能保证这样的性能，因为Rust必须从开头到索引位置遍历来确定有多少有效的字符。

### 字符串slice
为了更明确索引并表明你需要一个字符串slice, 相比使用[]和单个值的索引，可以使用[]和一个range来创建含特定字节的字符串slice:
```Rust
let hello = "Здравствуйте";
let s = &hello[0..4]; // Зд

// panic, 就跟访问vector中的无效索引时一样
// let s = &hello[0..1];
```

### 遍历字符串的方法
最好的选择是使用chars方法。
```Rust
for c in "नमस्ते".chars() {
  println!("{}", c);
}
```
bytes方法返回每一个原始字节，这可能会适合你的使用场景:
```Rust
for b in "नमस्ते".bytes() {
  println!("{}", b);
}
```

### 字符串并不简单
Rust 选择了以准确的方式处理String数据作为所有Rust程序的默认行为，这意味着必须更多的思考如何预先处理UTF-8数据。
