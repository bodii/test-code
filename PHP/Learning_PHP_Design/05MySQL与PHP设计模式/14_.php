<?php 
	
/*
	自由的PHP和观察者模式  

	使用用非SPL参与者,构建观察者模式
	使用SPL和不使用SPL二者的主要区:
		1.Subject参与者是一个抽象类而不是一个接口
		2.notify()方法在Subject中实现而不是在具体主题中实现.
		3.具体主题使用一个数组对象而不是一个SplObjectStorage对象来存储观察者
		实例.
*/

abstract class Subject 
{
	protected $stateNow;
	protected $observers = [];

	public function attachObser(Observer $obser)
	{
		if (in_array($obser, $this->observers) === false)
		{
			$this->observers[] = $obser;
		}
	}

	public function detachObser(Observer $obser)
	{
		if ($key = array_search($obser, $this->observers))
		{
			unset($this->observers[$key]);
		}
	}

	public function notify()
	{
		foreach($this->observers as $observer)
		{
			$observer->update($this);
		}
	}
}

class ConcreteSubject extends Subject 
{
	public function setState($stateSet)
	{
		$this->stateNow = $stateSet;
		$this->notify();
	}

	public function getState()
	{
		return $this->stateNow;
	}
}

/*
	在detachObser()方法中可以看到，必须迭代处理$observers数组来通知多个观察
	者。当然，即使没有使用SPL主题和观察者，也可以使用SplObjectStorage类。这
	样一来，可以使用一个SplObjectStorage对象（而不是一个数组）保存关联的观察
	者。另外，还可以使用SplObjectStorage类中内置的attach/detach方法。
 */

interface Observer
{
	function update(Subject $subject);
}

// 桌面计算机实现
class ConcreteObserverDT implements Observer
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='desktop/{$this->currentState}'><br>";
	}
}

// 平板电脑实现
class ConcreteObserverTablet implements Observer 
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='tablet/{$this->currentState}'><br>";
	}
}

// 智能手现实现
class ConcreteObserverPhone implements Observer
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='phone/{$this->currentState}'><br>";
	}
}


// 客户
class Client
{
	public function __construct()
	{
		$sub = new ConcreteSubject;

		$ob1 = new ConcreteObserverPhone;
		$ob2 = new ConcreteObserverTablet;
		$ob3 = new ConcreteObserverDT;

		$sub->attachObser($ob1);
		$sub->attachObser($ob2);
		$sub->attachObser($ob3);
		$sub->setState('decoCar.png');
	}
}

$worker = new Client;