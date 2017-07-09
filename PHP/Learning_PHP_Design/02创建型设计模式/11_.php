<?php 

/*
	克隆不会启动构造函数

	关于克隆过程，有一点要注意：克隆不会启动构造函数中的动作。克隆可以使用
	构造函数生成的默认赋值，但是如果构造函数生成一个动作，如一旦实例化就打
	印一条消息，克隆并不会显示这个消息(比如：‘Hello,clone!’)。

	*/

class HelloClone
{
	private $builtInConstructor;

	public function __construct()
	{
		echo 'Hello, clone!<br>';
		$this->builtInConstructor = 'Constructor creation<br>';
	}

	public function doWork()
	{
		echo $this->builtInConstructor;
		echo 'I\'m doing the work!<p>';
	}
}


$original = new HelloClone;
$original->doWork();

/*Hello, clone!
Constructor creation
I'm doing the work!*/


// 克隆不启动构造函数
$cloneIt = clone $original;
$cloneIt->doWork();

/*Constructor creation
I'm doing the work!*/


/*
	对于原型设计模式(以及其他使用克隆的方法)来说，这意味着不能依赖于构造
	函数提供重要的输出或返回结果。不过，这可能不是一件坏事。实际上，这是
	一个很好的编程实践。
 */

/*
	构造函数不要做具体工作
	如果一个类实例化时要完成大量初始化，结果往往不灵活，而且这是过度耦合
	的设计。同样地，如果一构造函数输出某些结果，客户除了能看到构造函数发
	送的结果外，没有任何其他选择，毕竟客户可能并不想要这些结果，或者至少
	并不想在给定时间得到这些结果。

	客户的构造函数可能与模式中其他参与者完全不同，因为它要向参与者做出请
	求。

	对于这个概念，即构造函数不应完成具体的工作，一种做法是忽略模式类中的
	构造函数，除非你有充分的理由包含这些构造函数；另外一种做法是，允许在
	需要时调用操作和，让客户负责实例化和克隆的有关事务。所以，尽管我们可
	能发现使用__clone()函数时存在限制，但这些限制可能更有助完成更好的OOP
	程序。
 */