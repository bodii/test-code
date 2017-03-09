<?php


# 克隆不会启动构造函数

# 克隆不会启动构造函数中的动作。克隆可以使用构造函数生成的默认值，但如果
# 构造函数生成一个动作，如一旦实例化就打印一条消息，克隆并不会显示这个消
# 息。

include_once '../display_errors.php';


class HelloClone
{
	private $builtInConstructor;

	public function __construct()
	{
		echo "Hello, clone!<br>";
		$this->builtInConstructor = "Constructor creation<br>";

	}

	public function doWork()
	{
		echo $this->builtInConstructor;
		echo "I'm doing the work!<p />";
	}
}

$original = new HelloClone();
$original->doWork();

// 克隆
$cloneIt = clone $original;
$cloneIt->doWork();


# 构造函数不要做具体工作
# 如果一个类实例化时要完成大量初始化，结果往往不灵活，而且这是过度耦合的设计。
# 同样地，如果一个构造函数输出某些结果，客户除了能看到构造函数发送的结果外，
# 没有任何其他选择，毕竟客户可能并不想要这些结果，或者至少并不想在给定时间
# 得到这些结果。
