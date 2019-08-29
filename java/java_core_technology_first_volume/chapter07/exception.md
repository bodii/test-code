#### Throwable
> 所有的异常都是由Throwable继承而来，但在下一层立即分解为两个分支: `Error` 和`Exception`。

Error类层次结构描述了Java运行时系统的内部错误和资源耗尽错误。应用程序不应该抛出这种类型的
对象。如出现了这样的内部错误，除了通告给用户，并尽力使用程序安全地终止之外，再也无能为力了。

在设计Java程序时，需要关注Exception层次结构。这个层次结构又分解为两个分支: 一个分支派生于
RuntimeException; 另一个分支包含其他异常。划分两个分支的规则是:由程序错误导致的异常属于
RuntimeException;而程序本身没有问题，但由于像I/O错误这类问题导致的异常属于其他异常。

派生于RuntimeException的异常包含下面几种情况:
* 错误的类型转换
* 数组访问越界
* 访问null指针

不是派生于RuntimeException的异常包括:
* 试图在文件尾部后面读取数据
* 试图打开一个不存在的文件
* 试图根据给定的字符串查找Class对象，而这个字符串表示的类并不存在

"如果出现RuntimeException异常，那么就一定是你的问题"是一条相当有道理的规则。
应该通过检测数组下标是否越界来避免ArrayIndexOutOfBoundsException异常;
应该通过在使用变量之前检测是否为null来杜绝NullPointerException异常的发生。

Java语言规范将派生于Error类或RuntimeException类的所有异常称为`非受查(unchecked)异常`，
所有其他的异常称为`受查(checked)异常`。

未预期的EOF后的信号:
```java
throw new EOFException();

// 或者

EOFException e = new EOFexception();
throw e;
```


