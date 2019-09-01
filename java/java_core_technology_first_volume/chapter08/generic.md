#### 泛型程序设计
> `泛型程序设计(Generic programming)`意味着编写的代码可以被很多不同的类型的对象所重用
例如，我们并不希望为聚焦String和File对象分别设计不同的类。实际上，也不需要这样做，因为
一个ArrayList类可以聚集任何类型的对象。这是一个泛型程序设计的实例。

在Java中增加范型类之前，泛型程序设计是用继承实现的。ArrayList类只推护一个Object引用的
数组:
```java
public class ArrayList {
	private Object[] elementData;

	...
	public Object get(int i) { ... }
	public void add(Object o) { ... }
}
```
这种方法两具问题。当获取一个值时必须时行强制类型转换:
```java
ArrayList files = new ArrayList();
String filename = (String) files.get(0);
```
此外，这里没有错误检查。可以向数组列表中添加任何类的对象。
```java
files.add(new File("..."));
```

> 泛型提供了一个更好的解决方案: 类型参数(type parameters)。
ArrayList类有一个类型参数用来指示元素的类型:
```java
ArrayList<String> files = new ArrayList<String>();
```

> 构造函数中可以省略泛型类型:
```java
ArrayList<String> files = new ArrayList<>();
```
省略的类型可以从变量的类型推断得出。
