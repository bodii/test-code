<?php 

/*
	【 原型设计模式 】
	原型设计模式(Prototype Design Pattern),它使用了一种克隆技术来复制实例化
	的对象。新对象是通过复制原型实例创建的。在这里，实例(instance)是指实例化
	的具体类。原型设计模式的目的是通过使用克隆以减少实例化对象的开销。与其从
	一个类实例化新对象，完全可以使用一个已有实例的克隆。

	？什么时候用到原型设模式
	原型模式还可以用来创建一种组织结构，可以根据实际的组织来创建和填充其中的
	位置(数目有限)。例如，可以使用组合创建一个绘图对象，然后克隆为不同版本的
	绘图对象，这就使用了原型模式。最后再来看一个使用原型模式的例子，在游戏开
	发中，可以通过克隆一个原型士兵，来增加军队的人数和兵种。

	PHP中使用原型设计模式的关键是要了解如何使用内置函数__clone()。
 */

abstract class CloneMe
{
	public $name;
	public $picture;
	abstract function __clone();
}

class Person extends CloneMe 
{
	public function __construct()
	{
		$this->picture = 'cloneMan.png';
		$this->name = 'Original';
	}

	public function display()
	{
		echo "<img src='{$this->picture}'>";
		echo "<br>{$this->name}<p />";
	}

	function __clone(){}
}

$worker = new Person();
$worker->display();

$slacker = clone $worker;
$slacker->name = 'Cloned';
$slacker->display();

/*
	对于所克隆的实例，clone关键字会为该实例的类实例化另一个实例(副本)。但不
	能直接调用对象的__clone()方法。
 */