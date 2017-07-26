<?php

# 类图(工厂模式）
# 设计模式类图显示了参与者（类和接口）之间的关系和通信。

# 工厂方法 模式 通过一个工厂实例化对象，从而将实例化过程与请求者分离。



# 统一建模语言（Unified Modeling Language, UML)


# 相识关系
# 设计模式中最基本、可能也是最常见的关系就是相识关系。相识关系是指一个参与
# 者（抽象类和接口）包含另一个参与者的引用。


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
		// 计算口令等
		if($un === 'Professional' && $pw === 'acp')
		{
			$this->request();
		}
		else
		{
			echo "Cry 'Havoc,' and let slip the dogs of war!";
		}
	}
}


// 如上 Proxy 类包含RealSubject类的一个引用，也就与它建立了一个相识关系。
