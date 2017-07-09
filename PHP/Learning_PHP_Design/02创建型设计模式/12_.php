<?php 

/*
	抽象类接口和具体实现
*/

abstract class IPrototype
{
	public $eyeColor;
	public $wingBeat;
	public $unitEyes;

	abstract public function __clone();
}

class MaleProto extends IPrototype
{
	const gender='MALE';
	public $mated;

	public function __construct()
	{
		$this->eyeColor = 'red';
		$this->wingBeat = '220';
		$this->unitEyes = '760';
	}

	public function __clone(){}
}

class FemaleProto extends IPrototype
{
	const gender = 'FEMALE';
	public $fecundity;

	public function __construct()
	{
		$this->eyeColor = 'red';
		$this->wingBeat = '220';
		$this->unitEyes = '760';
	}

	function __clone(){}
}

/*function __autoload($class_name)
{
	include $class_name . '.php';
}*/


class Client
{
	// 用于直接实例化
	private $fly1;
	private $fly2;

	// 用于克隆
	private $c1Fly;
	private $c2Fly;
	private $updatedCloneFly;

	public function __construct()
	{
		// 实例化
		$this->fly1 = new MaleProto;
		$this->fly2 = new FemaleProto;

		// 克隆
		$this->c1Fly = clone $this->fly1;

		$this->c2Fly = clone $this->fly2;

		$this->updatedCloneFly = clone $this->fly2;

		// 更新克隆
		$this->c1Fly->mated = 'true';
		$this->c2Fly->fecundity = '186';
		$this->updatedCloneFly->eyeColor ='purple';
		$this->updatedCloneFly->wingBeat = '220';
		$this->updatedCloneFly->unitEyes = '750';
		$this->updatedCloneFly->fecundity = '92';

		// 通过类型提示方法发送
		$this->showFly($this->c1Fly);
		$this->showFly($this->c2Fly);
		$this->showFly($this->updatedCloneFly);
	}

	private function showFly(IPrototype $fly)
	{
		echo 'Eye color: '. $fly->eyeColor . '<br>';
		echo 'Wing Beats/second: ' . $fly->wingBeat . '<br>';
		echo 'Eye units: ' . $fly->unitEyes . '<br>';

		$genderNow = $fly::gender;
		echo 'Gender: '. $genderNow . '<br>';
		if ($genderNow == 'FEMALE') {
			echo 'Number of eggs: ' . $fly->fecundity . '<br>';
		} 
		else 
		{
			echo 'Mated: ' . $fly->mated . '<p/>';
		}
	}
}


$worker = new Client;


/*
	为原型模式增加OOP

$this->updatedCloneFly->eyeColor = 'purple';
*/