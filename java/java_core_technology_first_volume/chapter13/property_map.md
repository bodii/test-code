### 应用首选项的存储
应用用户通常希望能保存他们的首选项和定制信息，以后再次启动应用时再恢复这些配置。

#### 属性映射(property map)是一种存储键／值对的数据结构。属性映射通常用来存储配置信息，
它有3个特性:
* 键和值是字符串。
* 映射可以很容易地存入文件以及从文件加载。
* 有一个二级表保存默认值。
实现属性映射的Java类名为Properties。
属性映射对于指定程序的配置选项很有用:
```java
Properties settings = new Properties();
settings.setProperty("width", "200");
settings.setProperty("title", "Hello, World!");
```

可以使用store方法将属性映射列表保存到一个文件中。在这里，我们将属性映射保存在文件
program.properties中:
```java
OutputStream out = new FileOutputStream("program.properties");
settings.store(out, "Program Properties");
```
这个示例会给出以下输出:
```
#Program Properties
#Mon Apr 30 07:22:52 2007
width=200
title=Hello, World!
```
要从文件加载属性，可以使用以下调用:
```java
InputStream in = FileInputStream("program.properties");
settings.load(in);
```
