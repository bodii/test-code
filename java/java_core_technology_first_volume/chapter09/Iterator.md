```java
public interface Iterator<E> {
	E next();
	boolean hasNext();
	void remove();
	default void forEachRemaining(Consumer<? super E> action);
}

Collection<String> c =  ... ;
Iterator<String> iter = c.iterator();
while (iter.hasNext()) {
	String element = iter.next();
	do something with element
}

// 用for each 循环可以更加简练地表示同样的循环操作:
for (String element : c) {
	do something with element
}
```


#### java.util.Iterator
* boolean hasNext()
	如果存在可访问的元素，返回true.

* E next()
	返回将要访问的下一个对象。如果已经到达了集合的尾部，将抛出一个NoSuchElementException.

* void remove()
	删除上次访问的对象。这个方法必须紧跟在访问一个元素之后执行。如果上次访问之后，集合已
	经发生了变化，这个方法将抛出一个IllegalStateException.
