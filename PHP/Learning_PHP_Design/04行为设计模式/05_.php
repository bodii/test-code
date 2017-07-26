<?php 

/*
	模板方法设计模式中的钩子

	有时模板方法函数中可能有一个你不想要的步骤，某些特定情况下你可能不希望执行
	这个步骤。这里就可以用到模板方法钩子。

	在模板方法设计模式中，利用钩子可以将一个方法作为模板方法的一部分，不过不一
	定会用到这个方法。换句话说，它是方法的一部分，不过它包含一个钩子，可以处理
	例外情况。子类可以为算法增加一个可选元素，这样一来，尽管仍按模板方法建立的
	顺序执行，但有可能并不完成模板方法期望的动作。

	你可能为为，这与好莱坞原则有冲突（子类没有遵循父类设置的顺序），这样想也没
	错。好莱坞原则要求只有父类能够改变框架。钩子很特殊，因为尽管它实现了模板方
	法中的方法，但实现的方法有一个“后门”，也就是说，它会处理例外情况。
*/


/* 【 建立钩子 】 */

abstract class IHook
{
	protected $purchased;
	protected $hookSpecial;
	protected $shippingHook;
	protected $fullCost;

	public function templateMehtod($total, $special)
	{
		$this->purchased = $total;
		$this->hookSpecial = $special;
		$this->addTax();
		$this->addShippingHook();
		$this->displayCost();
	}

	protected abstract function addTax();
	protected abstract function addShippingHook();
	protected abstract function displayCost();

	/* 
	这里，钩子方法放在中间，实际上模板方法指定的顺序中，钩子可以放在任意位
	置。
	 */
}

// 实现钩子

class ZambeziCalc extends IHook
{
	protected function addTax()
	{
		$this->fullCost = $this->purchased + ($this->purchased * .07);
	}

	protected function addShippingHook()
	{
		// 这就是钩子
		if (! $this->hookSpecial)
		{
			$this->fullCost += 12.95;
		}
	}

	protected function displayCost()
	{
		echo "Your full cost is {$this->fullCost}";
	}
}

/*
	用比较操作符设置布尔值
	不必使用条件语句来建立一个布尔状态，使用比较操作符会更容易，也更简洁。
	例如： $this->special = ($this->buyTotal >= 200);
	这会为$this->special赋一个状态true或false;
 */

class Client
{
	private $buyTotal;
	private $gpsNow;
	private $mapNow;
	private $boatNow;
	private $special;
	private $zamCalc;

	public function __construct()
	{
		$this->setHTML();
		$this->setCost();
		$this->zamCalc = new ZambeziCalc;
		$this->zamCalc->templateMehtod($this->buyTotal, $this->special);
	}

	private function setHTML()
	{
		$this->gpsNow = $_POST['gps'];
		$this->mapNow = $_POST['map'];
		$this->boatNow = $_POST['boat'];
	}

	private function setCost()
	{
		$this->buyTotal = $this->gpsNow;
		foreach($this->mapNow as $value)
		{
			$this->buyTotal += $value;
		}

		$this->special = ($this->buyTotal >= 200);
		$this->buyTotal += $this->boatNow;
	}
}

$worker = new Client;