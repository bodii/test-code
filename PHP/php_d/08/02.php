<?php

# 增加状态
# 关(Off) 开(On) 更亮(Brighter) 最亮(Brightest)
# {关} 只能变到 {开}， {开}不能直接变到{关}。{开} 只能变到{更亮}和{最亮}，
# 只有{最亮} 可以变到{关}

interface IState
{
	public function turnLightOff();
	public function turnLightOn();
	public function turnBrighter();
	public function turnBrightest();
}

// 状态
class OffState implements IState
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOn()
	{
		echo '<img src="lights/on.png">';
		$this->context->setState($this->context->getOnState());
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnBrightest()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png">';
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
		echo '<img src="lights/nada.png">';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/brighter.png">';
		$this->context->setState($this->context->getBrighterState());
	}

	public function turnBrightest()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png">';
	}
}

class BrighterState implements IState
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOn()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnBrightest()
	{
		echo '<img src="lights/brightest.png">';
		$this->context->setState($this->context->getBrightestState());
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png">';
	}
}

class BrightestState implements IState
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function turnLightOn()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnBrightest()
	{
		echo '<img src="lights/nada.png">';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/off.png">';
		$this->context->setState($this->context->getOffState());
	}
}

class Context
{
	private $offState;
	private $onState;
	private $brighterState;
	private $brightestState;

	private $currnetState;

	public function __construct()
	{
		$this->offState = new OffState($this);
		$this->onState = new OnState($this);
		$this->brighterState = new BrighterState($this);
		$this->brightestState = new BrightestState($this);

		// 开始状态为Off
		$this->currentState = $this->offState;
	}

	// 调用状态方法
	public function turnOnLight()
	{
		$this->currentState->turnLightOn();
	}

	public function turnOffLight()
	{
		$this->currentState->turnLightOff();
	}

	public function turnBrighterLight()
	{
		$this->currentState->turnBrighter();
	}

	public function turnBrightestLight()
	{
		$this->currentState->turnBrightest();
	}

	// 设置当前状态
	public function setState(IState $state)
	{
		$this->currnetState = $state;
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

	public function getBrighterState()
	{
		return $this->brighterState;
	}

	public function getBrightestState()
	{
		return $this->brightestState;
	}
}

class Client
{
	private $context;

	public function __construct()
	{
		$this->context = new Context();
		$this->context->turnOnLight();
		$this->context->turnBrighterLight();
		$this->context->turnBrightestLight();
		$this->context->turnOffLight();
		$this->context->turnBrightestLight();
	}
}

$worker = new Client();


