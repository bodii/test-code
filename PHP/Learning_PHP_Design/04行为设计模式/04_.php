<?php 

/*
	结合其他设计模式使用模板方法模式
	模板方法参与者
*/

abstract class TmAb
{
	protected $pix;
	protected $cap;

	public function templateMethod()
	{
		$this->addPix();
		$this->addCaption();
	}

	abstract protected function addPix();
	abstract protected function addCaption();
}


// 具体模板方法 (调用工厂方法)

class Tmfac extends TmAb
{
	protected function addPix()
	{
		$this->pix = new GraphicFactory;
		echo $this->pix->doFactory();
	}

	protected function addCaption()
	{
		$this->cap = new TextFactory;
		echo $this->cap->doFactory();
	}
}

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
		$product = new GraphicProduct;
		return $product->getProperties();
	}
}

class TextFactory extends Creator 
{
	protected function factoryMethod()
	{
		$product = new TextProduct;
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
		$this->mfgProduct = '<div style="color:#cc0000;font-size:12px;' 
		. 'font-family:Verdana, Geneva, sans-serif;">'
		. '<strong><em>Caption:</em></strong> Modigliani painted elongated'
		. 'painted elongated faces.</div>';
		return $this->mfgProduct;
	}
}


class Client 
{
	public function __construct()
	{
		$template = new Tmfac;
		$template->templateMethod();
	}
} 

$worker = new Client;