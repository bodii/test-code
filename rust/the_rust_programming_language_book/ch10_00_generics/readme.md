### 泛型、trait和生命周期
`泛型(generics)`是具体类型或其他属性的抽象替代。

### 提取函数来减少重复
```Rust
fn main() {
  let number_list = vec![34, 50, 25, 100, 65];

  let mut largest = number_list[0];

  for number in number_list {
    if number > largest {
      largest = number;
    }
  }

  println!("The largest number is {}", largest);
}
```


