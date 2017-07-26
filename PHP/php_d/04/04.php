<?php


# 客户
# 这个模式最后一个参与者是隐含的：即客户。

class Client
{
	private $someGraphicObject;
	private $someTextObject;

	public function __construct()
	{
		$this->someGraphicObject = new GraphicFactory();
		echo $this->someGraphicObject->startFactory() . '<br>';
		$this->someTextOjbect = new TextFactory();
		echo $this->someTextObject->startFactory() . '<br>';
	}
}

$worker = new Client();
# Client对象并没有向产品直接做出请求，而是通过工厂来请求。重要的是，
# 客户并不实现产品特性，而留给产品实现来体现。


# 设计模式的真正价值并不是提高操作的速度，而是加快开发的速度。

