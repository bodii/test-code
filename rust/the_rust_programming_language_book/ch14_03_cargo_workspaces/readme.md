### Cargo 工作空间
随着项目开发的深入，库crate持续增大，而你希望将其一步拆分成多个库crate。
Cargo提供了一个`工作空间(workspaces)`的功能。

#### 创建工作空间
`工作空间`是一系列共享同样的Cargo.lock和输出目录的包。
文件名：Cargo.toml
```toml
[worksapce]
members = [
  "adder",
]
```
```shell
cargo new adder
```
工作空间在顶级目录有一个target目录；

#### 在工作空间中创建第二个crate
文件名：Cargo.toml
```toml
[worksapce]
members = [
  "adder",
  "add-one",
]
```
新生成一个叫做add-one的库：
```shell
cargo new add-one --lib
```

让adder依赖库crate `add-one`:
文件名：adder/Cargo.toml
```toml
[dependencies]
add-one = { path = "../add-one" }
```
通过`-p`参数和包名称来运行cargo run指定工作空间中希望使用的包:
```shell
cargo run -p adder
```

#### 在工作空间中依赖外部crate
文件名：add-one/Cargo.toml
```toml
[dependencies]
rand = "0.5.5"
```
在现可以在`add-one/src/lib.rs`中增加`use rand;`
在整个工作空间中使用相同版本的外部crate可以节省可空，因为这样无需多个拷贝并确保工作作空间中crate将是相互兼容的。

#### 为工作空间增加测试
在工作作空间的根目录运行:
```shell
cargo test
```
也可以选择工作空间中特定的crate的测试，通过在根目录使用-p参数并指定希望测试的crate名称:
```shell
cargo test -p add-one
```

### 发布
如果选择向crates.io发布工作空间的crate，每一个工作空间中crate需要单独发布。
cargo publish命令没有`--all`或`-p`参数，所以必须进入每一个crate的目录并运行cargo pushlish
