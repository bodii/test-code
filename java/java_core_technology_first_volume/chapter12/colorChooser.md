#### 颜色选择器
Swing颜色选择器`JColorChooser`，也是一种组件。

利用颜色选择器显示模式对话框:
```java
Color selectedColor = JColorChooser.showDialog(parent,title, initialColor);
```

另外，也可以显示模式颜色选择器对话框，需要提供:
* 一个父组件。
* 对话框的标题。
* 选择模式/无模式对话框的标志。
* 颜色选择器。
* OK和Cancel按钮的监听器(如果不需要监听器可以设置为null)。
