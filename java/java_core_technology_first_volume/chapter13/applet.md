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