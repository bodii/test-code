<?php

# 聚合关系
# 聚合关系和相识关系类似，不过关系更强。聚合关系表示一个聚合对象与它的所有
# 者有相同的生命期。像人的心脏和肺。只要心脏供血，肺就能够在人体内输送氧化
# 的血液。它们都是可以独立运转的生理器官，但是如果一个停止工作，另一个也停
# 止，无法继续正常工作。

class Context
{
	private $strategy;

	public function __construct(IStrategy $strategy)
	{
		$this->strategy = $strategy;
	}

	public function algorithm($elements)
	{
		$this->strategy->algorithm($elements);
	}
}

# 代码提示引用指向IStrategy接口，不必在Context参与者（类）中实例化一个
# IStrategy实现。

interface IStrategy
{
	public function algorithm($elements);
}

