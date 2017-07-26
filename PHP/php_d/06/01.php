<?php


# 装饰器设计模式
#
# 装饰器模式就是对一个已有的结构增加“装饰”。装饰器模式会向现有对象增加对象。
# 装饰器也称为包装器（类似于适配器）。
# 有些设计模式包含一个抽象类，而且该抽象类还继承另一个抽象类，这种设计模式
# 为数不多，而装饰器模式就是其中之一。
# 
# 什么时间使用装饰器
# 如果想为现有对象增加新功能而不想影响其他对象，就可以使用装饰器模式。

include_once '../display_errors.php';


abstract class Icomponent
{
	protected $site;
	abstract public function getSite();
	abstract public function getPrice();
}

abstract class Decorator extends IComponent
{
	// 继承 getSite()和getPirce()
	// 这仍是一个抽象类
	// 这里不需要实现任何一个抽象方法
	// 任务是维护Component的引用
	// public function getSite(){}
	// public function getPrice(){}
}

// 具体组件
class BasicSite extends IComponent
{
	public function __construct()
	{
		$this->site = 'Basic Site';
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
		$fmat = "<br>&nbsp;&nbsp; Maintenance ";
		return $this->site->getSite() . $fmat;
	}

	public function getPrice()
	{
		return 950 + $this->site->getPrice();
	}
}

// 查看这个具体构造函数，你会发现，它看起来与装饰器几乎完全相同。不过，每个
// 具体装饰器在它包装的具体组件价格上还会增加它自己的一个价格。


// 具体装饰器2
class Video extends IComponent
{
	public function __construct(IComponent $siteNow)
	{
		$this->site = $siteNow;
	}

	public function getSite()
	{
		$fmat = '<br>&nbsp;&nbsp; Video ';
		return $this->site->getSite() . $fmat;
	}

	public function getPrice()
	{
		return 350 + $this->site->getPrice();
	}
}


// 具体装饰器3
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

# 装饰器实现中最重要的元素之一就是构造函数，要为构造函数提供一个组件类型。
# 由于这里只有一个具体组件，所有装饰器的实例化都会使用这个组件。使用多个
# 组件时，装饰器可以包装应用中的一部分或全部组件，也可以不包装任何组件。


class Client
{
	private $basicSite;

	public function __construct()
	{
		$this->basicSite = new BasicSite();
		$this->basicSite = $this->wrapComponent($this->basicSite);

		$siteShow = $this->basicSite->getSite();
		$format = '<br>&nbsp;&nbsp;<strong>Total= $';
		$price = $this->basicSite->getPrice();

		echo $siteShow . $format . $price . '</strong><p/>';
	}

	private function wrapComponent(IComponent $component)
	{
		$component = new Maintenance($component);
		$component = new Video($component);
		$component = new DataBase($component);
		return $component;
	}
}

$worker = new Client();
