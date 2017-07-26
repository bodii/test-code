<?php


# 动态对象实例化

# 结构型设计模式
# 结构型设计模式研究的是如何组合对象和类来构成更大的结构。
/*
 下面7种模式称为结构型模式：
 适配器模式（类和对象）
 桥接模式
 组合模式
 装饰器模式
 外观模式
 享元模式
 代理模式
 */
# 结构型设计模式的重点是创建新结构而不破坏原有的结构。在这个基础上，结构型
# 模式可以保持并提升松耦合标准以实现重用和灵活改变。



# 适配器模式


# 这里不仅继承了一个类，同时还实现了一个接口
# 如：class ChildClass extends ParentClass implements ISomeInterface
#
# 类适配器的例子

include_once '../display_errors.php';
// 美元
class DollarCalc
{
	private $dollar;
	private $product;
	private $service;
	public $rate = 1;

	public function requestCalc($productNow, $serviceNow)
	{
		$this->product = $productNow;
		$this->service = $serviceNow;
		$this->dollar = $this->product + $this->service;
		return $this->requestTotal();
	}

	public function requestTotal()
	{
		$this->dollar *= $this->rate;
		return $this->dollar;
	}
}

//欧元
class EuroCalc
{
	private $euro;
	private $product;
	private $service;
	public $rate = 1;

	public function requestCalc($productNow, $serviceNow)
	{
		$this->product = $productNow;
		$this->service = $serviceNow;
		$this->euro = $this->product + $this->service;
		return $this->requestTotal();
	}

	public function requestTotal()
	{
		$this->euro *= $this->rate;
		return $this->euro;
	}
}

// 创建一个欧元换算美元的适配器
// 接口
interface ITarget
{
	function requester();
}

class EuroAdapter extends EuroCalc implements ITarget
{
	public function __construct()
	{
		$this->requester();
	}

	function requester()
	{
		$this->rate = .8111;
		return $this->rate;
	}
}



class Client
{
	private $requestNow;
	private $dollarRequest;

	public function __construct()
	{
		$this->requestNow = new EuroAdapter();
		$this->dollarRequest = new DollarCalc();
		$euro = "&#8364;";
		echo "Euros: $euro" . $this->makeAdapterRequest($this->requestNow)
			. '<br>';
		echo 'Dollars: $'. $this->makeDollarRequest($this->dollarRequest);
	}

	private function makeAdapterRequest(ITarget $req)
	{
		return $req->requestCalc(40, 50);
	}

	private function makeDollarRequest(DollarCalc $req)
	{
		return $req->requestCalc(40, 50);
	}
}

$worker = new Client();

