### 定义并实例化结构体
结构体的每一部分可以是不同类型，并需要命名各部分数据以便能清楚的表明其值的意义。
定义结构体，需要使用struct关键字并为整个结构体提供一个名字。
在大括号中，定义每一部分数据的名字和类型，称为`字段(field)`
```Rust
struct User {
  active: bool,
  username: String,
  email: String,
  sign_in_count: u64,
}
```
### 使用没有命名字段的元组结构来创建不同的类型
```Rust
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
  // 元组结构体的实例
  let black = Color(0, 0, 0);
  let origin = Point(0, 0, 0);
}
```

### 没有任何字段的类单元结构体
类单元结构体(unit-like struct), 因为它们类似于(), 即元组类型中的unit类型
类单元结构体用于在某个类型上实现trait但不需要在类型中存储数据的时候
```Rust
struct AlwaysEqual;

fn main() {
  let subject = AlwaysEqual;
}
```
### 结构体数据的所有权
生命周期确保结构体引用的数据有效性跟结构体本身保持一致。
```
struct User {
  active: bool,
  // error: missing lifetime specifier
  username: &str,
  // error: missing lifetime specifier
  email: &str,
  sign_in_count: u64,
}

fn main() {
  let user1 = User {
    email: "someone@example.com",
    username: "someusername123",
    active: true,
    sign_in_count: 1,
  }
}
```
