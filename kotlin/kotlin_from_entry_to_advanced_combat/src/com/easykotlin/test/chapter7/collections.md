#### 3种集合类
集合类主要有3种: List(列表)、Set(集)和Map(映射)

#### 特征
List容器中的元素以线性方式存储，集合中可以存放重复对象。列表中的元素是有序地排列。
Set集容器的元素无序、不重复。
Map映射中持有的是"键值对"对象，每一个对象都包含一对键值K-V对象。Map映射容器中存储
的每个对象都有一个相关的关键字(Key)对象，关键字决定对象在映射中的存储位置。关键字
是唯一的。其实关键字本身并不能决定对象的存储位置，它通过散列(hashing)产生一个被称
做散列码(hash code)的整数值，这个散列码对应值（value)的存储位置。

List列表分为只读不可变的List和可变MutableList(可写入、删除数据)。

Set集也分为不可变Set和可变MutableSet(可写入、删除数据)。

Map与List、Set一样，Map也分为只读Map和可变MutableMap(可写入、删除数据)。Map没有继承
于Collection接口。