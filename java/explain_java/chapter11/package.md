### 包的定义
将数据和方法打包进行封装的是“类”，而将类集中起来进行封装的就是“包”。

### 包
包可以解决类名冲突的问题，因为通过包(package)这种逻辑“命名空间”，可以自由控制名称的适用范围。
属于包p的类Type记为p.Type,将类似p.Type这种类型的全名称为`完全限定名(fully qualified name)`，将
Type这种类名称为`简名(simple name)`。

包和类的关系就像是OS中目录(文件夹)和文件的关系。包的层次化也和目录非常相似。
(如，java包中的util包中的Scanner类)。各包名之间是通过“.”分隔的，因此该包名记为java.utila。完全
限定名为java.util.Scanner， 简名为Scanner。(java/util/Scanner, linux).


### 包的作用
1. 避免命名冲突
2. 根据特征进行分类
3. 封装(访问控制)


### 类型导入声明

> #### 单类型导入声明(single-type-import declaration)
```
import 完全限定名;
```

a. 有import声明
```java
import java.util.Scanner;

Scanner stdIn = new Scanner(System.in);
```

b. 无import声明
```java
// 如果没有进行import
java.util.Scanner stdIn = new java.util.Scanner(System.in);
```

> #### 按需类型导入声明(type-import-on-demand declaration)
```
import 包名.*;
```
“包名”所指定的包中包含的所有类型名称都可以只使用简名。

名为Date的类不只是在java.util包中，java.sql包中也有。当然，它们是名称相同的不同类。
因此，下面的程序在编译时就会发生错误:
```java
import java.util.*;
import java.sql.*;  // err

...
Date a = new Date(); // err,是java.util.Date，还是java.sql.Date ?
```

因此，不可以使用简名Date, 必须使用完全限定名java.util.Date 或java.sql.Date
```java
import java.util.*;
import java.sql.*;

...
java.util.Date a = new java.util.Date();
```

JarFile类和Manifest类位于java.util包的jar包中，这两个类不可以像下面这样进行导入。
```java
// 不能这样导入jar包
import java.util.*;  // err

// 必须指定类所在层次的包名
import java.util.jar.*;

```

#### java.lang包的自动导入
java.lang包中汇集了与Java语言密切相关的重要类，因此，该包中声明的类型名称会被自动导入。


#### 静态导入声明
```java
import static 类型名称.标识符名称;    // 单静态导入声明
import static 类型名称.* ;			  // 按需静态导入声明
```

### 包的声明
```
package 包名;
```
这个声明必须放在导入声明和类声明的前面。这是因为编译单元(translation unit)的语法结构。

example:

```java
package japan;

class Abc {
	// ...
}

class Def {
	// ...
}
```
简名为Abc,完全限定名为japan.Abc
简名为Def,完全限定名为japan.Def

> 同一个包中的类可以通过简名进行访问。
就是一个包中不可以存在同名的“包”和“类”。这与某个目录中不可以存在同名的目录和文件是一样的道理。
`包名的首字母是小写字母`


#### 包和目录
在创建包时，基本构成是将源文件和类文件放在与包名同名的目录上。
> 一个源文件中不可以声明多个public类，此外，非public类不对其他的包公开。


#### 唯一的包名
推荐使用与网址相反的排列方法。
例如，bohyoh.com的包math的名称如下
```
com.bohyoh.math
```
当域名中包含标识符不允许的连字符等特殊字符时，可以替换为下划线字符("_"), 
另外，与关键字相同的单词也会在后面加上下划线字符。
当首字符为标识符中不允许的开始字符(如数字或其他特殊字符)时，可以在该单词前面加下划线。


#### 类文件的查找和类路径
1. 引导程序类路径
这是构成Java平台的类

2. 安装类型的扩展功能
这是使用Java扩展功能结构的类。

3. 用户类路径
这是开发人员和第三方定义的、不使用扩展功能结构的类

```
为了使该程序无论保存在哪一个目录下都能运行，可以像下面这样启动java命令:
java -classpath .;c:\MyClass RandIdTester
```

对于不管在什么目录下，从任何位置都可以使用的自己创建的包，可以将其汇总在目录lib中(对于
lib,大家可以自己适当命名).各个包的类文件可以保存在目录lib中与包名同名的子目录中。
使用保存的自己创建的包的类A的运行命令如下:
> java -classpath .;lib A
