### 测试的组织结构

#### 单元测试
单元测试与他们要测试的代码共同存放在位于scr目录下相同的文件中。规范是在每个文件中创建包含测试函数的tests模块，并使用cfg(test)标注模块。

#### 测试模块和#[cfg(test)]
测试模块的#[cfg(test)]标注告诉Rust只在执行cargo test时才编译和运行测试代码，而在运行cargo build时不这么做。

> cfg属性代表`configuration`，它告诉Rust其之后的项只应该被包含进特定配置选项中。

#### 测试私有函数
可以在测试中导入私有函数

### 集成测试
为了创建集成测试，你需要先创建一个tests目录，与src同级。

可以通过指定测试函数的名称作为cargo test的参数来运行特定集成测试。也可以使用cargo test 的`--test`后跟文件的名称来运行某个特定集成测试文件中的所有测试
```shell
cargo test --test integration_test
```
tests目录中的文件都被编译为单独的crate

#### 集成测试中的子模块
例如：
```shell
tests/common/mod.rs
```

### 二进制crate的集成测试
如果项目是二进制crate并且只包含src/main.rs而没有src/lib.rs，这样就不可能在tetss目录创建集成测试并使用use语句导入src/main.rs中定义的函数。只有库crate才会向其他crate暴露了可供调用和使用的函数；二进制crate只意在单独运行。
