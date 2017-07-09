<?php 

/*
	【 封装 】
	封装： 信息隐藏

	封装 就是划分一个抽象的诸多元素的过程，这些元素构成该抽象的结构和行为；
	封装的作用就是将抽象的契约接口与其实现分离。

	一旦把一个复杂的大问题模块化为多个可解决的子问题，就可以利用封装来得到
	这些较小的抽象，并对它们完成划分。
 */

/* 
	获取方法和设置方法
	为了保持封装，同时提供可访问性，OOP设计建议使用获取方法(getters) 和设置
	方法(setters).
 */
 
 namespace GetSet;

 class GetSet
 {
 	private $dataWarehouse;

 	public function __construct()
 	{
 		$this->setter(200);
 		$got = $this->getter();
 		echo $got;
 	}

 	private function getter()
 	{
 		return $this->dataWarehouse;
 	}

 	private function setter($value)
 	{
 		$this->dataWarehouse = $value;
 	}
 }

 $worker = new GetSet;

 /*
 	不要直接请求完成一个工作所需的信息，而应当请求拥有这个信息的对象
 	为你完成工作。
  */
 

 /*
 	【 多态 】 
 	多态的真正价值在于，可以调用有相同接口的对象来完成不同的工作。
 	一个名字，多个实现。

 	设计模式中的内建多态性
 	
  */