#### AWT事件处理机制的概要:
* 监听器对象是一个实现了特定监听器接口(listener interface)的类的实例。
* 事件源是一个能够注册监听器对象并发送事件对象的对象。
* 当事件发生时，事件源将事件对象传递给所有注册的监听器。
* 监听器对象将利用事件对象中的信息决定如何对事件做出响应。

下面是监听器的一个示例:
```java
ActionListener listener = ... ;
JButton button = new JButton("Ok");
button.addActionListener(listener);
```
只要按钮产生了一个“动作事件”，listener对象就会得到通知。对于按钮来说，正像我们所想到的，
动作事件就是点击按钮。
为了实现ActionListener接口，监听器类必须有一个被称为actionPerformed的方法，该方法接收一
个ActionEvent对象参数。
```java
class MyListener implements ActionListener {
	...
	public void actionPerformed(ActionEvent event) {
		// reaction to button click goes here
	}
}
```
只要用户点击按钮，JButton对象就会创建一个ActionEvent对象，然后调用listener.action Performed
(event)传递事件对象。可以将多个监听器对象添加到一个像按钮这样的事件源中。这样一来，只要用
户点击按钮，按钮就会调用所有监听器的actionPerformed方法。


#### 简洁地指定监听器
一个监听器类有多个实例的情况并不多见。更常见的情况是：每个监听器执行一个单独的动作。在这种
情况下，没有必要分别建立单独的类。只需要使用一个lambda表达式:
```java
exitButton.addActionListener(event -> System.exit(0));
```

现在考虑这样一种情况: 有多个相互关联的动作。在这种情况下，可以实现一个辅助方法:
```java
public void makeButton(String name, Color backgroundColor) {
	JButton button = new JButton(name);
	buttonPanel.add(button);
	button.addActionListener(
		event -> buttonPanel.setBackground(backgroundColor)
	);
}
```

需要说明，lambda表达式指示参数变量backgroundColor.
然后只需要调用:
```java
makeButton("yellow", Color.YELLOW);
makeButton("blue", Color.BLUE);
makeButton("red", Color.RED);
```
