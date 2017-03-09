<?php

# MVC实现编程松耦合和重新聚焦

include_once '../display_errors.php';

abstract class IAbstract
{
	// 对所有实现都可用的属性
	protected $valueNow;

	/* 所有实现都必须包含以下两个方法：*/
	// 必须返回十进制值
	abstract protected function giveCost();
	// 必须返回字符串值
	abstract protected function giveCity();

	// 这个具体函数对所有类实现都可用， 而不覆盖内容
	public function displayShow()
	{
		$stringCost = $this->giveCost();
		$stringCost = (string)$stringCost;
		$allTogether = ("Cost: $" . $stringCost . " for " . $this->giveCity());
		return $allTogether;
	}
}


class NorthRegion extends IAbstract
{
	//必须返回十进制值
	protected function giveCost()
	{
		return 210.54;
	}

	//必须返回字符串值
	protected function giveCity()
	{
		return 'Moose Breath';
	}

}


class WestRegion extends IAbstract
{
	// 必须返回十进制值
	protected function giveCost()
	{
		$solarSavings = 2;
		$this->valueNow = 210.54 / $solarSavings;
		return $this->valueNow;
	}

	// 必须返回字符串值
	protected function giveCity()
	{
		return 'Rattlesnake Gulch';
	}

}

class Client
{
	public function __construct()
	{
		$north = new NorthRegion();
		$west = new WestRegion();
		$this->showInterface($north);
		$this->showInterface($west);
	}

	private function showInterface(IAbstract $region)
	{
		echo $region->displayShow() . '<br>';
	}

}

$worker = new Client();
