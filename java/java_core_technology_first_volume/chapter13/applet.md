#### 一个简单的applet
applet就是一个扩展了java.applet.Applet类的Java类。
使用Swing来实现applet。所有applet都将扩展JApplet类，它是Swing applet的超类。

如果你的applet包含Swing组件，必须扩展JApplet类。普通Applet中的Swing组件不能正确绘制。


要执行applet,需要完成以下3个步骤:
1. 将Java源文件编译为类文件。
2. 将类打包到一个JAR文件中。
3. 创建一个HTML文件，告诉浏览器首先加载哪个类文件，以及如何设定applet的大小。

下面给出这个文件的内容:
```xml
<applet class="applet/NotHelloWorld.class" archive="NotHelloWorld.jar", width="300" height="300"> </applet>
```

在浏览器中查看applet之前，最好先在JDK自带的applet viewer(applet查看器)程序中进行测试。
要使用applet查看测试applet，可以在命令行输入:
```shell
appletviewer NotHelloWolrdApplet.html
```
applet查看器程序的命令行参数是HTML文件名，而不是类文件。


很容易把一个图形化Java应用转换为可以嵌入在Web页面中的applet.
具体步骤:
1. 建立一个HTML页面，其中包含加载applet代码的适当标记。
2. 提供JApplet类的一个子类。将这个类标记为public。否则applet将无法加载。
3. 删去应用中的main方法。不要为应用构造构架窗口。你的应用将在浏览器中显示。
4. 把所有初始化代码从框架窗口移至applet的init方法。不需要明确构造applet对象，浏览器会实例化applet对象并调用init方法。
5. 删setSize调用；对applet来说，用HTML文件中的width和height参数就可以指定大小。
6. 删除SetDefaultCloseOperation调用。applet不能关闭；浏览器退出时applet就会终止运行。
7. 如果应用调用setTitle，则删除这个方法调用。applet没有标题栏。
8. 不要调用setVisible(true)。applet会自动显示。