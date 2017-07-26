<?php


# 包装器
# 适配器和装饰器模式都有另外一个名字“包装器”（wrapper).

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
echo $wrappedUp->showWrap();


//php内置包装器 如 file_get_contents()。
