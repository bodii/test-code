<?php 

/*
	开发人员约会服务
	为了说明如何实现一个包装多个组件的装饰器，下面给出一个例子，它为软件开发
	人员建立了一个约会服务。这里有两个组件，分别是Male和Female，可以分别为这
	两个组件装饰不同的约会关注点。可以用相同或不同的具体装饰采用任意组合来装
	饰这些组件。
*/

/*
	组件接口
	组件接口包括3个属性和5个方法。$date属性用来标识这是一个“约会”,而不是普通的
	日/月/年形式的日期对象。$ageGroup是该组件所属的组，$feature是由某个具体装饰
	提供的特性：
 */


abstract class IComponent
{
	protected $date;
	protected $ageGroup;
	protected $feature;

	abstract public function setAge($ageNow);
	abstract public function getAge();
	abstract public function getFeature();
	abstract public function setFeature($fea);
	/* 所有属性都是受保护的，所有方法都是抽象的 */
}


/* 具体组件 */

class Male extends IComponent
{
	public function __construct()
	{
		$this->date = 'Male';
		$this->setFeature('<br>Dude programmer features:');
	}

	public function getAge()
	{
		return $this->ageGroup;
	}

	public function setAge($ageNow)
	{
		$this->ageGroup = $ageNow;
	}

	public function getFeature()
	{
		return $this->feature;
	}

	public function setFeature($fea)
	{
		$this->feature = $fea;
	}
}


class Female extends IComponent
{
	public function __construct(){
		$this->date = 'Female';
		$this->setFeature('<br>Grrrl programmer features: ');
	}

	public function getAge()
	{
		return $this->ageGroup;
	}

	public function setAge($ageNow)
	{
		$this->ageGroup = $ageNow;
	}

	public function getFeature()
	{
		return $this->feature;
	}

	public function setFeature($fea)
	{
		$this->feature = $fea;
	}
} 

/*
	包含组件方法的装饰器
 */

abstract class Decorator extends IComponent
{
	public function setAge($ageNow)
	{
		$this->ageGroup = $this->ageGroup;
	}

	public function getAge()
	{
		return $this->ageGroup;
	}
}
/* 如果想增加一个else语句来捕获错误，这很容易做到。不过，目前采用了一种
“安静失败”的做法，即不会传递对象。 */

/*
	具体装饰器
 */

class ProgramLang extends Decorator
{
	private $languageNow;
	private $language = [
		'php' => 'PHP',
		'cs' => 'C#',
		'js' => 'JavaScript',
		'as3' => 'ActionScript 3.0',
	];

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	public function setFeature($lan)
	{
		$this->languageNow = $this->language[$lan];
	}

	public function getFeature()
	{
		$output = $this->date->getFeature();
		$fmat = '<br>&nbsp;&nbsp;';
		$output .= "$fmat Preferred programming language: ";
		$output .= $this->languageNow;
		return $output;
	}
}

class Hardware extends Decorator
{
	private $hardwareNow;
	private $box = [
		'mac' => 'Macintosh',
		'dell' => 'Dell',
		'hp' => 'Hewlett-Packard',
		'lin' => 'Linux',
	];

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	public function setFeature($hdw)
	{
		$this->hardwareNow = $this->box[$hdw];
	}

	public function getFeature()
	{
		$output = $this->date->getFeature();
		$fmat = '<br>&nbsp;&nbsp;';
		$output .= "$fmat Current Hardware: ";
		$output .= $this->hardwareNow;
		return $output;
	}
}

class Food extends Decorator
{
	private $chowNow;
	private $snacks = [
		'piz' => 'Pizza',
		'burg' => 'Burgers',
		 'nach' => 'Nachos',
		 'veg' => 'Veggies',
	];

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	public function setFeature($yum)
	{
		$this->chowNow = $this->snacks[$yum];
	}

	public function getFeature()
	{
		$output = $this->date->getFeature();
		$fmat = '<br>&nbsp;&nbsp;';
		$output .= "$fmat Favorite food: ";
		$output .= $this->chowNow ."<br>";
		return $output;
	}

}

class Client
{
	// $hotDate是组件实例
	private $hotDate;

	public function __construct()
	{
		$this->hotDate = new Female;
		$this->hotDate->setAge('Age Group 4');
		echo $this->hotDate->getAge();

		$this->hotDate = $this->wrapComponent($this->hotDate);
		echo $this->hotDate->getFeature();
	}

	private function wrapComponent(IComponent $component)
	{
		$component = new ProgramLang($component);
		$component->setFeature('php');
		$component = new Hardware($component);
		$component->setFeature('lin');
		$component = new Food($component);
		$component->setFeature('veg');
		return $component;
	}
}

$worker = new Client;