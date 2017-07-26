<?php 

/*
	关于包装器
	适配器和装饰器模式都有另外一个名字“包装器”(wrapper)。实际上，在一些定义
	中，就把包装器描述为适配器。

	包装器包装基本类型
*/

class PrimitiveWrap
{
	private $wrapeMe;

	public function __construct($prim)
	{
		$this->wrapeMe = $prim;
	}

	public function showWrap()
	{
		return $this->wrapeMe;
	}
}

$myPrim = 521;
$wrappedUp = new PrimitiveWrap($myPrim);
echo $wrappedUp->showWrap(), '<br>';

/* 
	PHP中的内置包装器
	file_get_contents()包装器。它将一个指定资源（如一个文件名，或一个文件名
	的url)绑定到一个流。
 */

class TextFileLoader
{
	private $textNow;

	public function __construct()
	{
		$this->textNow = file_get_contents('clery.txt');
		echo $this->textNow;
	}
}

$worker = new TextFileLoader;

/*
	设计模式包装器
	对象适配器模式中的适配器参与者“包装”被适配者。采用这种做法，
	就能创建一个与适配者兼容的接口。

	装饰器参与者可以“包装”一个组件对象，这样就能为这个已有的组件增加职责，
	而无须对它做任何修改。


	在设计模式中使用“包装器”是为了处理接口的不兼容，或者希望为组件增加功能，
	包装器就表示用来减少不兼容性的策略。
 */
