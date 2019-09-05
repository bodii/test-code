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


#### 类型擦除
无论何时定义一个泛型类型，都自动提供一个相应的`原始类型(raw type)`。原始类型的名子就是
删去类型参数后的泛型名。`擦除(erased)`类型变量，并替换为限定类型(无限定的变量用Object)。
例如, Pair<T>的原始类型, 因为T是一个无限定的变量，所以直接用Object替换。
在程序中可以包含不同类型的Pair，例如， Pair<String> 或Pair<LocalDate>。而擦除类型后就变
成原始的Pair类型了。
```java
public  class Interval<T extends Comparable & Serializable> implements Serializable {
	private T lower;
	private T upper;

	...
	public Interval(T first, T second) {
		if (first.compareTo(second) <= 0 ) { lower = first; upper = second; }
		else { lower = second; upper = first; }
	}
}
原型类型Interval如下所示:
```java
public class Interval implements Serializable {
	private Comparable lower;
	private Comparable upper;

	...
	public Interval(Comparable first, Comparable second) { ... }
}
```


#### 泛型转换
需要记住有关Java泛型转换的事实:
* 虚拟机中没有泛型，只有普通的类和方法。
* 所有的类型参数都用它们的限定类型替换。
* 桥方法被合成来保持多态。
* 为保持类型安全性，必要时插入强制类型转换。


#### 泛型类的静态上下文中类型变量无效
不能在静态域或方法中引用类型变量。如:
```java
public class Singleton<T> {
	private static T singleInstance;   // Error
	public static T getSingleInstance() {  // Error
		if (singleInstance == null) construct new instance of T
			return singleInstance;
	}
}
```
禁止使用带有类型变量的静态域和方法。


#### 不能抛出或捕获泛型类的实例
既不能抛出也不能捕获泛型类对象。实际上，甚至泛型类扩展Throwable都是不合法的。
```java
public class Problem<T> extends Exception { /* ... */ } // Error--can't extends Throwable
```

catch 子句中不能使用类型变量。
```java
public static <T extends Throwable> void doWork(Class<T> t) {
	try {
		do work
	} catch(T e) { // Error--can't catch type variable
		Logger.global.info(...);
	}
}
```
不过，在异常规范中使用类型变量是允许的。以下方法是合法的:
```java
public static <T extends Throwable> void doWork(T t) throws T // ok {
	try { 
		do work
	} catch (Throwable realCause) {
		t.initCause(realCause);
		throw t;
	}
}
```


#### 通配符类型

##### 通配符概念 
通配符类型中，允许类型参数变化。例如，通配符类型:
```java
Pair<? extends Employee>
```

#### 通配符的超类型限定
通配符限定与类型变量限定十分类似，但是，还有一个附加的能力，即可以指定一个超`类型限定(
supertype bound)`:
```java
? super Manager
```

