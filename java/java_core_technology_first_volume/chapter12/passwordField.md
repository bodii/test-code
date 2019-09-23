#### javax.swing.JPasswordField
* JPasswordField(String text, int columns)
    构造一个新的密码域对象。

* void setEchoChar(char echo)
    为密码域设置回显字符。注意:独特的观感可以选择自己的回显字符。0表示重新设置为默认的回显字符。

* char[] getPassword()
    返回密码域中的文本。为了安全起见，在使用之后应该覆写返回的数组内容(密码并不是以String的形式返回，这是因为字符串在被垃圾回收器回收之前会一直驻留在虚拟机中)。


#### 文本区
在JTextArea组件的构造器中，可以指定文本区的行数和列数。例如:
```java
textArea = new JTextArea(8, 40); // 8 lines of 40 columns each
```
可以通过开启换行特性来避免裁剪过长的行:
```java
textArea.setLineWarp(true); // long lines are wrapped
```


#### 滚动窗格
可以将文本区插入到滚动窗格(scroll pane)中。
```java
textArea = new JTextArea(8, 40);
JScrollPane scrollPane = new JScrollPane(textArea);
```