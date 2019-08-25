#### lambda构造器引用
构造器引用与方法引用很类似，只不过方法名为new。例如，Person::new是Person构造的一个引用。
```java
ArrayList<String> names = ...;
Stream<Person? stream = names.stream.map(Person::new);
List<Person> people = stream.collect(Collectors.toList());
```
可以用数组类型建立构造器引用。例如,int[]::new是一个构造器引用，它有一个参数:即数组的长度。
这等价于lambda表达式x-> new int[x]。
Java有一个限制，无法构造泛型类型T的数组。数组构造器引用对于克服这个限制很有用。表达式new
 T[n]会产生错误，因为这会改为new Object[n]。

 流库利用构造器引用解决了这个问题。可以把Person[]:new传入toArray方法:
 ```java
 Person[] people = stream.toArray(Person[]::new);
 ```
 toArray方法调用这个构造器来得到一个正确类型的数组。然后填充这个数组并返回。
 

#### 变量作用域
lambda表达式有3个部分:
1. 一个代码块；
2. 参数;
3. 自由变量的值，这是指非参数而且不在代码中定义的变量。

> lambda表达式中捕获的变量必须实际上是最终变量(effectively final)。
实际上的最终变量是指，这个变量初始化之后就不会再为它赋新值。
> lambda表达式的体与嵌套块有相同的作用域。
在lambda表达式中声明与一个局部变量同名的参数或局部变量是不合法的。
```java
Path first = Paths.get("/usr/bin");
Comparator<String> comp = 
	(first, second) -> first.length() - second.length();
```
在方法中，不能有两个同名的局部变量，因此，lambda表达式中同样也不能有同名的局部变量。
在一个lambda表达式中使用this关键字时，是指创建这个lambda表达式的方法的this参数:
```java
public class Application() {
	public void init() {
		ActionListener listener = event -> {
			System.out.println(this.toString());
		}
	}
}
```


#### 处理lambda表达式
使用lambda表达式的重点是`延迟执行(deferred execution)`。毕竟，如果想要立即执行代码，完全
可以直接执行，而无需把它包装在一个lambda表达式中。之所以希望以后再执行代码，这有很多原因:
1. 在一个单独的线程中运行代码;
2. 多次运行代码；
3. 在算法的适当位置运行代码(例如，排序中的比较操作);
4. 发生某种情况时执行代码(如，点击了一个按钮，数据到达，等等);
5. 只在必要时才运行代码。

假设要重复一个动作n次。将这个动作和重复次数传递到一个repeat方法:
```java
repeat(10, () -> System.out.println("Hello, world!"));
```
要接受这个lambda表达式，需要选择(偶尔可能需要提供)一个函数接口。
```java
public static void repeat(int n, Runnable action) {
	for (int i = 0; i < n; i++) action.run();
}
```
调用action.run() 时会执行这个 lambda表达式的主体。
