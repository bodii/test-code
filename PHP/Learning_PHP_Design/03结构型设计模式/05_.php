<?php 

/*
	最简单的装饰器例子

*/

abstract class IComponent
{
	protected $site;
	abstract public function getSite();
	abstract public function getPrice();
}

abstract class Decorator extends IComponent
{
	public function getSite(){}
	public function getPrice(){}
}

class BasicSite extends IComponent
{
	public function __construct()
	{
		$this->stie = 'Basic Site';
	}

	public function getSite()
	{
		return $this->site;
	}

	public function getPrice()
	{
		return 1200;
	}
}

// 具体装饰器
class Maintenance extends Decorator
{
	public function __construct(IComponent $siteNow)
	{
		$this->site = $siteNow;
	}

	public function getSite()
	{
		$fmat = '<br>&nbsp;&nbsp;Maintenance ';
		return $this->site->getSite() . $fmat;
	}

	public function getPrice()
	{
		return 950 + $this->site->getPrice();
	}
}

// 具体装饰器
class Video extends IComponent
{
	public function __construct(IComponent $siteNow)
	{
		$this->site = $siteNow;
	}

	public function getSite()
	{
		$fmat = '<br>&nbsp;&nbsp;Video ';
		return $this->site->getSite() . $fmat;
	}

	public function getPrice()
	{
		return 350 + $this->site->getPrice();
	}
}

// 具体装饰器
class DataBase extends Decorator
{
	public function __construct(IComponent $siteNow)
	{
		$this->site = $siteNow;
	}

	public function getSite()
	{
		$fmat = '<br>&nbsp;&nbsp; MySQL DataBase ';
		return $this->site->getSite() . $fmat;
	}

	public function getPrice()
	{
		return 800 + $this->site->getPrice();
	}
}

class Client
{
	private $basicSite;

	public function __construct()
	{
		$this->basicSite = new BasicSite();
		$this->basicSite = $this->wrapComponent($this->basicSite);

		$siteShow = $this->basicSite->getSite();
		$format = '<br>&nbsp;&nbsp;<strong>Total = $';
		$price = $this->basicSite->getPrice();

		echo $siteShow . $format . $price . '<strong><p/>';
	}

	private function wrapComponent(IComponent $component)
	{
		$component = new Maintenance($component);
		$component = new Video($component);
		$component = new DataBase($component);
		return $component;
	}
}

$worker = new Client;