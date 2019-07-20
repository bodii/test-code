### 继承和访问属性

#### 成员
在类的派生中，只有类的成员(member)可以被继承，类的成员如下：
* 字段
* 方法
* 类
* 接口

原则上，超类中的成员会被直接继承，不过，具有私有访问属性的成员，即声明为private的
成员不会被继承。

类中非成员资产:
* 实例初始化器
* 静态初始化器
* 构造函数


#### final类和方法
声明为final的类和方法在派生中会被特殊处理。

##### final类
不可以从final类中进行派生 。也就是说，无法创建以final类为超类的类。
```java
public final class String {
}

因此， 不可以像下面这样，创建String类的扩展类:
```java
class DeluxeString extends String {
}
```

##### final方法
在方法的开头上final进行声明的final方法不可以被子类重写：
```java
final void f() { ... }    // 该方法不可以重写
```
> 不应该被子类重写的方法要声明为final方法

> 另外，final类中的所有方法都会自动变为final方法。
final有“最后的”的意思，finanl类和final方法就蕴含“最终确定的版本，已经无法再扩展、
重写”的意思。


#### 重写和方法的访问属性
> 当重写方法时，必须赋给与上位类中的方法相同或更弱的访问控制修饰符。
```java
public class A {
	public		void m1() { }  // 公开
	protected	void m2() { }  // 限制公开
				void m3() { }  // 包（默认、无修饰符）
	private		void m4() { }  // 私有
}
```
即使下位类中定义了与私有方法具有相同签名、相同返回类型的方法，也不会被看作是重写，
只是碰巧具有相同规格但毫无关系的方法而已。

在Java的所有类的“老大类”Object类中，toString被定义为public方法。因此，当重写toString
类时，一定要加上public。
> 方法Sring toString() 必须在定义中加上public修饰符。

不可以将超类的类方法重写为实例方法:
```java
class A {
	static void f() { ... }
}

class B extends A {
	void f() { ... } // error
}
```

#### 声明中的修饰符顺序
类		注释 public protected private abstract static final strictfp
字段	注释 public protected private static	final transient volatile
方法	注释 public protected private abstract static final synchronized native strictfp
接口	注释 public protected private abstract static strictfp
