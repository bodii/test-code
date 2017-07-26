<?php 

/*
	【 适配器模式 】

	什么是适配器模式
	类适配器设计模式使用继承。
	这个模式实现中一个类有双重继承。双重继承在PHP5中是不允许的，不过实现双重
	继承有一些替代方法，可以结余继承和实现来正确地实现这个模式。

	一般来说，组合要优先于继承，因为参与者之间的绑定更宽松，在重用、结构和修
	改等方面有很多优点，这与使用继承不同，继承具体类或者所继承的类中包含已实
	现的方法时，存在一种紧密绑定，使用组合就没有这种紧密绑定的缺点。


	何时使用适配器模式
	分解这个问题可以得到：
	USB连接头--- 不变
	标准墙上插座--- 不变

 */


/*
	使用继承的适配器模式
	类适配器设计模式很简单，不过与对象适配器模式相比，类适配器模式的灵活性稍
	弱。类适配器模式更简单的原因在于，适配器(Adapter)会从被适配器继承功能，所
	以适配器模式中需要重新编写的代码较少。当然，由于给定了一个将由适配器继承
	的具体被适配者，这种绑定很紧密，所以使用类适配器模式模式创建应用时，必须
	非常清楚将在哪里发生适配。

 */


// 货币兑换(计算总额)
class DollarCalc
{
	private $dollar;
	private $product;
	private $service;
	public $rate=1;

	public function requestCalc($productNow, $serviceNow)
	{
		$this->product = $productNow;
		$this->service = $serviceNow;
		$this->dollar = $this->product * $this->service;
		return $this->requestTotal();
	}

	public function requestTotal()
	{
		$this->dollar *= $this->rate;
		return $this->dollar;
	}
}

// 加入欧元

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
		$this->euro = $this->product * $this->service;
		return $this->requestTotal();
	}

	public function requestTotal()
	{
		$this->euro *= $this->rate;
		return $this->euro;
	}
}



/*所有数据都是按美元计算的。为了加入EuroCalc,你需要一个适配器。*/

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

	public function requester()
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
		$this->requestNow = new EuroAdapter;
		$this->dollarRequest = new DollarCalc;
		$euro = "&#8364;";
		echo "Euros: $euro". $this->makeAdapterRequest($this->requestNow) . '<br>';
		echo "Dollars: $". $this->makeDollarRequest($this->dollarRequest);
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

$worker = new Client;