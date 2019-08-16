#### hashCode方法
散列码(hash code)是由对象导出的一整型值。散列码是没有规律的。
如果x和y是两个不同的对象，x.hashCode()与y.hashCode()基本上不会相同。

hashCode方法应该返回一个整型数值(也可以是负数),并合理地组合实例域的散
列码，以便能够让各个不同的对象产生的散列码更加均匀。

例如,下面是Employee类的hashCode方法
```java
public class Employee {
	public int hashCode() {
		return 7 * name.hashCode()
			+ 11 * new Double(Salary).hashCode()
			+ 13 * hireDay.hashCode();
	}
}
```

不过，还可以做得更好。首先，最好使用null安全的方法Objects.hashCode。如果
其参数为null, 这个方法会返回0，否则返回对参数调用hashCode的结果。
另外，使用静态方法Double.hashCode来避免创建Double对象:
```java
public int hashCode() {
	return 7 * Objects.hashCode(name)
		+ 11 * Double.hashCode(slary)
		+ 13 * Objects.hashCode(hireDay);
}
```

如果存在数组类型的域，那么可以使用静态的Arrays.hashCode方法计算一个散列码，
这个散列码由数组元素的散列码组成。


#### java.util.Object 
```java
int hashCode()
```
返回对象的散列码。散列码可以是任意的整数，包括正数或负数。两个相等的对象要求
返回相等的散列码。


#### java.util.Object
static int hash(Object... objects)
	返回一个散列码，由提供的所有对象的散列码组合而得到。

static int hashCode(Object a)
	如果a为null返回0，否则返回a.hashCode().


#### java.lang.(Integer|Long|Short|Byte|Double|Float|Character|Boolean)
static int hashCode((int|long|short|byte|double|float|char|boolean) value)
	返回给定值的散列码。


#### java.util.Arrays
static int hashCode(type[] a)
	计算数组a的散列码。组成这个数组的元素类型可以是object,int,long,short,char,byte,
	boolean,float或double。



