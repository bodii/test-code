### 将crate发布到Crates.io

#### 编写有用的文档注释
除了`//`注释Rust代码，Rust也有特定的用于文档的注释类型，通常被称为文档`注释(documentation comments)`，他们会生成HTML文档。

文档注释使用三斜杠`///`，以支付Markdown标记来格式化文本。
文件名：src/lib.rs
```Rust
/// Adds one to the number given.
/// 
/// # Example
///
/// ```
/// let arg = 5;
/// let answer = my_crate::add_one(arg);
/// 
/// assert_eq!(6, answer);
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}
```

#### 常用（文档注释）部分
* Example: 一个以"Example"为标题的部分
* Panics: 这个函数可能会panic!的场景。并不希望程序崩溃的函数调用者应该确保他们不会在这些情况下调用此函数。
* Errors: 如果这个函数返回Result，此部分描述可能会出现何种错误以及什么情况会造成这些错误，这有助于调用者编写代码来采用不同的方式处理不同的错误。
* Safety: 如果这个函数使用unsafe代码，这一部分应该会涉及到期望函数调用者支持的确保unsafe块中代码正常工作的不变条件(invarints)。

#### 文档注释作为测试
在文档注释中增加示例代码块是一个清楚的表明如何使用库的方法，这么做还有一个额外的好处: cargo test也会像测试那样运行文档中的示例代码！没有什么比有例子的文档更好的了，但更糟糕的莫过于写完文档后改动了代码，而导致例子不能正常工作。

### 注释包含项的结构
还有另一种风格的文档注释，//!，为包含的注释的项添加文档，而不是为注释之后的项增加文档。通常用于crate根文件(通常是scr/lib.rs)或模块的根文件，为crate或模块整体提供文档。
文件名：src/lib.rs
```Rust
//! # My Crate
//!
//! `my_crate` is a collection of utilities to make performing certain
//! calculations more convenient.

/// Adds one to the number give.
```
注意//!的最后一行之后没有任何代码。因为以`//!`开头而不是`///`，这是属于包含此注释的项而不是注释之后项的文档。在这个情况中，包含这个注释的项是src/lib.rs文件，也就是crate根文件。这些注释描述了整个crate。

### 使用pub use导出合适的公有API
可以选择使用pub use `重导出(re-export)`项来使公有结构不同于私有结构。重导出获取位于一个位置的公有项并将其公开到另一个位置，好像它就定义在这个新位置一样。

创建一个有用的公有API结构更像是一门艺术而非科学，你可以反复检视他们来找出最适合用户的API。pub use提供了解耦组织crate内部结构和与终端用户体现的灵活性。


### 发布新crate之前
有了账号之后，比如说你已经有一个希望发布的crate。在发布之前，你需要在crate的Cargo.toml文件的[package]部分增加一些本crate的元信息(metadata)。
文件名： Cargo.toml
```Toml
[Package]
name = "guessing_game"
```
即使你选择了一个唯一的名称，如果此时尝试运行`cargo publish`发布该crate的话，会得到一个警告接着是一个错误:
这是因为我们缺少一些关键信息：关于该crate用途的描述和用户可能在何种条款下使用该crate的license。
```Toml
[package]
name = "guessing_game"
license = "MIT"
```
如果你希望使用不存在于SPDX的license，则需要将license文本放入一个文件，将该文件包含进项目中，接着使用license-file来指定文件名而不是使用license字段
```Toml
[package]
name = "guessing_game"
version = "0.1.0"
authors = ["Your Name <you@example.com>"]
edition = "2018"
description = "A fun game where you guess what number the computer has chosen."
license = "MIT OR Apache-2.0"

[dependencies]

```

### 发布到Crates.io
crates.io的一个主要目标是作为一个存储代码的永久文档服务器，这样所有依赖crates.io中的crate的项目都能一直正常工作。而允许删除版本没办法达成这个目标。

### 发布现存crate的新版本
改变Cargo.toml中version所指定的值。

### 使用cargo yank从Crates.io撤回版本
虽然不能删除之前版本的crate，但是可以阻止任何将来的项目将他们加入到依赖中。
运行cargo yank并指定希望撤加的版本:
```Shell
cargo yank --vers 1.0.1
```
也可以撤销撤回操作
```Shell
cargo yank --vers 1.0.1 --undo
```


