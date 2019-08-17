#### 泛型数组列表
ArrayList是一个采用`类型参数(type parameter)`的泛型类(generic class)。
为了指定数组列表保存的元素对象类型，需要用一对尖括号将类名括起来加在后面，
例如，
```java
ArrayList<Employee>
```

下面声明和构造一个保存Employee对象的数组列表:
```java
ArrayList<Employee> staff = new ArrayList<>();
```
这被称为"菱形"语法，因为空尖括号<>就像是一个菱形。可以结合new操作符使用菱形语法。
编译器会检查这个变量、参数或方法的泛型类型，然后将这个类型放在<>中。

使用add方法可以将元素添加到数组列表中。
```java
staff.add(new Employee(...);
```

数组列表管理着对象引用的一个内部数组。最终，数组的全部空间有可能被用尽。这就显现
出数组列表的操作魅力：`如果调用add且内部数组已经满了，数组列表就将自动地创建一个
大的的数组，并将所有的对象从较小的数组中拷贝到较大的数组中。`

如果已经清楚或能够估计出数组可能存储的元素数量，就可以在填充数组之前调用ensureCapacity
方法:
```java
staff.ensureCapacity(100);
```
这个方法调用将分配一个包含100个对象的内部数组。然后调用100次add,而不用重新分配空间。


另外，还可以把初始容量传递给ArrayList构造器：
```java
ArrayList<Employee> staff = new ArrayList<>(100);
```

分配数组列表，如下:
```java
new ArrayList<>(100); // capacity is 100
```
它与为新数组分配空间有所不同:
```java
new Employee[100]; // size is 100
```

> 数组列表的容量与数组的大小有一个非常重要的区别。如果为数组分配100个元素的存储空间，
数组就有100个空位置可以使用。而容量为100个元素的数组列表只是拥有保存100个元素的潜力
（实际上，重新分配空间的话，将会超过100），但是在最寝，甚至完成初始化构造之后，数组
列表根本就不含有任何元素。


* size 方法将返回数组列表中包含的实际元素数目。
```java
staff.size();
```
将返回staff数组列表的当前元素数量，它等价于数组a的a.length.
一旦能够确认数组列表的大小不再发生变化，就可以调用trimToSize方法。这个方法将存储区域
的大小调整为当前元素数量所需要的存储空间数目。垃圾回收器将回收多余的存储空间。
一旦整理了数组列表的大小，添加新元素就需要花时间再次移动存储块，所以应该在确认不会添
加任何元素时，再调用trimToSize。


#### java.util.ArrayList<E>
* ArrayList<E>()
	构造一个空数组列表

* ArrayList<E>(int initialCapacity)
	用指定容量构造一个空数组列表。
	参数:initalCapacity		数组列表的最初㝏容量

* boolean add(E obj)
	在数组列表的尾端添加一个元素。永远返回true.
	参数: obj		添加的元素

* int size()
	返回存储在数组列表中的当前元素数量。(这个值将小于或等于数组列表的容量。)

* void ensureCapacity(int capacity)
	确保数组列表在不重新分配存储空间的情况下就能够保存给定数量的元素。
	参数: capacity	需要的存储容量

* void trimtoSize()
	将数组列表的存储容量削减到当前尺寸。
