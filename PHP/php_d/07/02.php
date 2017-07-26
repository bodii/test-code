<?php

# 好莱坞原则
# “反向控制结构”概念另外也称为“好莱坞原则”。这个原则是指父类调用子类的
# 操作，而子类不调用父类的操作。
#
# 与好莱坞原则关联最紧密的模式就是模板方法模式，因为它在父类中实现。除了
# templateMethod(),父类中的其他操作（方法）都是抽象和保护方法。


# 在过程编程中，关键问题是控制流，在OOP中，关键问题则是对象职责。


// 工厂方法参与者
abstract class Creator
{
	protected abstract function factoryMethod();

	public function doFactory()
	{
		$mfg = $this->factoryMethod();
		return $mfg;
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

class TextFactory extends Creator
{
	protected function factoryMethod()
	{
		$product = new TextProduct();
		return $product->getProperties();
	}
}

interface Product
{
	public function getProperties();
}

class GraphicProduct implements Product
{
	private $mfgProduct;

	public function getProperties()
	{
		$this->mfgProduct = '<img src="pix/modig.png">';
		return $this->mfgProduct;
	}
}

class TextProduct implements Product
{
	private $mfgProduct;

	public function getProperties()
	{
		$this->mfgProduct = '<div style="color:#cc0000; font-size:12px;
							font-family:Verdana, Geneva, sans-serif">
								<strong><em>Caption:</em></strong>Modigliani
								 painted elongated faces.</div>';
		return $this->mfgProduct;
	}
}


