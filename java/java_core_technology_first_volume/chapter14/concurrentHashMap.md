#### 映射条目的原子更新
concurrentHashMap原来的版本只有为数不多的方法可以实现原子更新，这使得编程多少有些麻烦。假
设我们希望统计观察到的某些特性的频度。作为一个简单的例子，假设多个线程会遇到单词，我们想
统计它们的频率。


#### 对并发散列映射的批操作
Java SE 8为并发散列映射提供了批操作，即使有其他线程在处理映射，这些操作也能安全地执行。批
操作会遍历映射，处理遍历过程中找到的元素。无须冻结当前映射的快照。除非恰好知道批操作运行时
映射不会被修改，否则就要把结果看作是映射状态的一个近似。

有3种不同的操作:
* 搜索(search)为每个键或值提供一个函数，直到函数生成一个非null的结果。然后搜索终止，返回这
个函数的结果。
* 归约(reduce)组合所有键或值，这里要使用所提供的一个累加函数。
* forEach为所有键或值提供一个函数。

每个操作都有4个版本:
* operationKeys: 处理键。
* operationValues: 处理值。
* operation: 处理键和值。
* operationEntries: 处理Map.Entry对象。

对于上述各个操作，需要指定一个参数化阈值(parallelism threshold)。如果映射包含的元素多于这个
阈值，就会并行完成批操作。如果希望批操作在一个线程中运行，可以使用阈值Long.MAX_VALUE。
如果希望用尽可能多的线程运行操作，可以使用阈值1。

下面首先看search方法。有以下版本:
```java
U searchKeys(long threshold, BiFunction<? super K, ? extends U> f)
U searchValues(long threshold, BiFunction<? super V, ? extends U> f)
U search(long threshold, BiFunction<? super K, ? super V, ? extends U> f)
U searchEntries(long threshold, BiFunction<Map, Entry<K, V>, ? extends U> f)
```
例如，假设我们希望找出第一个出现次数超过1000次的单词，需要搜索键和值:
```java
String result = map.search(threshold, (k, v) -> v > 1000 ? k : null);
```

forEach方法有两种形式。第一个只为各个映射条目提供一个消费者函数，例如:
```java
map.forEach(threshold, (k, v) -> System.out.println(k + " -> " + v));
```

第二种形式还有一个转换器函数，这个函数要先提供，其结果会传递到消费者:
```java
map.forEach(threshold, 
		(k, v) -> k + " -> " + v, // Transformer
		System.out::println);  // Consumer
```

转换器可以用作为一个过滤器。只要转换器返回null,这个值就会被悄无声息地跳过。
例如，下面只打印有大值的条目:
```java
map.forEach(threshold,
		(k, v) -> v > 1000 ? k + " -> " + v : null,
		System.out::println);
```
reduce操作用一个累加函数组合其输入。例如，可以如下计算所有值的总和:
```java
Long sum = map.reduceValues(threshold, Long::sum);
```
与forEach类似，也可以提供一个转换器函数。可以如下计算最长的键的长度:
```java
Integer maxLength = map.reduceKeys(threshold,
		String::length,
		Integer::max);
```
转换器可以作为一个过滤器，通过返回null来排除不想要的输入。
在这里，我们要统计多少条目的值 > 1000：
```java
Long count = map.reduceValues(threshold,
		v -> v > 1000 ? 1L : null,
		Long::sum);
```
如果映射为空，或者所有条目都被过滤掉，reduce操作会返回null。如果只有一个元素，则返回其转换
结果，不会应用累加器。
