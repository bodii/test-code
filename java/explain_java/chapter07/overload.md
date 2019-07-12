Java允许一个类中存在多个相同名称的方法。同一个类中声明多个相同名称的方法称为`重载(overload)`方法。
不过，也存在“相同签名(signature)的方法不可以进行重载”的限制。
所谓方法的签名，就是方法名和形参的个数、类型的组合。
相同签名的方法不可以重载，换言之，就是为了让调用方明确区分出要调用哪个方法，如果形参的类型和个数
相同，就不可以重载。

重载有时还被称为`多重定义`,大家需要记住这两个术语。

```java
class Abc {
	static int ave(int x, int y) {
		return (x + y) / 2;
	}

	static int ave(int[] x, int[] y) {
		int v;
		for (int i = 0; i < x.length; i++)
			v = x[i] + y[i];
		return v / 2;
	}

	public static void main(String[] args) {
		System.out.println( ave(3, 5) );
	}
}
```
