只要在字段声明时加上static,就可以实现该数据。由于是类中共享的变量，因此被称为`类变量(class variable`。
另外，由于在声明时加上了static，因此也被称为`静态字段(static field)`。

### 类变量的访问
正如其名，类变量不属于各个实例，而是`属于类`.
```
// 访问
类名 . 字段名
```

> 一个源程序中不可以声明两个以上的public类

### 一般的方针
除了应该在包中通过的内部类，原则上类都应该是public类，并被作为一个源程序实现。

对于规则较小的测试类，可以将多个类汇总到一个源程序中。此时，只将包含main方法的类
声明为public,其他的类都声明为非public类。


### Math类
Math类中定义了自然对数的底数E和圆周率PI的类变量。它们都是double型的常量，声明中都
加上了final.

```java
// java.lang.Math类的摘录
public final class Math {
	// 最接近自然对数的底数e的double值
	public static final double E = 2.7182818284590452354;

	// 最接近圆周率和其直径比PI的double值
	public static final double PI = 3.14159265358979323846;
}
```
