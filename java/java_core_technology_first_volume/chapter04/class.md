### 类之间的关系

在类之间，最常见的关系有:
* 依赖("uses-a")
* 聚合("has-a")
* 继承("is-a")

依赖(dependence)，即"uses-a"关系，是一种最明显的、最常见的关系。例如，Order类使用
Accout类是因为Order对象需要访问Account对象查看信用状态。但是Item类不依赖于Account类，
这是因为Item对象与客户账户无关。因此，如果一个类的方法操纵另一个类的对象，我们就说一
个类依赖于另一个类。
应该尽可能将相互信赖的类减至最少。如果类A不知道B的存在，它就不会关心B的任何改变(这意味
着B的改变不会导致A产生任何bug)。用软件工程的术语来说，就是让类之间的耦合度最小。

聚合(aggregation), 即"has-a"关系，是一种具体且易于理解的关系。例如，一个Order对象包含
一些Item对象。聚合关系意味着类A的对象包含类B的对象。

继承(inheritance), 即"is-a"关系，是一种用于表示特殊与一般关系的。例如，RushOrder类由Order
类继承而来。在具有特殊性的RushOrder类中包含了一些用于优先处理的特殊方法，以及一个计算
运费的不同方法；而其他的方法，如添加商品、生成账单等都是从Order类继承来的。一般而言，如
果类A扩展类B，类A不但包含从类B继承的方法，还会拥有一些额外的功能。
