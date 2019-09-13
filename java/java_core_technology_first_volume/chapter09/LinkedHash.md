#### 链接散列集与映射
LinkedHashSet和LinkedHashMap类用来记住插入元素项的顺序。这样就可以避免在散列表中的
项从表面上看是随机排列的。
```java
Map<String, Employee> staff = new LinkedHashMap<>();
staff.put("144-25-5464", new Employee("Amy Lee"));
staff.put("567-24-2546", new Employee("Harry Hacker"));
staff.put("157-62-7935", new Employee("Gary Cooper"));
staff.put("456-62-5527", new Employee("Francesca Cruz"));

// 然后， staff.keySet().iterator()以下面的次序枚举键:
/*
144-25-546
567-24-2546
157-62-7935
456-62-5527
*/

// 并且staff.values().iterator()以下列顺序枚举这些值:
/*
Amy Lee
Harry Hacker
Gary Cooper
Francesca Cruz
*/
```

甚至可以让这一过程自动化。即构造一个LinkedHashMap的子类，然后覆盖下面这个方法:
```java
protected boolean removeEldestEntry(Map.Entry<K, V> eldest)
```

每当方法返回true时，就添加一个新条目，从而导致删除eldest条目。例如，下面的高速缓存可以存
放100个元素:
```java
Map<K, V> cache = new 
	LinkedHashMap<>(128, 0.75F, true) {
		protected boolean remvoeEldestEntry(Map.Entry<K, V> eldest) {
			return size() > 100;
		}
	}
```
