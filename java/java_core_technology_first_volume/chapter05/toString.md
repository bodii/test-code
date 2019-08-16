#### toString方法
在Object中还有一个重要的方法，就是toString方法，它用于返回表示对象值的字符串。

绝大多数(但不是全部)的toString方法都遵循这样的格式: 类的名字，随后是一对方括号
括起来的域值。

下面是Employee类中的toString方法的实现:
```java
public String toString() {
	return "Employee[name=" + name
		+ ",salary=" + salary
		+ ",hireDay=" + hireDay
		+ "]";
}
```

实际上，还可以设计得更好一些，最好通过调用getClass().getName()获得类名的字符串，
而不要将类名硬加到toString方法中。
```java
public String toString() {
	return getClass().getName()
		+ "[name=" + name
		+ ",salary=" + salary
		+ ",hireDay=" + hireDay
		+ "]";
}
```

> toString方法可以供子类调用。

只要对象与一个字符串通过操作符 "+" 连接起来，Java编译就会自动地调用toString方法，
以便获得这个对象的字符串描述。

> 在调用x.toString()的地方可以用`"" + x` 替代。这条语句将一个空串与x的字符串表示相
连接。这里的x就是x.toString()。与toString不同的是，如果x是其本类型，这条语句照样能
够执行。


数组继承了object类的toString方法，数组类型将按照旧的格式打印：
```java
int[] luckyNumbers = { 2, 3, 5, 7, 11, 13 };
String s = "" + luckyNumbers;
```
生成字符串"[I@1a46e30"(前缀[I表明是一个整型数组)。修正的方式是调用静态方法Arrays.toString
```java
String s = Arrays.toString(luckyNumbers);
```
将生成字符串"[2,3,5,7,11,13]"。
