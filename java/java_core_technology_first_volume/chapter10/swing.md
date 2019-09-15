#### Swing
在每个Swing程序中，有两个技术问题需要强调:
1. 首先，所有的Swing组件必须由`事件分派线程(event dispatch thread)` 进行配置，线程将
	鼠标点击和按键控制转移到用户接口组件。
```java
EventQueue.invokeLater(
		() -> {
			statements
		}
);
```

> 许多Swing程序并没有在事件分派线程中初始化用户界面。在主线程中完成初始化是通常采用的方式。
遗憾的是，由于Swing组件十分复杂，JDK和程序员无法保证这种方式的安全性。虽然发生错误的概率
非常小，但任何人不愿意成为遭遇这个问题的不幸者之一。即使代码看起来有些神秘，也最好能够保
证其正确性。

接下来，定义一个用户关闭这个框架时的响应动作。对于这个程序而言，只让程序简单地退出即可。
选择这个响应动作的语句是:
```java
frame.setDefaultColseOperation(JFrame.EXIT_ON_CLOSE);
```
在默认情况下，用户关闭窗口时只是将框架隐藏起来，而程序并没有终止(在最后一个构架不可见之后，
程序再终止，这样处理比较合适，而Swing却不是这样工作的)。
简单地构造框架是不会自动地显示出来的，框架起初是不可见的。这就给程序员了一个机会，可以在框
架第一次显示之前往其中添加组件。为了显示框架，main方法需要调用框架的setVisible方法。


#### 框架定位
JFrame类本身只包含若干个改变框架外观的方法。当然，通过继承，从JFrame的各个超类中继承了许多
用于处理框架大小和位置的方法。其中最重要的有下面几个:
* setLocation 和 setBounds方法用于设置框架的位置。
* setIconImage用于告诉窗口系统在标题栏、任务切换窗口等位置显示哪个图标。
* setTitle用于改变标题栏的文字。
* setResizable利用一个boolean值确定框架的大小是否允许用户改变。

#### 确定合适的框架大小
为了得到屏幕的大小，需要按照下列步骤操作。调用Toolkit类的静态方法getDefaultToolkit得到一个
Toolkit对象(Toolkit类包含很多与本地窗口系统打交道的方法)。然后，调用getScreenSize方法，这
个方法以Dimension对象的形式返回屏幕的大小。Dimension对象同时用公有实例变量width和height保
存着屏幕的宽度和高度。下面是相关的代码:
```java
Toolkit kit = Toolkit.getDefaultToolkit();
Dimension screenSize = kit.getScreenSize();
int screenWith = screenSize.widht;
int screenHeight = screenSize.height;
```
下面，将框架大小设定为上面取值的50%,然后，告知窗口系统定位框架:
```java
setSize(screenWidth / 2, screenHeight / 2);
setLocationByPlatform(true);
```
另外，还提供一个图标。
```java
Image img = new ImageIcon("icon.gif').getImage();
setIconImage(img);
```
在通常情况下，将程序的主框架尺寸设置为最大。可以通过调用下列方法将框架设置为最大:
```java
frame.setExtendedState(Frame.MAXIMIZED_BOTH);
```

GraphicsDevice类还允许以全屏模式执行应用。
