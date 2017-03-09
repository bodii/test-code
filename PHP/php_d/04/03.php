<?php


# 继承和实现关系
 
# 多重关系

# 创建关系
# 在一个设计模式中，一个对象创建另一个对象的实例。

# 对象图
# 类图会显示类，与之不同，对象图只显示实例。


# 创建型设计模式
# 创建型设计模式强调的是实例化过程。设计这些模式的目的是隐藏实例的创建过程，
# 并封装对象使用的知识。
# 5个创建型设计模式包括：
# 抽象工厂 （Abstract Factory）
# 生成器 （Builder）
# 工厂方法 （Factory Method）
# 原型 （Propertype）
# 单例	（Singleton）

# 工厂方法模式是这5个设计模式中唯一的一种类设计模式，尽管相当简单，却是一种
# 非常有效的模式。原型模式属于对象类模式，可以使用PHP __clone()方法实现。首
# 先基于原型实例化（创建）一个对象，然后由这个实例化对象进一步克隆其他对象。


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
		$this->mfgProduct = 'This is a graphic.<3';
		return $this->mfgProduct;
	}
}

class TextFactory extends Creator
{
	protected function factoryMethod()
	{
		$product = new TextProduct();
		return($product->getProperties());
	}
}

class GraphicFactory extends Creator
{

	protected function factoryMethod()
	{
		$product = new GraphicProduct();
		return($product->getProperties());
	}
}

# 这两个工厂实现是类似的，只不过一个创建TextProduct实例，而另一个创建
# GraphicProduct实例。


