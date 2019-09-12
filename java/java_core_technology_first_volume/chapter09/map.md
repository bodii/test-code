#### 映射
集是一个集合，它可以快速地查找现有的元素。但是，要查看一个元素，需要有要查找元素的精确副本。
这不是一种非常通用的查找方式。
映射(map)数据结构就是为此设计的。映射用来存放键/值对。如果提供了键，就能够查找到值。

#### 基本映射操作
Java类库为映射提供了两个通用的实现: HashMap和TreeMap。这两个类都实现了Map接口。
散列映射对键进行散列，树映射用键的整体顺序对元素进行排序，并将其组织成搜索树。
散列或比较函数只能作用于键。与键关联的值不能进行散列或比较。

```java
Map<String, Employee> staff = new HashMap<>();
Employee harry = new Employee("Harry hacker");
staff.put("987-98-9996", harry);
```
每当往映射中添加对象时，必须同时提供一个键。在这里，键是一个字符串，对应的值是Employee对象。
要想检索一个对象，必须使用一个键。
```java
String id = "987-98-9996";
e = staff.get(id);
```
如果在映射中没有与给定键对应的信息，get将返回null。
null返回值可能并不方便。有时可以有一个好的默认值，用作为映射中不存在的键。然后使用getOrDefault
方法。
```java
Map<String, Integer> scores = ... ;
int score = scores.get(id, 0);
```
键必须是唯一的。不能对同个键存放两个值。如果对同个键两次调用put方法，第二个值就会取代第一
个值。实际上，put将返回用这个键参数存储上一个值。
remove 方法用于从映射中删除给定键对应的元素。
size方法用于返回映射中的元素数。
forEach方法用于迭代处理映射的键值。可以提供一个接收键和值的lambda表达式
```java
scores.forEach(k, v) ->
	System.out.println("key=" + k + ", value=" + v));
```
