#### JAR文件
Java归档（JAR）文件是含有大量类文件的目录。
一个JAR文件既可以包含类文件，也可以包含诸如图像和声音这些其他类型的文件。
JAR文件是压缩的，它使用了ZIP压缩格式。
pack200是一种较通常的ZIP压缩算法更加有效的压缩类文件的方式。


#### 创建JAR文件
可以使用jar工具制作JAR文件(在默认的JDK安装中，位于jdk/bin目录下)。
创建一个新的JAR文件应该使用的常见命令格式为:
```shell
jar cvf JARFileName file1 file2...
# or 
jar cvf CalculatorClasses.jar *.class icon.gif
```

可以将应用程序、程序组件(有时称为"beans")以及代码库打包在JAR文件中。例如，
JDK的运行时库包含一个非常庞大的文件rt.jar中。


#### 清单文件
除了类文件、图像和其他资源外，每个JAR文件还包含一个用于描述归档特性的`清单文件(manifest)`。
清单文件被命名为MANIFES.MF, 它位于JAR文件的一个特殊MEAT-INF子目录中。
最小的符合标准的清单文件：
```
Manifest-Version: 1.0
```

复杂的的清单文件:
```
Manifest-Version: 1.0
描述这个归档文件的行

Name: Woozle.class
描述这个文件的行
Name: com/mycompany/mypkg/
描述这个包的行
```
要想编辑清单文件，需要将希望添加到清单文件中的行放到文本文件中，然后运行：
```shell
jar cfm JARFileName ManifestFileName ...

# 例如， 要创建一个包含清单的JAR文件，应该运行:
jar cfm MyArchive.jar manifest.mf com/mycompany/mypkg/*.class

# 更新一个已有的JAR文件的清单:
jar ufm MyArchive.jar manifest-additions.mf
```


#### 可执行JAR文件
可以使用jar命令中的e选项指定程序的入口点，即通常需要在调用java程序加载器时指定的类:
```shell
jar cvfe MyProgram.jar com.mycompany.mypkg.MainAppClass files to add
```
或者，可以在清单中指定应用程序的主类，包括以下形式的语句:
```
Main-Class: com.mycompany.mypkg.MainAppClass
```
不要将扩展名.class添加到主类名中。

启动应用程序:
```shell
java -jar MyProgram.jar
```
根据操作系统的配置，用户甚至可以通过双击JAR文件图标来启动应用程序。


#### 资源
在applet和应用程序中使用的类通常需要使用一些相关的根据文件，例如:
* 图像和声音文件。
* 带有消息字符串和按钮标签的文本文件。
* 二进制数据文件，例如，描述地图布局的文件。
在java中，这些关联的文件被称为`资源(resource)`.

类加载器，操作非类文件:
1. 获得具有资源的Class对象，例如AboutPanel.class.
2. 如果资源是一个图像或声音文件，那么就需要调用getresource(filename)获得作为URL的资源位置，
然后利用getImage或getAudioClip方法进行读取。
3. 与图像或声音文件不同，其他资源可以使用getResourceAsStream方法读取文件中数据。

例如，要想利用about.gif图像文件制作图标:
```java
URL url = ResourceTest.class.getResource("about.gif");
Image img = new ImageIcon(url).getImage();
```
要想读取about.txt文件:
```java
InputStream stream = ResourceTest.class.getResourceAsStream("about.txt");
Scanner in = new Scanner(stream, "UTF-8");
```
还可以将它放在子目录中:
data/text/about.txt

