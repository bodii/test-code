### 采用发布配置自定义构建
在Rust中`发布配置(release profiles)`是预定义的、可定制的带有不同选项的配置。
每一个配置都彼此相互独立。
Cargo有两个主要配置：运行cargo build时采用的dev配置和运行cargo build --release的release配置。

当项目的Cargo.toml文件中没有任何[profile.*]部分的时候，Cargo会对每一个配置都采用默认设置。
```toml
[profile.dev]
opt-level = 0

[profile.release]
opt-level = 3
```
opt-level设置控制Rust会对代码进行何种程序的优化。
越高的优化级别需要更多的时间编译。
