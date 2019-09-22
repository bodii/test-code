#### 布局管理
`流布局管理器(flow layout manager)`管理，这是面板的默认布局管理器。

通常，组件放置在容器中，布局管理器决定容器中的组件具体的位置和大小。
按钮、文本域和其他的用户界面元素都继承于Component类，组件可以放置在
面板这样的容器中。由于Container类继承于Component类，所以容器也可以放
置在另一个容器中。

每个容器都有一个默认的布局管理器，但可以重新时行设置。例如，使用下列语句:
```java
panel.setLayout(new GirdLayout(4, 4));
```
这个面板将用GirdLayout类布局组件。可以往容器中添加组件。容器的add方法将把组件
和放置的方位传递给布局管理器。


#### java.awt.Container
* void SetLayout (layoutManager m)
	为容器设置布局管理器

* Component add(Component c)
* Component add(Component c, Object constraints)
	将组件添加到容器中，并返回组件的引用。
	参数: c				要添加的组件
		constraints		布局管理器理解的标识符


#### java.awt.FlowLayout
* FlowLayout()
* FlowLayout(int align)
* FlowLayout(int align, int hgap, int vgap)
	构造一个新的FlowLayout对象。
	参数: align			LEFT、 CENTER 或者RIGHT
		  hgap			以像素为单位的水平间距(如果为负值，则强行重叠)
		  vgap          以像素为单位的垂直间距(如果为负值，则强行重叠)

