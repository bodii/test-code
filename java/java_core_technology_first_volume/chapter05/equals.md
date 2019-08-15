#### java语言规范要求equals方法具有下面的特性:
> 1. 自反性: 对于任何非空引用x, x.equals(x)应该返回true;
> 2. 对称性: 对于任何引用x和y,当且仅当y.equals(x)返回true, x.equals(y)也应该返回true。
> 3. 传递性: 对于任何引用x、y和z,如果x.equals(y)返回true,y.equals(z)返回true,
	x.equals(z)也应该返回true.
> 4. 一致性: 如果x和y引用的对象没有发生变化，反复调用x.equals(y)应该返回同样的结果。
> 5. 对于任意非空引用x, x.equals(null)应该返回false。


#### 编写一个完美的equals方法的建议:
1. 显式参数命名为otherObject, 稍后需要将它转换成另一个叫做other的变量

2. 检测this与otherObject是否引用同一个对象:
```java
if (this == otherObject) return true;
```
这条语句只是一个优化。实际上，这是一种经常采用的形式。因为计算这个等式要比一个个地比较
类中的域所付出的代价小得多。

3. 检测otherObject是否为null, 如果为null, 返回false.这项检测是很必要的。
```java
if (otherObject == null) return false;
```

4. 比较this与otherObject是否属于同一个类。如果equals的语义在每个子类中有所改变，就使用
getClass检测:
```java
if (getClass() != otherObject.getClass()) return false;
```

5. 将otherObject转换为相应的类类型变量:
```java
ClassName other = (ClassName) otherObject
```

6. 现在开始对所有需要比较的域进行比较了。使用==比较基本类型域，使用equals比较对象域。如
果所有的域都匹配，就返回true; 否则返回false;
```java
return field1 == other.field1
		&& Objects.equals(field2, other.field2)
		&& ...;
```


如果在子类中重新定义equals,就要在其中包含调用super.equals(other)。



#### java.util.Arrays
> static Boolean equals(type[] a, type[] b)
	如果丙个数组长度相同，并且在对应的位置上数据元素也均相同，将返回true,数组的元素类型可
	以是Object、int、long、short、char、byte、boolean、float或double.


#### java.util.Objects
> static boolean equals(Ojbect a, Object b)
	如果a和b都为null, 返回true; 如果只有其中之一为null，则返回false;否则返回a.equeals(b)
