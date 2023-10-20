### 包和crate
crate是一个二进制项或者库
crate root是一个源文件，Rust编译器以它为起始点，并构成你的crate的根模块。
包(package)是提供一系列功能的一个或者多个crate。`一个包会包含有一个Cargo.toml文件`，阐述如何去构建这些crate。

`包`中所包含的内容由几条规则来确立：
一个包中至多只能包含一个库crate(library crate);
包中可以包含任意多个二进制crate（binary crate）；
包中至少包含一个crate，无论是库的还是二进制的。

> 如果一个包同时含有src/main.rs和src/lib.rs，则它有两个crate: 一个库和一个二进制项，且名字都与包相同。通过将文件放在src/bin目录下，一个包可以拥有多个二进制crate: 每个src/bin下的文件都会被编译成一个独立的二进制crate。