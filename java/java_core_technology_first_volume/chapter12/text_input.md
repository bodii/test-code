#### 文本输入
Swing用户界面组件，具有用户输入和编辑文本功能的组件。
｀文本域(JTextFiled)｀ 和｀文本区(JTextArea)｀组件用于获取文本输入。文本域只能接收单行文本的输入，而文本区能够接收多行文本的输入。JPassword也只能接收单行文本的输入，但不会将输入的内容显示出来。
这三个类都继承于JTextComponent类。由于JTextComponent是一个抽象类，所以不能够构造这个类的对象。另外，在Java中常会看到这种情况。
在一个文本域和文本区内获取(get)、设置(set)文本的方法实际上都是JTextComponent类中的方法。


#### javax.swing.text.JTextComponent
* String getText()
* void setText(String text)
    获取或设置文本组件中的文本。

* boolean isEditable()
* void setEdiTable(boolean b)
    获取或设置editable特性，这个特性决定了用户是否可以编辑文本组件中的内容。


#### 文本域
把文本域添加到窗口的常用办法是将它添加到面板或者其他容器中，这与添加按钮完全一样:
```java
JPanel panel = new JPanel();
JTextField textField = new JTextField("Default input", 20);
panel.add(textField);
```
这段代码将添加一个文本域，同时通过传递字符串"Default input"时行初始化。
使用setColumns方法改变了一个文本域的大小之后，需要调用包含这个文本框的容器的revalidate方法。
```java
textField.setColumns(10);
panel.revalidate();
```
realidate方法会重新计算容器内所有组件的大小，并且对它们重新进行布局。调用revalidate方法以后，布局管理器会重新设置容器的大小，然后就可以看到改变尺寸后的文本域了。
revalidate方法是JComponent类中的方法。它并不是马上就改变组件大小，而是给这个组件加一个需要改变大小的标记。这样就避免了多个组件改变大小时带来的重复计算。但是，如果想重新计算一个JFrame中的所有组件，就必须调用validate方法－－JFrame没有扩展JComponent。

构造一个空白文本域：
```java
JTextField textField = new JTextFiled(20);
```


#### javax.swing.JTextField
* JTextField(int cols)
    构造一个给定列表的空JTextField对象。

* JTextField(String text, int cols)
    构造一个给定列表，给定初始字符串的JTextField对象。

* int getColumns()
* void setColumns(int cols)
    获取或设置文本域使用的列数。


#### javax.swing.JComponent
* void revalidate()
    重新计算组件的位置和大小。

* void setFont(Font f)
    设置组件的字体


#### java.awt.Component
* void validate()
    重新计算组件的位置和大小。如果组件是容器，容器中包含的所有组件的位置和大小也被重新计算。

* Font getFont()
    获取组件的字体。
