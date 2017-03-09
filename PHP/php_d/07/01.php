<?php

# 行为模式设计
/*
	11个行为型模式：
	
	职责链模式
	命令模式
	解释器模式（类设计模式）
	迭代器模式
	中介者模式
	备忘录模式
	观察者模式
	状态模式
	策略模式
	模板方法模式（类设计模式）
	访问者模式

*/

# 抽象类中的模板方法操作， 可以看做是“基本操作的组织者”。

// 简单的例子：对图像和图题使用模板方法模式

abstract class AbstractClass
{
	protected $pix;
	protected $cap;

	public function templateMethod($pixNow, $capNow)
	{
		$this->pix = $pixNow;
		$this->cap = $capNow;
		$this->addPix($this->pix);
		$this->addCaption($this->cap);
	}

	abstract protected function addPix($pix);
	abstract protected function addCaption($cap);
}

// 具体实现
class ConcreteClass extends AbstractClass
{
	protected function addPix($pix)
	{
		$this->pix = $pix;
		$this->pix  = "pix/" . $this->pix;
		$formatter = "<img src=$this->pix><br>";
		echo $formatter;
	}

	protected function addCaption($cap)
	{
		$this->cap = $cap;
		echo "<em>Caption:</em>" . $this->cap . '<br>';
	}
}


class Client
{
	function __construct()
	{
		$caption = "Modigliani painted elongated faces.";
		$mo = new ConcreteClass();
		$mo->templateMethod('moding.png', $caption);
	}
}

$worker = new Client();
