如果类中未定义toString方法，则程序会默认定义toString方法。这时，toString方法返回
的是下述字符串。
```java
getClass().getName() + '@' + Integer.toHexString(hashCode());
```
这是由类名、＠符、实例哈希码的无符号十六进制表示拼接起来的字符序列。
对于Day类类型，toString方法返回的是"Day@e09713"字符串（十六进制数e09713部分根据实
例不同而有所不同)。

在构造函数的开头，可以使用this()来调用类中的其他构造函数，委托其处理。
