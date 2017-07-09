<?php 

/*
	增加状态

	“关”(off)状态只能变到“开”(on)状态，on状态不能变到off状态。实际上，现在
	规则有变化，on状态只能变到"更亮"(brighter)状态和“最亮”(brightest)状态。
	只有brightest状态可变到“关”(off)状态。
*/

// 改变接口

interface IState
{
	public function turnLightOn();
	public function turnBrighter();
	public function turnBrightest();
	public function turnLightOff();
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
		echo '<img src="lights/on.png" alt="on"><br>';
		$this->context->setState($this->context->getOnState());
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnBrightest()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
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
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/on.png" alt="on"><br>';
		$this->context->setState($this->context->getOnState());
	}

	public function turnBrightest()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
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
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
		
	}
	public function turnBrightest()
	{
		echo '<img src="lights/on.png" alt="on"><br>';
		$this->context->setState($this->context->getOnState());
	}

	public function turnLightOff()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
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
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnBrighter()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
		
	}
	public function turnBrightest()
	{
		echo '<img src="lights/nada.png" alt="nada"><br>';
	}

	public function turnLightOff()
	{
		echo '<img src="lights/on.png" alt="on"><br>';
		$this->context->setState($this->context->getOnState());
	}
}


class Context
{
	private $offState;
	private $onState;
	private $brighterState;
	private $brightestState;

	private $currentState;

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

	public function setState(IState $state)
	{
		$this->currentState = $state;
	}

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
		$this->context = new Context;
		$this->context->turnOnLight();
		$this->context->turnBrighterLight();
		$this->context->turnBrightestLight();
		$this->context->turnOffLight();
	}
}

$worker = new Client;