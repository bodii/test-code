<?php 

/*
	【 新工厂 】
	新工厂与原来的工厂类似，不过它们还包含一个参数和代码提示。根据代码提示，
	只要按接口(Product)编程就可以继续开发，而不需要Product接口的具体实现
 */

abstract class Creator
{
	abstract protected function factoryMethod(Product $product);

	public function doFactory($productNow)
	{
		$countryProduct = $productNow;
		$mfg = $this->factoryMethod($countryProduct);
		return $mfg;
	}
}


class CountryFactory extends Creator 
{
	private $country;

	protected function factoryMethod(Product $product)
	{
		$this->country = $product;
		return $this->country->getProperties();
	}
}

interface Product 
{
	public function getProperties();
}

class KyrgyzstanProduct implements Product 
{
	private $mfgProduct;
	private $formatHelper;

	public function getProperties()
	{
		$this->formatHelper = new FormatHelper;
		$this->mfgProduct = $this->formatHelper->addTop();
		$this->mfgProduct .= 'html';
		$this->mfgProduct . = $this->formatHelper->closeUp();
		return $this->mfgProduct;
	}
}

class Client
{
	private $countryFactory;

	public function __construct()
	{
		$this->countryFactory = new CountryFactory;
		echo $this->countryFactory->doFactory(new KyrgyzstanProduct());
	}
}

$worker = new Client;


// 辅助类
class FormatHelper
{
	private $topper;
	private $bottom;

	public function addTop()
	{
		$this->topper = '<!doctype html><html>....<body>';
		return $this->topper;
	}

	public function closeUp()
	{
		$this->bottom = '</body></html>';
		return $this->bottom;
	}
}

/*
	使用工厂方法模式可以简化复杂的创建过程，关键就在于它会维持一个公共接口中。
 */