#### 并行数组算法
在Java SE 8中， Arrays类提供了大量并行化操作。静态Arrays.parallelSort方法可以对一个基本类
型值或对象的数组排序。例如:
```java
String contents = new String(Files.readAllBytes(
			Paths.get("alice.txt")), StandardCharsets.UTF_8);
String[] words = contents.split("[\\P{L}]+");
Arrays.parallelSort(words);
```
对对象排序时，可以提供一个Comparator。
```java
Arrays.parallelSort(words, Comparator.comparing(String::length));
```
对于所有方法都可以提供一个范围的边界，如:
```java
values.parallelSort(values.length / 2, values.length);
```

parallelSetAll方法会用由一个函数计算得到值填充一个数组。这个函数接收元素索引，然后计算机应
位置上的值。
```java
Arrays.parallelSetAll(values, i -> i % 10);
// Files values with 0 1 2 3 4 5 6 7 8 9 0 1 2 ...
```
显然，并行化对这个操作很有好处。这个操作对于所有基本类型数组和对象数组都有相应的版本。

parallelPerfix方法，它会用对应一个给定结合操作的前缀的累加结果替换各个数组元素。
```java
Arrays.parallelPrefix(values, (x, y) -> x * y) 
// [1, 1 * 2, 1 * 2 * 3, 1 * 2 * 3 * 4 ....]
```
