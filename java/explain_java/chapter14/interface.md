### 接口
如果将类比作"电路的设计图", 那么接口(interface)就是"遥控器的设计图".
所谓interface,就是"交界面", "公共区域"的意思。

#### 接口的声明
```java
// 播放器 接口
interface Player {
	void play();
	void stop();
}
```

接口声明(interface declaration).开头的关键字并不是class,而是inteface.

此外，接口的所有方法都为public且abstract(也可以在声明中加上public和abstract,但这只会让程序
变得冗长).
必须用;来替换方法体{ ... } 进行声明，这一点与类中的抽象方法相同。


#### 接口的实现
实现(implement)该接口。
```java
// 该类会实现接口Player
class VideoPlayer implements Player {
	public void play() {
		// 方法的定义
	}

	public void stop() {
		// 方法的定义
	}
}
```

重写的方法必须声明为public.这是因为接口的方法为public.已无法再强化其访问控制。
> 接口中的方法为public且abstract.在实现该接口的类中，各方法在实现时需要加上public修饰符。


#### 无法创建接口类型的实例
#### 接口类型的变量可以引用实现该接口的类的实例
```java
Player p1 = new CDPlayer();
Player p2 = new VideoPlayer();
```

#### 接口实现时必须实现所有的方法
```java
interface I {
	void a();
	void b();
}

abstract class A implements I {			// 实现方法a。由于未实现方法b,因此该类为抽象类
	public void a() { ... }
}

class B extends A {
	public void b() { ... }
}
```

#### 可以持有常量
接口可以持有下述成员:
* 类
* 接口
* 常量 ... public、 static 、final字段
* 抽象方法  ... public、abstract 方法
> 接口可以持有常量成员，但不可以持有非常量字段。接口与“只持有抽象方法和常量成员的抽象类类似"。

```java
// 换肤接口
public interface Skinnable {
	/* 字段为public static final */
	int BLACK = 0;
	int RED   = 1;
	int GREEN = 2;
	int BLUE  = 3;
	int LEOPARD = 4;

	/* 方法为public abstract */
	void changeSkin(int skin);		// 换肤
}
```

> 接口中声明的字段为public且static且final，即为不可以改写数值的类变量。
声明中带有static的类变量可以通过"类名.字段名"时行访问，与此同时，接口中的常量可以通过"接口名.字段名"
进行访问.
在本示例中，黑色可以通过Skinnable.BLACK进行访问。

#### 命名方法与类相同
原则上，接口的名称与类一样使用名词，不过，像本示例这种表示"可....的"之意的接口名称可以使用~able.
