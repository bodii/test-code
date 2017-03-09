<?php

# 状态设计模式

# 游戏通常就采用状态模式，因为游戏中的对象往往会非常频繁地改变状态。
# 状态模式的作用就是允许对象在状态改变时改变其行为。
# 对象中频繁的状态改变非常依赖于条件语句。不过，如果选项太多，以至于
# 程序开始出现混乱，或者增加或改变选项需要花费太多时间，甚至成为一种
# 负担，这就出现了问题。


# 情境为王
# 所有状态模式都需要一个参与者来跟踪对象所处的状态。

include_once "../display_errors.php";

class Context
{
	private $offState;
	private $onState;
	Private $currentState;

	public function __construct()
	{
		$this->offState = new OffState($this);
		$this->onState = new OnState($this);
		// 开始状态为Off
		$this->currentState = $this->offState;
	}

	// 调用状态方法解发器
	public function turnOnLight()
	{
		$this->currentState->turnLightOn();
	}

	public function turnOffLight()
	{
		$this->currentState->turnLightOff();
	}

	// 设置当前状态
	public function setState(IState $state)
	{
		$this->currentState = $state;
	}

	// 获得状态
	public function getOnState()
	{
		return $this->onState;
	}
	
	public function getOffState()
	{
		return $this->offState;
	}
}

// 在这个实例化过程用到了一种递归，称为自引用。构造函数参数中的实参写为this
// ，这是Context类自身的一个引用。状态类希望接收一个ConText类实例作为参数，
// 为了在ConText类中实例化一个状态实例，它必须使用$this作为实参。

interface IState
{
	public function turnLightOn();
	public function turnLightOff();
}

class OnState implements IState
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOn()
	{
		echo "Light is already on-> take no action<br>";
	}

	public function turnLightOff()
	{
		echo "Light off!<br>";
		$this->context->setState($this->context->getOffState());
	}
}

class OffState implements IState
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOn()
	{
		echo "Lights on! Now I can see!<br>";
		$this->context->setState($this->context->getOnState());
	}

	public function turnLightOff()
	{
		echo "Light is alread off-> take no action<br>";
	}

}

class Client
{
	private $context;

	public function __construct()
	{
		$this->context = new Context();
		$this->context->turnOnLight();
		$this->context->turnOnLight();
		$this->context->turnOffLight();
		$this->context->turnOffLight();
	}
}

$worker = new Client();
