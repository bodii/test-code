#### 反射
> 能够分析类能力的程序称为`反射(reflective)`。

反射机制的功能极其强大：
* 在运行时分析类的能力
* 在运行时查看对象，例如，编写一个toString方法供所有类使用。
* 实现通用的数组操作代码。
* 利用Method对象，这个对象很像C++中的函数指针。

#### Class类
在程序运行期间，Java运行时系统始终为所有的对象维护一个被称为运行时的类型标识。
这个信息跟踪着每个对象所属的类。虚拟机利用运行时类型信息选择相应的方法执行。
然而，可以通过专门的Java类访问这些信息。保存这些信息的类被称为Class,这个名字很
容易让人混淆。Object类中的getClass()方法将会返回一个Class类型的实例。
```java
Employee e;
...
Class c1 = e.getClass();
```
如同用一个Employee对象表示一个特定的雇员属性一样，一个Class对象将表示一个特定类的属性。
最常用的Class方法是getName.这个方法将返回类的名字:
```java
System.out.println(e.getClass().getName() + " " + e.getName());
```
如果类在一个包里，包的名字也作为类名的一部分:
```java
Random generator = new Random();
Class c1 = generator.getClass();
String name = c1.getName(); // name is set to "java.util.Random"
```

还可以调用静态方法forName获得类名对应的Class对象
```java
String className = "java.util.Random"
Class c1 = Class.forName(className);
```
如果类名保存在字符串中，并可在运行中改变，就可以使用这个方法。当然，这个方法只有在className
是类名或接口名时才能够执行。否则，forName方法将抛出一个checked exception（已检查异常)。
无论何时使用这个方法，都应该提供一个异常处理器(exception handler).

> 在启动时，包含main方法的类被加载。它会加载所有需要的类。这些被加载的类又要加载它们需要
的类，以此类推。
对于一个大型的应用程序来说，这将会消耗很多时间，用户会因此感到不耐烦。可以
使用下面这个技巧给用户一种启动速度比较快的幻觉。不过，要确保包含main方法的类没有显式地引用
其他的类。首先，显示一个启动画面；然后，通过调用Class.forName手动加载其他的类。

获得Class类对象的第三种方法, 如果T是任意的Java类型(或void关键字),T.class将代表匹配的类对象:
```java
Class c1 = Random.class;   // if you import java.util.*;
Class c2 = int.Class;
Class c3 = Double[].class;
```

> Class类实际上是一个泛型类。例如，Employee.class的类型是Class<Employee>。

> 鉴于历史原因，getName方法在应用于数组类型的时候返回一个很奇怪的名字:
```java
Double[].class.getName() // 返回 "[Ljava.lang.Double;"
int[].class.getName() // 返回"[I"
```

虚拟机为每个类型管理一个Class对象。因此，可以利用==运算符实现两个类对象比较的操作:
```java
if (e.getClass() == Employee.class)
```

还有一个很有用的方法newInstance(),可以用来动态创建一个类的实例:
```java
e.getClass().newInstance();
```
创建了一个与e具有相同类类型的实例。newInstance方法调用默认的构造器(没有参数的构造器)
初始化新创建的对象。如果这个类没有默认的构造器，就会抛出一个异常。

将forName与newInstance配合起来使用，可以根据存储在字符串中的类名创建一个对象:
```java
String s = "java.util.Random"
Object m = Class.forName(s).newInstance();
```


#### java.lang.reflect.Array
* static Object get(Object array, int index)
* static xxx getxxx(Object array, int index)
	(xxx是boolean、byte、char、double、float、int、long、short之中的一种基本类型.)
	这些方法将返回存储在给定位置上的给定数组的内容

* static void set(Object array, int index, Object newValue)
* static setxxx(Object array, int index, xxx newValue)
	这些方法将一个新值存储到给定位置上的给定数组中.

* static int getLength(Object array)
	返回数组的长度

* static Object newInstance(Class componentType, int length)
* static Object newInstance(Class componentType, int[] lengths)
	返回一个具有给定类型、给定维数的新数组。

