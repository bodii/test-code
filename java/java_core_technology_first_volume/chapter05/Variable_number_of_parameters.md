#### 参数数量可变
```java
System.out.printf("%d %s", n , "widgets");

// printf方法是这样定义的：
public class PrintStream {
    public PrintStream printf(String fnt, Object... args) { 
		return format(fnt, args);
	}
}
```

自定义可变参数，并将参数指定为任意类型,甚至是基本类型:
```java
public static double max(double... values) {
	double largest = Double.NEGATIVE_INFINITY;
	for (double v: values)
		if (v > largest)
			largest = v;
	retrun largest; 
}

// 可以像下面这样调用:
double m = max(3.1, 4.04, -5);

// 编译器将new double[]{3.1, 4.04, -5} 传递给max方法。
```
> 允许将一个数组传递给可变参数方法的最后一个参数:
```java
System.out.printf("%d %s", new Object[] { new Integer(1), "widgets" });
```
