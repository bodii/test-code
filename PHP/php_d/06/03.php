<?php


# 设计模式包装器
# 在设计模式中使用“包装器”是为了处理接口的不兼容，或者希望为组件增加功能，
# 包装器就表示用来减少不兼容性的策略。

include_once '../display_errors.php';
// 组件接口
abstract class IComponent
{
	protected $date;
	protected $ageGroup;
	protected $feature;

	abstract public function setAge($ageNow);
	abstract public function getAge();
	abstract public function getFeature();
	abstract public function setFeature($fea);

}

// 具体组件
class Male extends IComponent
{
	public function __construct()
	{
		$this->date = 'Male';
		$this->setFeature('<br>Dude programmer features: ');
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

// Female 具体组件
class Female extends IComponent
{
	public function __construct()
	{
		$this->date = 'Female';
		$this->setFeature('<br>Grrrl Programmer features: ');
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


// 装饰器参与者
abstract class Decorator extends IComponent
{
	public function setAge($ageNow)
	{
		$this->ageGroup = $ageNow;
	}

	public function getAge()
	{
		return $this->ageGroup;
	}
}

// 具体装饰器
class ProgramLang extends Decorator
{
	private $languageNow;

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	private $language = array('php'=>'php',
							'cs' => 'C#',
							'js' => 'JavaScript',
							'as3' => 'ActionScript 3.0');

	public function setFeature($lan)
	{
		$this->languageNow = $this->language[$lan];
	}

	public function getFeature()
	{
		$output = $this->date->getFeature();
		$fmat = "<br>&nbsp;&nbsp;";
		$output .= "$fmat Preferred programming language: ";
		$output .= $this->languageNow;
		return $output;
	}
}

// 具体装饰器
class Hardware extends Decorator
{
	private $hardwareNow;

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	private $box = array('mac' => 'Macintosh',
						'dell' => 'Dell',
						'hp' => 'Hewlett-packard',
						'lin' => 'linux');

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

// 具体装饰器
class Food extends Decorator
{
	private $chowNow;

	public function __construct(IComponent $dateNow)
	{
		$this->date = $dateNow;
	}

	private $snacks = array('piz' => 'Pizza',
							'burg' => 'Burgers',
							'nach' => 'nachos',
							'veg' => 'Veggies'
						);

	public function setFeature($yum)
	{
		$this->chowNow = $this->snacks[$yum];
	}

	public function getFeature()
	{
		$output = $this->date->getFeature();
		$fmat = '<br>&nbsp;&nbsp;';
		$output .= "$fmat Favorite food: ";
		$output .= $this->chowNow . "<br>";
		return $output;
	}

}


class Client
{	
	// $hotDate是组件实例
	private $hotDate;

	public function __construct()
	{
		$this->hotDate = new Female();
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

$wroker = new Client();
