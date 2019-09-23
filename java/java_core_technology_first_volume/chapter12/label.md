#### 标签和标签组件
标签是容纳文本的组件，它们没有任何的修饰(例如没有边缘)，也不能响应用户输入。可以利用标签标识组件。
1. 用相应的文本构造一个JLabel组件。
2. 将标签组件放置在距离需要标识的组件足够近的地方，以便用户可以知道标签所标识的组件。

可以指定右对齐标签:
```java
JLable label = new JLabel("User name:", SwingConstants.RIGHT)
// 或者
JLable label = new JLabel("User name:", JLabel.RIGHT)
```
利用setText和setIcon方法可以在运行期间设置标签的文本和图标。

可以在按钮、标签和菜单项上使用无格式文本或HTML文本。我们不推荐在按钮上使用HTML文本－－这样会影响观感。但是HTML文本在标签中是非常有效的。只要简单地将标签字符串放置在<html>...</html>中即可：
```java
label = new JLabel("<html><b>Required</b>entry:</html>");
```

#### javax.swing.JLabel
* JLabel(String text)
* JLabel(Icon icon)
* JLabel(String text, int align)
* JLabel(String text, Icon icon, int align)
    构造一个标签。
    参数： text             标签中的文本
                  icon              标签中的图标
                  align             一个SwingConstants的常量LEFT(默认)、CENTER或者RIGHT

* String getText()
* void setText(String text)
    获取或设置标签的文本。
* Icon getIcon()
* void setIcon(Icon icon)
    获取或设置标签的图标。