<?php


# 模板方法设计模式中的钩子
# 在模板方法设计模式中，利用钩子可以将一个方法作为模板方法的一部分，不过不
# 一定会用到这个方法。换句话说，它是方法的一部分，不过它包含一个钩子，可以
# 处理例外情况。子类可以为算法增加一个可选元素，这样一来，尽管仍按模板方法
# 建立的顺序执行，但有可能并不完成模板方法期望的动作。

abstract class IHook
{
	protected $purchased;
	protected $hookSpecial;
	protected $shippingHook;
	protected $fullCost;

	public function templateMethod($total, $special)
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
		if (! $this->hookSpecial)
		{
			$this->fullCost += 12.95;
		}
	}

	protected function displayCost()
	{
		echo "Your full cost is $this->fullCost";
	}
}
