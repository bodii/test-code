<?php 

/*
	开灯关灯：最简单的状态设计模式

	% 所有状态模式都需要一个参与者来跟踪对象所处的状态。%
*/

class Context
{
	private $offState;
	private $onState;
	private $currentState;

	public function __construct()
	{
		$this->offState = new OffState($this);
		$this->onState = new OnState($this);

		// 开始状态为Off
		$this->currentState = $this->offState;
	}

	// 调用状态方法触发器
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

interface IState 
{
	public function turnLightOff();
	public function turnLightOn();
}

class OffState implements IState 
{
	private $context;
	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOff()
	{
		echo 'Light is already off-> take on action<br>';
	}

	public function turnLightOn()
	{
		echo 'Lights on! Now I can see!<br>';
		$this->context->setState($this->context->getOnState());
	}
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
		echo 'Light is already on-> take on action<br>';
	}

	public function turnLightOff()
	{
		echo 'Lights off!<br>';
		$this->context->setState($this->context->getOffState());
	}
}


class Client
{
	private $context;

	public function __construct()
	{
		$this->context = new Context;
		$this->context->turnOnLight();
		$this->context->turnOnLight();
		$this->context->turnOffLight();
		$this->context->turnOffLight();

	}
}

$worker = new Client;