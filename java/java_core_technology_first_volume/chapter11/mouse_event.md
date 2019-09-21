#### java.awt.event.MouseEvent
* int getX()
* int getY()
* Point getPoint()      
    返回事件发生时，事件源组件左上角的坐标X(水平)和Y(竖直)，或点信息。

* int getClickCount()
    返回与事件关联的鼠标连击次数("连击"所指定的时间间隔与具体系统有关)。


#### java.awt.event.InputEvetn
* int getModifiersEx()
    返回事件扩展的或"按下"(down)的修饰符。使用下面的掩码值检测返回值:
    BUTTON1_DOWN_MASK
    BUTTON2_DOWN_MASK
    BUTTON3_DOWN_MASK
    SHIFT_DOWN_MASK
    CTRL_DOWN_MASK
    ALT_DOWN_MASK
    ALT_GRAPH_DOWN_MASK
    META_DOWN_MASK

* static String getModifiersExText(int modifiers)
返回用给定标志集描述的扩展或“按下"(down)的修饰符字符串，例如"Shift+Button1"。


#### java.awt.ToolKit
* public Cursor createCustomCursor(Image image, Point hotSpot, String name)
	创建一个新的定制光标对象。
	参数: image			光标活动时显示的图像
		  hotSpot		光标热点(箭头的顶点或十字中心)
		  name			光标的描述，用来支持特殊的访问环境


#### java.awt.Component
* public void setCursor(Cursor cursor)
	用光标图像设置给定光标
