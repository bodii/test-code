#### 基本类型和包装类

基本类型			包装类
byte				Byte
short				Short
int					Integer
long				Long
float				Float
double				Double
char				Character
boolean				Boolean

#### 通过类变量提供基本类型的特性
包装类通过类变量为对应的基本类型提供其可以表示的最小值MIN_VALUE 和最大值MAX_VALUE
(Boolean型除外)

包装类中还定义了类变量，以表示基本类型所占用的位数。

#### 可以创建持有对应的基本类型的值的类类型实例
各个包装类通过字段来持有对应的基本类型的值。例如，Integer类持有int型的字段，Double类持有
double型的字段。

由于提供了用于接收对应的基本型参数的构造函数，因此，我们可以像下面这样来创建包装类类型的
实例。
```java
Integer i = new Integer(5);
Double d = new Double(3.14);
```

包括包装类在内，Java中的类都是Object类的子孙。为此，只能应用于引用类型的操作也可以应用于
整数值或实数值(包装它们的包装类类型的实例)。

自动装箱(auto boxing):
```java
Integer i = 5; // int 到Integer的自动装箱
Double d = 3.14; // double到Double的自动装箱
```

#### 通过方法提供各种操作
```java
int n = 5;
System.out.println("n = " + n);
// 相当于
System.out.println("n = " + Integer(n).toString());
```
也就是说，当执行"字符串+数值"的运算时，处理顺序如下：
1. 创建持有该数值的包装类的实例，对该实例应用toString方法，创建字符串。
2. 执行字符串的拼接(字符串和转换后的字符串的拼接)

对于所有的包装类，都提供了将数值转换为字符串的toString方法。

toString方法的逆向转换，即将字符串转换为数值的就是parse...方法。这是类方法，其中的...部分是
首字母在大写的基本类型的类型名。Integer类中提供了parseInt方法,Float类中则提供parseFloat方法

例如,Integer.parseInt("3145") 会返回整数值3145,Long.parseLong("12345"),则返回long型的整数值
12345L。

