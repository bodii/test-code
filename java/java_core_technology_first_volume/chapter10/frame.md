#### 创建框架
在Java中，顶层窗口(就是没有包含在其他窗口中的窗口)被称为框架(frame)。在AWT库中有一个称
为Frame的类，用于描述顶层窗口。这个类的Swing版本名为JFrame,它扩展于Frame类。
JFrame是极少数几个不绘制在画布上的Swing组件之一。因此，它的修饰部件(按钮、标题栏、
图标等)由用户的窗口系统绘制，而不是由Swing绘制。

> 绝大多数Swing组件类都以"J"开头，例如，JButton、JFrame等。在Java中有Button和Frame这样的
类，但它们属于AWT组件。如果偶然地忘记书写“J”，程序仍然可以进行编译和运行，但是将Swing和
AWT组件混合在一起使用将会导致视觉和行为的不一致。


#### javax.swing.JFrame
* Container getContentPane()
	返回这个JFrame的内容窗格对象。

* Component add(Component c)
	将一个给定的组件添加到该框架的内容窗格中


#### java.awt.Component
* void repaint()
	"尽量可能快地“重新绘制组件。

* Dimension getPreferredSize()
	要覆盖这个方法，返回这个组件的首选大小。


#### javax.swing.JComponent
* void paintComponent(Grphics g)
	覆盖这个方法来描述应该如何绘制自己的组件。


#### java.awt.Window
* void pack()
	调整窗口大小，要考虑到其组件的首选大小。


