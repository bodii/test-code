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


#### java.util.Map<K, V>
* V get(Object key)
	获取与键对应的值；返回与键对应的对象，如果在映射中没有这个对象则返回null。键可以为null。

* default V getOrDefault(Object key, V defaultValue)
	获得与键关联的值；返回与键关联的对象，或者如果未在映射中找到这个键，则返回defaultValue

* V put(K key, V value)
	将键与对应的值关系插入到映射中。如果这个键已经存在，新对象将取代与这个键对应的旧对象。
	这个方法将返回键对应的旧值。如果这个键以前没有出现过则返回null.键可以为null，但值不能
	为null。

* void putAll(Map<? extends K, ? extends V> entries)
	将给定映射中的所有条目添加到这个映射中。

* boolean containsKey(Object key)
	如果在映射中已经有这个键，返回true。

* boolean containsValue(Object value)
	如果映射中已经有这个值，返回true.

* default void forEach(BiConsumer<? super K, ? super V> action)
	对这个映射中的所有键／值对应用这个动作。


#### java.uti.HashMap<K, V>
* HashMap()
* HashMap(int initialCapacity)
* HashMap(int initialCapacity, float loadFactor)
	用给定的容量和装填因子构造一个空散列映射（装填因子是一个0.0 ~ 1.0之间的数值。这个数值
	决定散列表填充的百分比。一旦到了这个比例，就要将其再散列到更大的表中）。默认的装填因
	子是0.75

#### java.util.TreeMap<K, V>
* TreeMap()
	为实现Comparable接口的键构造一个空的树映射。

* TreeMap(Comparator<? super K> c)
	构造一个树映射，并使用一个指定的比较器对键进行排序。

* TreeMap(Map<? extends k, ? extends V> entries)
	构造一个树映射，并将某个映射中的所有条目添加到树映射中。

* TreeMap(SortedMap<? extends K, ? extends V> enteries)
	构造一个树映射，将某个有序映射中的所有条目添加到树映射中，并使用与给定的有序映射相同的
	比较器。

#### java.util.SortedMap<K, V>
* Comparator<? super K> comparator()
	返回对键进行排序的比较器。如果键是用Comparable接口的compareTo方法进行比较的，返回null。

* K firstKey()
* K lastKey()
	返回映射中最小元素和最大元素



#### java.util.Map<K, V>
* defailt V merge(K key, V value, BiFunction<? super V, ? super V, ? extends V> 
remappingFunction)
	如果key与一个非null值v关联，将函数应用到v和value,将key与结果关联，或者如果结果为null,
	则删除这个键。否则，将key与value关联，返回get(key)。

* default V compute(K key, BiFunction<? super K, ? super V, ? extends V> reappingFunction)
	将函数应用到key和get(key)。将key与结果关联，或者如果结果为null,则删除这个键。返回get(key).

* default V computeIfPresent(K key, BiFunction<? super K, ? super V, ? extends V>
remappingFunction)
	如果key与一个非null值v关联，将函数应用到key和v，将key与结果关联，或者如果结果为null,则删除
	这个键，返回get(key)。

* default V computeIfAbsent(K key, Function<? super k, ? extends V> mapppingFunction)
	将函数应用到key，除非key与一个非null值关联。将key与结果关联，或者如果结果为null，则
	删除这个键。返回get(key).

* default void replaceAll(BiFunction<? super K, ? super V, ? extends V> function)
	在所有映射项上应用函数。将键与非null结果关联，对于null结果，则将相应的键删除。
