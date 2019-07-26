#### String型变量
```java
String s = "ABC";
```

String型并不是基本类型(int、double等内置类型),而是java.lang包中的类.
> String型之所以无需导入便可通过简名来使用，是因为它是java.lang包中的类。

字符串常量是String型实例的引用。
不管是字符串常量，还是显式声明的变量，这些字符串都是String型。

#### 空引用和空字符串的区别
如果:
```java
String s1 = null;
s1.length(); // 运行时错误 java.lang.NullPointerException
```

#### 引用其他的字符串
```java
String s1 = "ABC";  // 变量s1初始化为"ABC"的引用。
```


#### String构造函数
由于String型为类类型，因此可以显式调用构造函数来创建实例。
