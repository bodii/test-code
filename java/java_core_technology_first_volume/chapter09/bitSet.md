#### 位集
Java平台的BitSet类用于存放一个位序列(它不是数学上的集，称为位`向量`或位数组更为合适)。
如果需要高效地存储序列(例如，标志)就可以使用位集。由于位集将位包装在字节里，所以，使用
位集比使用Boolean对象的ArrayList更加高效。
BitSet类提供了一个便于读取、设置或清除各个位的接口。使用这个接口可以避免屏蔽和其他麻烦
的位操作。如果将这些存储在int或long变量中就必须时行这些繁琐的操作。
例如，对于一个名为bucketOfBits的BitSet,
```java
bucketOfBits.get(i)
```
如果第i位处于“开“状态，就返回true;否则返回false。同样地,
```java
bucketOfBits.set(i)
```
将第i位置为"开"状态。最后，
```java
bucketOfBits.clear(i)
```
将第i位置为"关"状态。


#### java.util.BitSet
* BitSet(int initialCapacity)
	创建一个位集。

* int length()
	返回位集的“逻辑长度”，即1加上位集的最高设置的索引。

* boolean get(int bit)
	获得一个位。

* void set(int bit)
	设置一个位。

* void clear(int bit)
	清除一个位。

* void and(BitSet set)
	这个位集与另一个位集进行逻辑“AND”。

* void or(BitSet set)
	这个位集与另一个位集进行逻辑“OR”。

* void xor(BitSet set)
	这个位集与另一个位集进行逻辑"XOR".

* void andNot(BitSet set)
	清除这个位集中对应另一个位集中设置的所有位。
