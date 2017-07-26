<?php

# 使用SPL实现观察者模式
#
# 可用于观察者设计模式的3个SPL接口/类如下：
# SplSubject
# SplObserver
# SplObjectStorage
#
#
# SplSubject
# SplSubject接口有3个抽象方法，这与原来观察者类图中的方法类似。不过，SplSubject
# 是一个接口，而原来Subject参与者通常是一个抽象类，因为它的方法不是抽象的。
#
# 《PHP Manual》给出了SplSubject的方法，如下所示：
# abstract public void attach ( SplObserver $observer )
# abstract public void detach ( SplObserver $observer )
# abstract public void notify ( void )
#
# 从void可以看出，这些方法不返回任何结果。标准PHP接口中还要包含一个function
# 关键字。如果自行创建这个接口，可能实现如下：
#
# public function attach( SplObserver $observer );
# public function detach( Splobserver $observer );
# public function notify();
#
# 实现SplSubject接口时无需创建一个subject接口或抽象类。重要的是，这个接口
# 指定attach()和detach()方法参数中$observer的数据类型必须是一个SplObserver
# 对象。
#
# 
# SplObserver
# SplObserver 接口有一个方法update(),如下所示：
# abstract public void update ( SplSubject $subject )
#
# 采用标准PHP接口格式，可以实现为：
#
# public function update(SplSubject $subject);
#
# update()方法对于观察者模式至关重要，因为它会得到Subject状态的最新变化，
# 并交给观察者实例。 
#
# SqlObjectStorage
# SplObjectStorage类与观察者设计模式没有内在的关系，不过通过它可以很方便
# 地将观察者实例与一个主题相关联以及解除关联。尽管一般将SqlObjectStorage
# 类描述为从对象到数据或对象集的一映射，但我更喜欢把它想成是一个数组（
# 带有内置的attach和detach()方法)。这个类提供了一种简单的方法，可以使观察
# 者与发出状态改变通知的主题对象关联以及解除关联。
#
# SPL具体主题
# 主题需要与观察者关联和解除关联，并通知订阅者（关联的观察者）发生了改变。
# 私有变量$observers封装了这个属性。在这个实现中，$observers属性实例化为一
# 个SplObjectStorage对象。可以把它想成是一个数组，然后关联各个$observer实例
# （SplObserverc对象），存储单元是存储各个$observer实例的$observers。
#
# SplObjectStorage类最重要的方法是其内置attach()和detach()方法，利用这两个
# 方法，可以很容易地指定哪些观察者实例将“订阅”和“解除订阅”更新通知：
#

class ConcreteSubject implements SplSubject
{
	private $observers;
	private $data;

	public function setObservers()
	{
		$this->observers = new SplObjectStorage();
	}

	public function attach(SplObserver $observer)
	{
		$this->observers->attach($observer);
	}

	public function detach(SplObserver $observer)
	{
		$this->observers->detach($observer);
	}

	public function notify()
	{
		foreach ($this->observers as $observer) {
			$observer->update($this);
		}
	}

	public function setData($dataNow)
	{
		$this->data = $dataNow;
	}

	public function getData()
	{
		return $this->data;
	}

}

# SplSubject接口不包括获取方法和设置方法，不过这是观察者设计模式中的一部分，
# 所以需要增加获取方法和设置方法。设置方法setData()包含一个参数，这是要增加
# 的任何类型的数据。获取方法getData存储当前主题状态，由具体观察者用来更新
# 观察者数据。
# 另外还增加了setObservers()方法。并不是在构造函数中设置SplObjectStorage()
# 实例（这需要ConcreteSubject类的一个新实例），也没有在setData()方法中设置
# 观察者实例，这里实现了一个单独的setObservers()方法，可以提供更松的耦合，
# 并允许有多组观察者。
