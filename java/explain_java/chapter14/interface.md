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


