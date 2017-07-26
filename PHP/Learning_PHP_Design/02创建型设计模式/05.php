<?php 

/* 
	【 相识关系 】 
	设计模式中最基本、可能也是最常见的关系就是相识关系。相识关系是指一个
	参与者包含另一个参与者的引用。
 */

abstract class ISubject
{
	abstract public function login($un, $pw);
}

class RealSubject
{
	public function request()
	{
		echo 'request';
	}
}

class Proxy extends ISubject
{
	private $realSubject;

	protected function request()
	{
		$this->realSubject = new RealSubject();
		$this->realSubject->request();
	}

	public function login($un, $pw)
	{
		// 计算口今等
		if ($un === 'Proessional' && $pw === 'acp')
		{
			$this->request();
		}
		else
		{
			print_r('Cry \'Havoc,\' and let slip the dogs of war!');
		}
	}
}