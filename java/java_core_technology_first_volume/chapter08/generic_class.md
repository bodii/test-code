#### 定义简单泛型
一个泛型类(generic class)就是具有一个或多个类型变量的类。本章使用一个简单的Pair类作为
例子:
```java
public class Pair<T> {
	private T first;
	private T second;

	public Pair() {  first = null; second = null; }
	public Pair(T first, T second) { this.first = first; this.second = second; }

	public T getFirst() { return first; }
	public T getSecond() { return second; }

	public void setFirst(T newFirst) { first = newFirst; }
	public void setScond(T newSecond) { second = newSecond; }
}
```
Pair 类引入了一个类型变量T, 用尖括号(<>) 括起来，并放在类名的后面。泛型类可以有多个类型
变量。例如，可以定义Pair类，其中第一个域和第二个域使用不同的类型:
```java
public class Pair<T, U> { ... }
```

> 类型变量使用大写形式，且比较短，这是很常见的。在Java库中，使用变量E表示集合的元素类型，
> K和V分别表示表的关键字与值的类型。T（需要时还可以用临近的字母U和S)表示“任意类型”。

用具体的类型替换类型变量就可以实例化泛型类型，例如:
```java
Pair<String>
```
可以将结果想象成带有构造器的普通类:
```java
Pair<String>();
Pair<String>(String, String)
```
换句话说，泛型类可看作普通类的工厂。
