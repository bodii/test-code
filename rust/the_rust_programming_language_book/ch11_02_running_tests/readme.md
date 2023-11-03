### 控制测试如何运行
可以将一部分命令行参数传递给cargo test，而将另外一部分传递给生成的测试二进制文件。为了分隔这两种参数，需要先列出传递给cargo test的参数，接着是分隔符--，再之后是传递给测试二进制文件的参数。运行cargo test --help会提示cargo test的有关参数，而运行cargo test -- -- help 可以提示在分隔符--之后使用的有关参数。

### 并行或连续的运行测试
当运行多个测试时，Rust默认使用线程来并行运行。

可以传递--test-threads参数和希望使用线程的数量给测试二进制文件
```shell
cargo test -- --test-threads=1
```
这里将测试线程设置为1， 告诉程序不要使用任何并行机制。

### 显示函数输出
如果你希望看到通过测试中打印的值，可以通过在末尾增加--show-output参数来告知Rust显示通过测试的输出：
```shell
cargo test -- --show-output
```
### 通过指定名字来运行部分测试 
只有名称被匹配到的才会被测试到
> 运行单个测试
```Rust
cargo test one_hundred
```

> 过滤运行多个测试
```Rust
cargo test add 
```

### 忽略某些测试
对于想要排除的测试，我们在#[test]之后增加了#[ignore]行。

如果我们只希望运行被忽略的测试，可以使用
```shell
cargo test -- --ignored
```
