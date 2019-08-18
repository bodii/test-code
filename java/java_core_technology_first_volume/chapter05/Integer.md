#### java.lang.Integer
* int intValue()
	以int的形式返回Integer对象的值(在Number类中覆盖了intValue方法)。

* static String toString(int i)
	以一个新String对象的形式返回给定数值i的十进制表示。
	
* static String toString(int i , int radix)
	返回数值i的基于给定radix参数进制的表示

* static int parseInt(String s)
* static int parseInt(String s, int radix)
	返回字符串s表示的整型数值，给定字符串表示的是十进制的整数(第一种方法),
	或者是radix参数进制的整数(第二种方法)。

* static Integer valueOf(String s)
* static Integer valueOf(String s, int radix)
	返回用s表示的整型数值进行初始化后的一个新Integer对象，给定字符串表示的是
	＋进制的整数(第一种方法), 或者是radix参数进制的整数(第二种方法)。


#### java.text.NumberFormat
* Number parse(String s)
	返回数字值，假设给定的String表示了一个数值。