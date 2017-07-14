<?php 

/*
	行为型设计模式

	备忘录模式(Memento Pattern)
	状态模式(State Pattern)
*/

/*
	备忘录模式
	1，有时一些发起人对象的内部信息必须保存在发起人对象以外的地方，但是必须要由发起人
	对象自已读取，这时，使用备忘录模式可以把复杂的发起人内部信息对其他的对象屏蔽起
	来，从而可以恰当地保持封装的边界。
	
	2，本模式简化了发起人类。发起人不再需要管理和保存其内部状态的一个个版本，客户端可
	以自行管理他们所要的这些状态的版本。
*/


// 发起人，所需备份者
class Originator
{
	// 内部状态
	private $state;

	// 设置状态
	public function setState($state)
	{
		$this->state = $state;
	}

	// 查看状态
	public function getState()
	{
		return $this->state;
	}

	// 创建一个备忘录
	public function createMemento()
	{
		return new Memento($this->state);
	}

	// 恢复一个备份
	public function restoreMemento(Memento $memento)
	{
		$this->state = $memento->getState();
	}
}

// 备忘录角色
class Memento
{
	private $state;

	public function __construct($state)
	{
		$this->state = $state;
	}

	public function getState()
	{
		return $this->state;
	}
}

// 备忘录管理者
class Caretaker
{
	private $memento;

	public function setMemento(Memento $memento)
	{
		$this->memento = $memento;
	}

	public function getMemento()
	{
		return $this->memento;
	}
}


$role = new Originator;
$role->setState('满血');

$caretaker = new Caretaker;
$caretaker->setMemento($role->createMemento());

$role->setState('死亡');
$role->getState();

$role->restoreMemento($caretaker->getMemento());

$role->getState();