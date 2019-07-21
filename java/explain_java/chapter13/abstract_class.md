### 抽象类
* 无法创建或不应该创建实例
* 无法定义方法的主体。其内容应该在子类中具体实现
```java
abstract class A { // 抽象类
	abstract void aDef();  // 抽象方法
}
```

#### 抽象方法
由于抽象方法中不存在主体，因此在其声明中，我们将{ }替换为;。即使方法的主体为空，
也不可以写为程序块{ }。
```java
abstract void aDef() { }   // err 无法定义
```

> 在从抽象派生的类中，重写抽象方法，定义主体的操作称为`实现(implement)` 方法。

#### 抽象类
只要包含1个抽象方法，该类就必须声明为抽象类。为了将类声明为抽象类，需要在class
的前面加上abstract。

> > 对于抽象类， 不可以指定final、static、private。
```java
class P {				// err 只要包含一个抽象方法，该类就必须声明为抽象类。
	abstract void a();
	void b() { ... }
}

abstract class Q {		// ok
	abstract void a();
	void b() { ... }
}

abstract class R {		// ok, 抽象类中可以没有抽象方法
	void a () { ... }
	void b () { ... }
}
```

#### 无法创建抽象类的实例
由于抽象类中的方法未具体定义，因此无法使用new 来创建实例.

如果从抽象类派生的子类中未实现抽象方法，则抽象方法会被直接继承。
```java
abstract class A {
	abstract void a();
	abstract void b();
}

// 由于继承了抽象方法b,但在B类中未实现，因此B类也是
// 抽象类
abstract class B extends A {  

	void a() { ... }
}

class C extends B {
	void b() { ... }
}

```


#### 具有抽象性的非抽象方法的设计

> 可以将超类的非抽象方法重写为抽象方法。

