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


#### 创建异常类
```java
class FileFormatException extends IOException {
	public FileFormatException() {}
	public FileFormatException(String gripe) {
		super(gripe);
	}
}

// 现在就可以抛出自已定义的异常类型了
String readDate(BufferedReader in) throws FileFormatException {
	while(...) {
		if (ch == -1) // EOF encountered
		{
			if (n < len)
				throw new FileFormatException();
		}
	}
}
```

#### java.lang.Throwable
* Throwable()
	构造一个新的Throwable对象，这个对象没有详细的描述信息。

* Throwable(String message)
	构造一个新的throwable对象，这个对象带有特定的详细描述信息。
	习惯上，所有的派生的异常类都支持一个默认的构造器和一个带有详细描述信息的构造器

* String getMessage()
	获得Throwable对象的详细描述信息。


#### 捕获异常
```java
try {
	code 
	more code
} catch （Exception e) {
	handler for this type
}

```

下面给出一个读取数据的典型程序代码:
```java
public void read(String filename) {
	try {
		InputStream in = new FileInputStream(filename);
		int b;
		while((b = in.read()) != -1) {
			process input
		}
	} catch (IOException exception) {
		exception.printStackTrace();
	}
}
```

同一个catch子句中可以捕获多个异常类型:
```java
try {
	code that minght throw exceptions
} catch (FileNotFoundException | unknownHostExction e) {
	emergency action for missing files an unknown hosts
} catch (IOException e) {
	emergency action for all other I/O probems
}
```
只有当捕获的异常类型彼此之间不存在子类关系时才需要这个特性。

> 捕获多个异常时，异常变量隐含为final变量。例如，不能在以下子句体中为e赋不同的值:
```java
catch (FileNotFoundException | UnknownHostException e) { ... }
```

#### 再次抛出异常与异常链
```java
try {
	access the database
} catch (SQLExcetpion e) {
	throw new ServletExctption("database error: " + e.getMessage());
}
```
捕获多个异常不仅让你的代码看起来更简单，还会更高效。生成的字节码只包含一个对应公共catch
子句的代码块。

更好的处理方法，并且将原始异常设置为新异常的“原因":
```java
try {
	access database
} catch (SQLException e) {
	Throwable se = new ServleException("database error");
	se.initCause(e);
	throw se;
}
// 这里，ServleException用带有异常信息文本的构造器来构造。
```
当捕获到异常时，就可以使用下面的语句重新得到原始异常:
```java
Throwable e = se.getCause();
```
强烈建议使用这种包装技术。这样可以让用户抛出子系统中的高级异常，而不会丢失原始异常的细节。


#### finally
无论在try语句块中是否遇到异常，finally子句都会被执行。

> 这里，强烈建议解耦合try/catch和try/finally语句块。


#### 分析堆栈轨迹元素
> `堆栈轨迹(stack trace)`是一个方法调用过程的列表，它包含了程序执行过程中方法调用的特定位置。

可以调用Throwable类的printStackTrace方法访问堆栈轨迹的文本描述信息。
```java
Throwable t = new Throwable();
StringWriter out = new StringWriter();
t.printStackTrace(new PrintWriter(out));
String description = out.toString();
```
一种更灵活的方法是使用getStackTrace方法，它会得到StackTraceElement对象的一个数组，可以在你
的程序中分析这个对象数组:
```java
Throwable t = new Throwable();
StackTraceElement[] frames = t.getStackTrace();
for (StackTraceElement frame : frames)
	analyze frame
```


#### java.lang.Throwable
* Throwable(Throwable cause)
* Throwable(String message, Throwable cause)
	用给定的"原因"构造一个Throwable对象

* Throwable initCause(Throwable cause)
	将这个对象设置为"原因"。如果这个对象已经被设置为"原因",则抛出一个异常。
	返回this引用

* Throwable getCause()
	获得设为这个对象的"原因"的异常对象。如果没有设置"原因", 则返回null。

* StackTraceElement[] getStacktrace()
	获得构造这个对象时调用堆栈的跟踪。

* void addSuperessed(Throwable t)
	为这个异常增加一个"抑制"异常。这出现在带资源的try语句中，其中t是close方法抛出的
	一个异常

* Throwable[] getSuperessed()
	得到这个异常的所有"抑制"异常 。一般来说，这些是带资源的try语句中close方法抛出的异常。


#### java.lang.Exception
* Exception(Throwable cause)
* Exception(String message, Throwable cause)
	用给定的"原因"构造一个异常对象。


#### java.lang.RuntimeException
* RuntimeException(Throwable cause)
* RuntimeException(String message, Throwable cause)
	用给定的"原因"构造一个RuntimeException对象。


#### java.lang.StackTraceElement
* String getFileName()
	返回这个元素运行时对应的源文件名。如果这个信息不存在，则返回null

* int getLineNumber()
	返回这个元素运行时对应的源文件行数。如果这个信息不存在，则返回-1。

* String getClassName()
	返回这个元素运行时对应的类的完全限定名。

* String getMethodName()
	返回这个元素运行时对应的方法名。构造器名是<init>;静态初始化器名是<clinit>。

* boolean isNativeMethod()
	如果这个元素运行时在一个本地方法中，则返回true。

* String toString()
	如果存在的话，返回一个包含类名、方法名、文件名和行数的格式化字符串。


#### assert断言
```java
assert a != null;
assert i >= 0;
```
断言是一种测试和调试阶段所使用的战术性工具； 而日志记录是一种在程序的整个生命周期都可以
使用的策略性工具。


#### java.lang.Classloader
* void setDefaultAssertionStatus(boolean b)
	对于通过类加载器加载的所有类来说，如果没有显式地说明类或包的断言状态，就启用或禁用断言。

* void setClassAssertionStatus(String className, boolean b)
	对于给定的类和它的内部类，启用或禁用断言。

* void setPackageAssertionStatus(String packageName, boolean b)
	对于给定和其子包中的所有类，启用或禁用断言。

* void clearAsserionStatus()
	移去所有类和包的显式断言状态设置，并禁用所有通过这个类加载器加载的类的断言。
