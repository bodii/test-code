#### 网格布局
网格布局像电子数据表一样，按行列排列所有的组件。不过，它的每个单元大小都是一样的。
在网格布局对象的构造器中，需要指定行数和列数(计算器):
```java
panel.setLayout(new GridLayout(4, 4));
```
添加组件，从第一行的每一列开始，然后是第一行的第二列，以此类推:
```java
panel.add(new JButton("1"));
panel.add(new JButton("2"));
```


#### java.awt.GridLayout
* GridLayout(int rows, int columns)
* GridLayout(int rows, int columns, int hgap, int wgap)
    构造一个新的GridLayout对象。rows或者columns可以为零，但不能同时为零，指定的每行或每列的组件数量可以任意的。
        参数:  rows             网格的行数
                    columns     风格的列数
                    hgap            以像素为单位的水平间距(如果为负值，则强行重叠)
                    vgap            以像素为单位的垂直间距(如果为负值，则强行重叠))