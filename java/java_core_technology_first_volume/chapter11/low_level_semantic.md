#### 语义事件和底层事件
AWT将事件分为`底层(low-level)`和`语义(semantic)`事件。语义事件是表示用户动作的事件，例如，
点击按钮；因此,ActionEvent是一种语义事件。底层事件是形成那些事件的事件。在点击按扭时，包含
了按下鼠标、连续移动鼠标、抬起鼠标(只有鼠标在按钮中抬起才引发)事件。或者在用户利用TAB键选择
按钮，并利用空格键激活它时，发生的敲击键盘事件。同样，调节滚动条是一种语义事件，但拖动鼠标
是底层事件。

java.awt.event包中最常用的语义事件类:
* ActionEvent(对应按钮点击、菜单选择、选择列表或在文本框中enter);
* AjustmentEvent(用户调节滚动条);
* ItemEvent(用户从复选框或列表中选择一项).

常用的5个底层事件类:
* KeyEvent(一个键被按下或释放);
* MouseEvent(鼠标键被按下、释放、移动或拖动);
* MouseWheelEvent（鼠标滚动被转动);
* FouseEvent(某个组件获得焦点或失去焦点);
* WindowEvent(窗口组件被改变).

下列接口将监听这些事件:
ActionListener			MouseMorionListener
AdjustmentListener		MouseWhellListener
FocusListener			WindownListener
ItemListener			WindownFocusListener
KeyListener				WindowStateListener
MouseListener

常用的适配器类:
FocusAdapter			MouseMotionAdapter
MouseAdapter			WindownAdapter	
MouseAdapter
