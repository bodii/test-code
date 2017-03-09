<?php


# 自由的PHP和观察者模式
#
# 抽象Subject类和ConcreteSubject实现

abstract class Subject
{
	protected $stateNow;
	protected $observers = array();

	public function attachObser(Observer $obser)
	{
		array_push($this->observers, $obser);
	}

	public function detachObser(Observer $obser)
	{
		$position = 0;
		foreach($this->observers as $viewer)
		{
			++$position;
			if($viewer == $obser)
			{
				array_splice($this->observers, ($position), 1);
			}
		}
	}

	protected function notify()
	{
		foreach($this->observers as $viewer)
		{
			$viewer->update($this);
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

interface Observer
{
	function update(Subject $subject);
}

// 用于桌面
class ConcreteObserverDT implements Observer
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='desktop/$this->currentState'><br>";
	}
}

// 用于平板
class ConcreteObserverTablet implements Observer
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='tablet/$this->currentState'><br>";
	}
}

// 智能手机
class ConcreteObserverPhone implements Observer
{
	private $currentState;

	public function update(Subject $subject)
	{
		$this->currentState = $subject->getState();
		echo "<img src='phone/$this->currentState'><br>";
	}
}

class Client
{
	public function __construct()
	{
		$sub = new ConcreteSubject();

		$ob1 = new ConcreteObserverPhone();
		$ob2 = new ConcreteObserverTablet();
		$ob3 = new ConcreteObserverDT();

		$sub->attachObser($ob1);
		$sub->attachObser($ob2);
		$sub->attachObser($ob3);
		$sub->setState("decoCar.png");
	}
}

$worker = new Client();
