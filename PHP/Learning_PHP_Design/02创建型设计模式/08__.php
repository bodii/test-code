<?php 

/*
	【 工厂方法设计模式 】
	工厂方法（Factory Method）模式就是要创建“某种东西”。
	对工厂方法模式，要创建的“东西”是一个产品，这个产品与创建它的类之间不
	存在绑定。实际上，为了保持这种松耦合，客户会通过一个工厂发出请求。再
	由工厂创建所请求的产品。也可以换种方式考虑，利用工厂方法模式，请求者
	只发出请求，而不具体创建产品。
 */

abstract class Creator
{
	protected abstract function factoryMethod();

	public function startFactory()
	{
		$mfg = $this->factoryMethod();
		return $mfg;
	}
}

interface Product 
{
	public function getProperties();
}

class TextProduct implements Product
{
	private $mfgProduct;

	public function getProperties()
	{
		$this->mfgProduct = 'This is text.';
		return $this->mfgProduct;
	}
}

class GraphicProduct implements Product
{
	private $mfgProduct;

	public function getProperties()
	{
		$this->mfgProduct = 'This is a Graphic.';
		return $this->mfgProduct;
	} 
}

class TextFactory extends Creator
{
	protected function factoryMethod()
	{
		$product = new TextProduct();
		return $product->getProperties();
	}
}

class GraphicFactory extends Creator
{
	protected function factoryMethod()
	{
		$product = new GraphicProduct();
		return $product->getProperties();
	}
}

class Client
{
	private $someGraphicObject;
	private $someTextObject;

	public function __construct()
	{
		$this->someGraphicObject = new GraphicFactory;
		echo $this->someGraphicObject->startFactory() .'<br>';
		$this->someTextObject = new TextFactory;
		echo $this->someTextObject->startFactory() . '<br>';
	}
}

$worker = new Client;