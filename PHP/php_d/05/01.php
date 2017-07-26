<?php

# 原型设计模式
#
# 原型设计模式使用了一种克隆技术来复制实例化的对象。

# 如果一个项目要求你创建某个原型对象的多个实例，就可以使用原型模式。


# 克隆函数
# __clone()

include_once '../display_errors.php';

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
		$this->picture = 'CloneMe.png';
		$this->name  = 'Original';
	}

	public function display()
	{
		echo "<img src='$this->picture'>";
		echo "<br><p>$this->name<p />";
	}

	public function __clone() {}
}

$worker = new Person();
$worker->display();

$slacker = clone $worker;
$slacker->name ="Cloned";
$slacker->display();
