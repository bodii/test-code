<?php 
	
/*
	自动职责链和工厂方法 
	职责链与工厂方法模式结合。

	职责链和日期驱动请求
	链接到这个应用的地址后，会通过一个index.php文件自动启动client。client
	再把当前日期放在一个变量$queryNow中，这个变量将传递到一个Request辅助类。
	沿着职责链，查询会查找第一个合适的具体处理器作为工厂方法的客户，由它创
	建和显示图像和文字素材。
*/
class Client
{
	private $queryNow;
	private $dateNow;

	public function __construct()
	{
		// 得到所选时区的日期
		$tz = 'America/New_York';
		date_default_timezone_set($tz);
		$this->queryNow = getdate();

		$d1 = new D1;
		$d2 = new D2;
		$d3 = new D3;
		$d4 = new D4;
		$d5 = new D5;

		$d1->setSuccessor($d2);
		$d2->setSuccessor($d3);
		$d3->setSuccessor($d4);
		$d4->setSuccessor($d5);

		// 生成和处理加载请求
		$loadup = new Request($this->queryNow);
		$d1->handleRequest($loadup);
	}
}

class Request
{
	private $value;

	public function __construct($service)
	{
		$this->value = $service;
	}

	public function getService()
	{
		return $this->value;
	}
}

abstract class Handler 
{
	protected $hungerFactory;
	protected $successor;
	protected $monthNow;
	protected $dayNow;
	protected $handleNow;
	abstract public function handleRequest($request);
	abstract public function setSuccessor($nextService);
}

class D1 extends Handler
{
	public function setSuccessor($nextService)
	{
		$this->successor = $nextService;
	}

	public function handleRequest($request)
	{
		$dateCheck = $request->getService();
		$this->monthNow = intval($dateCheck['mon']);
		$this->dayNow = intval($dateCheck['mday']);

		// $this->handleNow是一个布尔值
		// 基于一个带日期范围的布尔表达式得到
		$this->handleNow = ($this->monthNow == 9 && $this->dayNow >= 15) && 
		($this->monthNow == 9 && $this->dayNow <= 22);

		if ($this->handleNow)
		{
			$this->hungerFactory = new HungerFactory;
			echo $this->hungerFactory->feedFactory(new C1());
		}
		elseif ($this->successor != null)
		{
			$this->successor->handleRequest($request);
		}
	}
	/*
		通过使用一个布尔变量($handleNow)和布尔表达式，处理器首先查询
		$handleNow，如果为true,则指向一个工厂方法，由它加载必要的素材
		从这个角度来看，职责链具体处理器(D1到D15)是一个客户，将从工厂
		方法做出一个请求。
	 */
}

/* 工厂方法完成任务 */

// Creator和HungerFactory

abstract class Creator
{
	protected $countryProduct;
	protected abstract function factoryMethod(Product $product);

	public function feedFactory(Product $productNow)
	{
		$this->countryProduct = $productNow;
		$mfg = $this->factoryMethod($this->countryProduct);
		return $mfg;
	}
}

class HungerFactory extends Creator
{
	private $country;

	protected function factoryMethod(Product $product)
	{
		$this->country = $product;
		return  $this->country->getProperties();
	}
}

// 最终结果显示了一web页面，其中包含一个照片、一个地图和一段简短的文学说明
// 学生们使用这个应用研究全球饥荒情况，必须找出国家名、婴儿死亡率和不同文化
// 程度的性别差异

class C1 implements Product 
{
	private $mfgProduct;
	private $formatHelper;
	private $countryNow;

	public function getProperties()
	{
		// 从外部文本文件加载文体说明
		$this->countryNow = file_get_Contents('../hunger/c01/clue.txt');

		$this->formatHelper = new FormatHelper;
		$this->mfgProduct = $this->formatHelper->addTop();

		// 加载图像
		$this->mfgProduct .= '<img src="../hunger/c01/map.gif" width="300"
		height="322">';
		$this->mfgProduct .= '<img class="pixLeft" src="../hunger/c01/pic.jpg
		 width="200" height="400">';
		$this->mfgProduct .= "<p>{$this->countryNow}</p>";
		$this->mfgProduct .= $this->formatHelper->closeUp();
		return $this->mfgProduct; 
	}
}

class FormatHelper
{
	private $topper;
	private $bottom;

	public function addTop()
	{
		$this->topper =<<<HEAD
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Map Factory</title>
	</head>
	<body>
HEAD;
	}

	public function closeUp()
	{
		$this->bottom = '<br><a href="Answer.html" target="_blank">Click for 
		Credit</a>';
		$this->bottom .= '</body></html>';
	}
}