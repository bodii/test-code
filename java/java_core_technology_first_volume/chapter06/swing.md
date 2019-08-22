#### javax.swing.JOptionPane
* static void showMessageDialog(Component parent, Object message)
	显示一个包含一条消息和OK按钮的对话框。这个对话框将位于其parent组件的中央。
	如果parent为null,对话框将显示在屏幕的中央。


#### java.swing.Timer
* Timer(int interval, ActionListener listener)
	构造一个定时器，每隔interval毫秒通告listener一次。

* void start()
	启动定时器，一旦启动成功，定时器将调用监听器的actionPerformed.

* void stop()
	停止定时器。一旦停止成功，定时器将不再调用监听器的actionPerformed.

#### java.awt.Toolkit
* static Toolkit getDefaultToolkit()
	获得默认的工具箱。工具箱包含有关GUI环境的信息

* void beep()
	发出一声铃响。


