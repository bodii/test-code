#### 考虑实现Comparable接口
compareTo方法并没有在Object类中声明。相反，它是Comparable接口中唯一的方法。compareTo方法不但允许进行简单的等同性比较，而且允许执行顺序比较，除此之外，它与Object 的equals方法具有相似的特征，它还是个泛型(generic)。类实现了Comparable接口，就表明它的实例具有内在的排序关系(natural ordering)。

为实现Comparable接口的对象数组进行排序就这么简单:
```java
Arrays.sort(a);
```
对存储在集合中的Comparable对象进行搜索、计算极限值以及自动维护也同样简单。例如，下面的程序信赖于实现了Comparable接口的String类，它去掉了命令行参数列表中的重复参数，并按字母顺序打印出来:
```java
public class WordList {
    public static void main(String[] args) {
        Set<String> s = new TreeSet<>();
        Collections.addAll(s, args);
        System.out.println(s);
    }
}
```

#### compareTo方法的通用约定与equals方法的约定相似:
将这个对象与指定的对象进行比较。当该对象小于、等于或大于指定对象的时候，分别返回一个负整数、零或正整数。如果由于指定对象的类型而无法与该对象进行比较，则抛出ClassCastException异常。

在下面的说明中，符号sgn(expression)表示数学中的signum函数，它根据表达式(expression)的值为负值、零和正值，分别返回-1、0或1。
* 实现者必须确保所有的x和y都满足sgn(x.compareTo(y)) == -sgn(y.compareTo(x))。(这也暗示着，当且仅当y.compareTo(x)抛出异常时，x.compareTo(y)才必须抛出异常。)
* 实现者还必须确保这个比较关系是可传递的:(x.compareTo(y) > 0 && y.compareTo(z) > 0) 暗示着x.compareTo(z) > 0。
* 最后，实现者必须确保x.compareTo(y) == 0暗示着所有的z都满足sgn(x.compareTo(z)) == sgn(y.compareTo(z))。
* 强烈建议(x.compareTo(y) == 0) == x(x.equals(y)),但这并非绝对必要。一般说来，任何实现了Comparable接口的类，若违反了这个条件，都应该明确予以说明。推荐使用这样的说法: “注意：该类具有内在的排序功能，但是与equals不一致。”


#### compareTo约定中的条款
1. 如果颠倒了两个对象引用之间的比较方向，就会发生下面的情况:如果第一个对象小于第二个对象，则第二个对象一定大于第一个对象；如果第一个对象等于第二个对象，则第二个对象一定等于第一个对象；如果第一个对象大于第二个对象，则第二个对象一定小于第一个对象。
2. 如果一个对象大于第二个对象，并且第二个对象又大于第三个对象，那么第一个一定大于和第三个对象。
3. 在比较时被认为相等的所有对象，它们跟别的对象做比较时一定会产生同样的结果。