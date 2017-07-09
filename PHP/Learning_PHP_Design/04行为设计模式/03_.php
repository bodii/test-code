<?php 

/*
	抽象类
*/

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


/*
	具体类
 */

class ConcreteClass extends AbstractClass
{
	protected function addPix($pix)
	{
		$this->pix = $pix;
		$this->pix = 'pix/' . $this->pix;
		$formatter = '<img src="' . $this->pix . '"><br>';
		echo $formatter;
	}

	protected function addCaption($cap)
	{
		$this->cap = $cap;
		echo "<em>Caption:</em>" . $this->cap . "<br>";
	}
}

class Client
{
	public function __construct()
	{
		$caption = 'Modigliani painted elongated faces.';
		$mo = new ConcreteClass;
		$mo->templateMethod('modig.png', $caption);
	}
}

$worker = new Client;

/*
	好莱坞原则
	“反向控制结构” 概念另外也称为“好莱坞原则”（the Hollywood Principle）。
	这个原则是指父类调用子类的操作，而子类不调用父类的操作。（就像试镜之
	后，导演告诉年轻的演员，“如果你拿到这个角色，我们会通知你。不要打电
	话来询问；我们会打电话给你的”。正因如此，这种反向控制结构被称为好莱坞
	原则。）

	与好莱坞原则关联最紧密的模式就是模板方法模式，因为它在父类中实现。除了
	templateMethod(),父类中的其他操作（方法）都是抽象和保护方法。所以，尽管
	客户实例化一个具体类，但它调用了父类中实现的方法。
 */

/* 在过程编程中，关键问题是控制流(flow of control)。 在OOP中，关键问题则是
	对象职责(object responsibility)。由于过程编程的重点类于控制流，一些解释
	会使用“控制反向”(inversion of control)来解释好莱坞原则。控制反向从过程
	编程的角度来讲很有意义，不过在OOP中，大多数控制流都会通过对象职责和协作
	抽取出来。也就是说，并不是考虑控制流，而应考虑哪些对象将处理某些职责，另
	外对象如何协作来完成任务。
 */

/*
	模板方法定义了框架，子类可以扩展或重新实现算法的可变部分，不过它们不能改
	变模板方法的控制流。【从子类发出“调用”是重新实现父类的方法】，这是好莱坞
	原则不允许的一种“调用”。只有父类可以做出“调用”来建立或改变框架（操作的执
	行顺序）。

	要理解好莱坞原则，更好的办法可能是按幼儿园里老师和学生之间的关系来考虑，
	也就是幼儿园原则(Kindergarten Principle)。老师建立一些项目，让孩子们按
	某种顺序完成。

	由父类建立顺序，子类按自已特有的方式完成这些操作。
	如果讨论好莱坞原则时谈到控制反向，要记住这个“反向”只在过程编程的控制流
	上下文中才有意义。

 */