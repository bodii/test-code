#### 记录日志API
记录日志API的优点:
* 可以很容易地取消全部日志记录，或者仅仅取消某个级别的日志，而且打开和关闭这个操作也很
容易。
* 可以很简单地禁止日志记录的输出，因此，将这些日志代码留在程序中的开销很小。
* 日志记录可以被定向到不同的处理器，用于在控制台中显示，用于存储在文件中等。
* 日志记录器和处理器都可以对记录进行过滤。过滤器可以根据过滤实现器制定的标准丢弃那此无
用的记录项。
* 日志记录可以采用不同的方式格式化，例如，纯文本或XML.
* 应用程序可以使用多个日志记录器，它们使用类似包名的这种具有层次结构的名字，例如，com.
mycompany.myapp.
* 在默认情况下，日志系统的配置由配置文件控制。如果需要的话，应用程序可以替换这个配置。


#### 基本日志
要生成简单的日志记录，可以使用全局日志记录器(global logger)并调用其info方法:
```java
Logger.getGlobal().info("File->Open menu item selected");
```
如果在适当的地方(如main形如）调用：
```java
Logger.getGlobal().setLevel(Level.OFF);
// 将会取消所有的日志。
```


#### 高级日志
企业级(industrial-strength)日志。
在一个专业的应用程序中，不要将所有的日志都记录到一个全局日志记录器中，而是可以自定义日志
记录器。
可以调用getLogger方法创建或获取记录器:
```java
private static final Logger myLogger = Logger.getLogger("com.mycompany.myapp");
```

通常，有以下7个日志记录器级别:
```java
* SEVERE
* WARNING
* INFO
* CONFIG
* FINE
* FINER
* FINEST
```

在默认情况下，只记录前三个级别。也可以设置其他的级别:
```java
Logger.setLevel(Level.FINE);
```

另外，还可以使用Level.All开启所有级别的记录，或者使用Level.OFF关闭所有级别的记录。
对于所有的级别有下面几种记录方法:
```java
logger.warning(message);
logger.fine(message);
```
同时，还可以使用log方法指定级别:
```java
logger.log(Level.FINE, message);
```

> 默认的日志配置记录了INFO或更高级别的所有记录，因此，应该使用CONFIG、FINE、FINER和
> FINEST级别来记录那些有助于诊断，但对于程序员又没有太大意义的调试信息。


> 记录日志的常见用途是记录那些不可预料的异常。可以使用下面两个方法提供日志记录中包含
> 的异常描述内容。
```java
void throwing(String className, String methodName, Throwable t)
void log(Level l, String message, Throwable t)
```

典型的用法是:
```java
if ( ... ) {
	IOException exception = new IOException("...");
	logger.throwing("com.mycompany.mylib.Reader", "read", exception);
	throw exception;
}
```
还有:
```java
try {
} catch (IOException e) {
	Logger.getLogger("com.mycompany.myapp").log(Level.WARING, "Reading image", e);
}
```


#### 修改日志管理器配置
可以通过编辑配置文件来修改日志系统的各种属性。在默认情况下，配置文件存在于:
```
jre/lib/logging.properties
```
要使用另一个配置文件，就要将java.util.logging.config.file特性设置为配置文件的存储位置，并
用下列命令启动应用程序:
```shell
java -Djava.util.logging.config.file=configFile MainClass
```


#### 格式化器
ConsoleHandler类和FileHandler类可以生成文本和XML格式的日志记录。
也可以自定义格式。这需要扩展Formatter类并覆盖下面这个方法:
```java
String format(LogRecord record);
```
这个方法对记录中的部分消息进行格式化，参数替换和本地应用操作。

很多文件格式(如XML)需要在已格式化的记录的前面加上一个头部和尾部。在这个例子中，要覆盖下面
两个方法:
```java
String getHead(Handler h);
String getTail(Handler h);
```
最后，调用setFormatter方法将格式化器安装到处理器中。

下列代码确保将所有的消息记录到应用程序特定的文件中:
```java
if (System.getProperty("java.util.logging.config.class") == null
		&& System.getProperty("java.util.logging.config.file") == null
) {
	try {
		Logger.getLogger("").setLevel(Level.All);
		final int LOG_ROTATION_COUNT = 10;
		Handler handler = new FileHandler("%h/myapp.log", 0, LOG_ROTATION_COUNT);
		Logger.getLogger("").addHandler(handler);
	} catch (IOException e) {
		logger.log(Level.SEVERE, "Can't create log file handler", e);
	}
}
```
