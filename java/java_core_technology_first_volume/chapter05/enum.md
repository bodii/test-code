#### 枚举类
一个典型的例子:
```java
public enum Size { SMALL, MEDIUM, LARGE, EXTRA_LARE };
```
实际上，这个声明定义的类型是一个类，它刚好有4个实例，在此尽量不要构造新对象。
因此，在比较两个枚举类型的值时，永远不需要调用equals, 而直接使用== 就可以了。

可以在枚举类型中添加一些构造器、方法和域。当然，构造器只是在构造枚举常量的时候
被调用:
```java
public enum Size {
	SMALL("S"), MEDIUM("M"), LARGE("L"), EXTRA_LARGE("XL");

	private String abbreviation;

	private Size(Sring abbreviation) { this.abbreviation = abbreviation; }
	public String getAbbreviation() { retrun abbreviation; }
}
```

> 所有的枚举类型都是Enum类的子类。它们继承了这个类的许多方法。其中最有用的一个
是toString,这个方法能够返回检举常量名。例如，Size.SMALL.toString()将返回字符串
“SMALL”。

toString的逆方法是静态方法valueOf:
```java
// 将s设置成sie.SMALL
Size s = Enum.valueOf(Size.class, "SMALL");
```

每个枚举类型都有一个静态的values方法，它将返回一个包含全部检举值的数组:
```java
Size[] values = Size.values();
// 返回包含元素Size.SMALL，Size.MEDIUM, Size.LARGE, Size.EXTRA_LARGE的数组。
```
ordinal方法返回enum声明中枚举常量的位置，位置从0开始计数。例如:Size.MEDIUM.ordinal()
返回1。


> 如同Class 类一样，鉴于简化的考虑，Enum类省略了一个类型参数。例如，实际上，应该将枚举
类型Size扩展为Enum<Size>。类型参数在compareTo方法中使。


#### java.lang.Enum
* static Enum valueOf(Class enumClass, String name)
	返回指定名字、给定类的枚举常量

* String toString()
	返回枚举常量名

* int ordinal()
	返回枚举常量在enum声明中的位置，位置从0开始计数

* int compareTo(E other)
	如果枚举量出现在other之前，则返回一个负值；
	如果this == other, 则返回0；
	否则，返回正值。
	枚举常量的出现次序在enum声明中给出。
