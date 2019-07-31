#### try语句
引发异常的操作称为`抛出(throw)`异常。

检查、捕获(catch)抛出的异常，并对其进行处理的就是try语句(try statement)

try 语句由三部分构成:
* try 语句块(try block)
* catch 子句(catch clause)
* finally 子句(finally clause)
另外，catch子句可以有多个。此外，finally子句可以省略。


#### try语句块
其结构为关键字try的后面紧跟着语句块{}。如果try语句块在执行时抛出异常，处理就会中断，
程序流程转移到catch子句，捕获异常。反之，如果没有try语句块，就会捕获异常。
当try语句块在执行过程中未抛出异常时，try语句块会执行到最后。程序流程会跳过catch子句，直接
转移到fianlly处


####catch子句
catch子句部分会捕获try语句块在执行过程中发生的异常，并执行具体的处理，也称为`异常处理器(
exception handler)`。

catch 就是捕获的意思。

关键字catch后面的括号中会声明表示捕获的异常种类的类型和类型的形参名。虽然与函数的形参声明
类似，但这里只可以声明(即可以接收) 一个形参

异常处理器的主体中会对捕获的异常执行相应的处理，当异常处理器执行完毕后，程序流程会转移到最
后一个异常处理器的下一个位置。
```java
// 从Java SE7开始的版本可以在一个异常处理器中执行多种异常的处理。各个异常用|隔开:
catch (Exp1 | Exp2 e) { /* 对Exp1和Exp2的(相同) 处理 */ }
```

#### finally 子句
无论try语句块中是否发生异常，finally子句都一定会被执行。finally子句执行的是资源的释放处理
（例如，将打开的文件进行关闭的处理)等善后操作。

"字符串 + e"的运算，是"字符串 + 类类型变量" 的运算就是"字符串 + 类类型变量.toString()"。

对捕获的异常e调用toString方法，可以得到表示异常内容的简单字符串。

表示异常内容的字符串分为"简单的"和"复杂的"两种。前者通过toString方法获取，后者则通过getMessage
方法获取。



### 异常处理
Throwable 类为Object类的子类。另外，Throwable、Error、Exception都属于java.lang包

#### Throwable类
Throwable类位于异常类的层次结构的顶端。也就是说，Java中所有的异常类都是它的下位类。
因此，存在如下规则
* 当声明catch子句中的形参时，如果指定的类型不是Throwable的下位类，就会发生编译错误。
* 当自已创建异常类时，必须将其创建为Throwable的下位类
Throwable的子类为Error类和Exception类

#### Error类
这是程序没有希望(无法)恢复的重大异常。正如其名称所示，与其说是"异常",倒不如说是"错误"更为
准确
通常情况下，程序中无需对此类进行捕获、处理，因为即使捕获，也难以甚至无法处理.

#### Exception类
这是程序有希望(可以)恢复的异常。Exception类的下位类基本上都是称为`检查异常(checked exception)`的异常。
不过，timeException类及其下位类为非检查异常(unchecked exception)。

#### Throwable类的构造函数
Throwable() 构建详细消息为null的异常对象
Throwable(String message) 构建详细消息为message的异常对象
Throwable(String message, Throwable cause) 构建详细消息为message, 原因为cause的异常对象
Throwable(Throwable cause) 构建原因为cause的异常对象，如果cause为null, 详细消息则为null,
否则为cause.toString();

#### Throwable类的主要方法
String getMessage() 返回详细消息
Throwable getLocalizedMessage() 
	返回详细消息本地化后的描述。如果在Throwable的下位类中重写本方法，则可以创建本地(地域)
	固有的消息。如果未重写，则返回与getMessage相同的字符串

Throwable getCause() 返回原因。如果原因不存在或未知，则返回null

Throwable initCause(Throwable cause)
	设置原因。本方法只可调用一次。由于构造函数的内部会自动调用本方法，因此当使用这个构造函数
	进行构建时，一次也不可以调用。

String toString() 返回将下面三个内容拼接后的简短描述的字符串。
* 对象的类名
* ": "(冒号和空格)
* 对对象调用getLocalizedMessage方法后的结果
当getLocalizedMessage返回null时，则只返回类名

void printStackTrace()
	将对象及跟踪输出到标准错误流中，输出的第一行中包含对对象调用toString方法后返回的字符串，
	其他行则表示通过fillIInStackTrace方法记录的数据

void printStackTrace(PrintStream s) 将对象及其跟踪输出到PrintStream s 中

void printStackTrace(PrintWriter s) 将对象及其跟踪输出到PrintWriter s中

Throwable fillInStackTrace()
	将当前线程的栈跟踪的当前状态的相关信息记录到对象中。不过，如果栈跟踪不可写入，则不执行任何操作

StackTraceElement[] getStackTrace() 返回元素为printStackTrace 输出的栈跟踪的各个信息的数组

void setStackTrace(StackTraceElement[] stackTrace) 
	设置由getStackTrace返回，并由printStackTrace和相关方法输出的栈跟踪元素


#### throws 子句(声明可能抛出的检查异常)
```java
static void (check) throws Exception, RuntimeException { ... }
```

#### 方法的抛出
throw语句(throw statement)用于抛出异常，其形式为"throw表达式;".
指定的表达式为异常类类型实例的引用。
另外，不可以指定Throwable的下位类之外的类(的实例的引用)




