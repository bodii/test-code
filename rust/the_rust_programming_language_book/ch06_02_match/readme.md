### 匹配(match)
Rust中的匹配是`穷举式(exhaustive)`的: 必须穷举到最后的可能性来使代码有效。

### 通配模式和_占位符
```Rust
let dic_roll = 9;
match dice_roll {
    3 => add_fancy_hat(),
    7 => remove_fancy_hat(),
    // 最后一个分支则涵盖了所有其他可能的值，模式是命名为other的一个变量。
    other => move_player(other),
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
fn move_player(num_spaces: u8) {}
```
Rust提供了一个模式，当不想使用通配模式获取值时，请使用`_`,这是一个特殊的模式，可以匹配任意值而不绑定到该值。
```Rust
let dic_roll = 9;
match dice_roll {
    3 => add_fancy_hat(),
    7 => remove_fancy_hat(),
    // _ 代替变量other
    _ => reroll(),
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
fn reroll() {}
```

如果不想运行任何代码：
```Rust
let dic_roll = 9;
match dice_roll {
    3 => add_fancy_hat(),
    7 => remove_fancy_hat(),
    // 这里不运行任何代码(空元组)
    _ => (),
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
```