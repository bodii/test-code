#### 泛型方法
```java
class ArrayAlg {
	public static <T> getMiddle(T... a) {
		return  a[a.length / 2];
	}
}
```
> 类型变量放在修饰符(这里是public static)的后面，返回类型的前面。

泛型方法可以定义在普通类中，也可以定义在泛型类中。
当调用一个泛型方法时，在方法名前的尖括号中放入具体的类型:
```java
String middle = ArrayAlg.<String>getMiddle("John", "Q.", "PUblic");
```

在这种情况下，方法调用中可以省略<String>类型参数。编译器有足够的信息能够推断出所调用的方法。
它用names的类型(即String[])与泛型类型T[]进行匹配并推断出T一定是String。
```java
String middle = ArrayAlg.getMiddle("John", "Q.", "Public");
```

可以对类型变量设定限定(bound)实现:
```java
public static <T extends Comparable> T min(T[] a) ...
```
现在，泛型的min方法只能被实现Comparable接口的类(如String、LocalDate等)的数组调用。

> <T extends BoundingType>
表示T应该是绑定类型的`子类型(subtype)`。T和绑定类型可以是类，也可以是接口。选择关键字extends
的原因是更接近子类的概念，并且Java的设计者也不打算在语言中再添加一个新的关键字。

一个类型变量或通配符可以有多个限定，例如:
```java
T extends comparable & Serializable
```
限定类型用"&"分隔，而逗号用来分隔类型变量。
在java的继承中，可以根据需要拥有多个接口超类型，但限定中至多有一个类。如果用一个类作为限定，
它必须是限定列表中的第一个。

