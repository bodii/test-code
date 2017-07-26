<?php 

/*
	【 接口 】
	类型提示：类似数据类型

	类型提示基本格式：

		function doWork(TypeHint $someVar)...
 */

namespace IProduct;

interface IProduct 
{
	function apples();
	function oranges();
}


namespace FruitStore;

use IProduct\IProduct;

class FruitStore implements IProduct
{
	public function apples()
	{
		return 'FruitStore sez -- We have apples.<br>';
	}

	public function oranges()
	{
		return 'FruitStore sez -- We have no citrus fruit.<br>';
	}
}


namespace CitrusStore;

use IProduct\IProduct;

class CitrusStore implements IProduct 
{
	public function apples() 
	{
		return 'CitrusStore sez -- We do not sell apples.<br>';
	}

	public function oranges()
	{
		return 'CitrusStore sez -- We have citrus fruit.<br>';
	}
}


namespace UseProducts;

use FruitStore\FruitStore;
use CitrusStore\CitrusStore;
use IProduct\IProduct;

class UseProducts
{
	public function __construct()
	{
		$appleSauce = new FruitStore;
		$orangeJuice = new CitrusStore;
		$this->doInterface($appleSauce);
		$this->doInterface($orangeJuice);
	}

	public function doInterface(IProduct $product)
	{
		/* 类型提示(type hint) IProduct能够识别实现IProduct接口的两个类。
		换句话说，并不是把它们分别识别为一个FruitStore实例和另一个Citrus-
		Store实例，而会识别它们共同的接口IProduct.
		 */
		echo $product->apples();
		echo $product->oranges();
	}
}

$worker = new UseProducts();

/**
 * 强制数据类型可以确保倘若给定方法中使用了代码提示，那么其中使用的
 * 对象(类)必然有给定的接口。另外，如果把一个接口(可以是一个抽象类或
 * 接口)作为代码提示，绑定会更宽松；它会绑定到接口而不是绑定到一个特
 * 定的实现。随着程序变得越来越大，只要遵循接口，就可以做任何改变而
 * 不会对程序造成破坏。不仅如此，所做的修改也不会与具体实现纠缠不清。
 *
 * 不能使用标题类型(如string或int)作为代码提示，不过可以使用数组、接
 * 口和类作为代码提示。所以尽管没有另外一些语言那么灵活，但PHP中可以
 * 通过类型提示实现类型，这在OOP和设计模式编程中起着重要作用。
 */