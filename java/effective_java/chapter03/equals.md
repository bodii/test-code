#### equals
在覆盖equals方法的时候，必须要遵守它的通用约定(来自Object的规范)。
equals方法实现了等价关系(equaivalence relation)：
* 自反性(reflexive): 对于任何非null的引用值x，x.equals(x)必须返回true。
* 对称性(symmetric): 对于任何非null的引用值x和y，当且仅当y.equals(x)返回
  true时，x.equals(y)必须返回true。
* 传递性(transitive): 对于任何非null的引用值x、y和z,如果x.equals(y)返回true,
  并且y.equals(z)也返回true，那么x.equals(z)也必须返回true。
* 一致性(consistent)：对于任何非null的引用值x和y，只要equals的比较操作在对象中
  所用的信息没有被修改，多次调用x.equals(y)就会一致地返回true,或者一致地返回false。
* 对于任何非null的引用值x, x.equals(null)必须返回false。

什么是等价关系?
不严格地说，它是一个操作符，将一组元素划分到其元素与另一个元素等价的分组中。
这些分组被称作`等价类(equivalence class)`。从用户的角度来看，对于有用的equals方法，
每个等价类中的所有无素都必须是可交换的。


> 在每个覆盖了equals方法的类中，都必须覆盖hashCode方法。
一个好的散列函数通常倾向于“为不相等的对象产生不相等的散列码”。