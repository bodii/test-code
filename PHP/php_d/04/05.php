<?php

/*
新工厂
新工厂与原来的工厂类似，不过它们还包含一个参数和代码提示。
*/

include_once '../display_errors.php';

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
		$this->formatHelper = new FormatHelper();
		$this->mfgProduct = $this->formatHelper->addTop();
		$this->mfgProduct .= <<<KYRGYZSTAN
<img src='Countries/Kyrgyzstan.png' class='pixRight' width='600' height='304'>
<header>Kyrgyzstan</header>
<p>A ....</p>
KYRGYZSTAN;
		$this->mfgProduct .= $this->formatHelper->closeUp();
		return $this->mfgProduct;
	}
}

abstract class Creator
{
	protected abstract function factoryMethod(Product $product);

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

	protected function factoryMethod(Product $product){
		$this->country = $product;
		return($this->country->getProperties());
	}
}


# 新产品


class Client
{
	private $countryFactory;

	public function __construct()
	{
		$this->countryFactory = new CountryFactory();
		echo $this->countryFactory->doFactory(new KyrgyzstanProduct());
	}
}

$worker = new Client();

# 辅助类
class FormatHelper
{
	private $topper;
	private $bottom;

	public function addTop()
	{
		$this->topper = '<!DOCTYPE html><html><head>
			<meat charset="utf-8">
			<title>Map Factory</title>
			</head>
			<body>';
		return $this->topper;
	}

	public function closeUp()
	{
		$this->bottom = '</body></html>';
		return $this->bottom;
	}
}
