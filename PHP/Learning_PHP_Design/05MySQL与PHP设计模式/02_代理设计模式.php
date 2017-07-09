<?php 
	
/*
	保护代理 完成登录

	代理(Proxy)模式是一种结构型设计模式。“四人帮”(GoF)提出了4种代理模式：
	1.远程代理（Remote proxy）
		代理对象在一个地址空间，而实际对象在另一个地址空间中，此时代理就
		是远程代理

	2. 虚拟代理(Virtual proxy)
		虚拟代理可以缓存一个真实主题的有关信息，从而能延迟对这个真实主题的
		访问。有时在真实对象处理登录数据之前，高安全性登录会使用一个虚拟代
		理来完成登录。

	3. 保护代理（Protection proxy）
		保护代理只有在验证过请求之后，才会把请求发送到真实主题。这个真实主
		题就是请求的目标，如访问数据库信息。根据用户的登录信息，很多保护代
		理会提供不同级别的访问；并不是建立一个真实主题，真实主题可能是多个，
		而且是受限的。

	4. 智能引用(A smart reference)
		程序可以使用GoF所称的“祼指针”， 不过，如果“裸指针”不能满足程序的需
		要，代理就相当于一个智能引用（或智能指针），可以在引用对象时完成额
		外的动作。例如，可能首先由作为智能引用的代理参与者加载一个数据库的
		数据。


	代理模式有两个主要的参与者：
		一个代理主题(proxy subject)和一个真实主题(real subject)。
		客户通过Subject接口向Proxy提交请求，不过只有当请求首先通过Proxy之后
		才有可能访问RealSubject。

 */

/* 
	Proxy和RealSubject参与者都实现了Subject接口，不过Proxy的角色是一个
	“守门人”或封装模块。如果请求经由Proxy确认，Proxy再调用RealSubject为
	用户提供请求内容。
 */

class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();

		$drop = 'drop table if exists ' . $this->tableMaster;

		if ($this->hookup->query($drop) === true)
		{
			printf( "Old table %s has been dropped.<br>", $this->tableMaster);
		}

		$sql = 'ceate table ' .$this->tableMaster . '(uname varchar(15), pw varchar(120))';

		if ($this->hookup->query($sql) === true)
		{
			echo "Table {$this->tableMaster} has been created successfully.<br>";
		}

		$this->hookup->close();
	}
}

$worker = new CreateTable;


class HashRegister
{
	public function __construct()
	{
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();
		$username = $this->hookup->real_escape_string(trim($_POST['uname']));
		$pwNow = $this->hookup->real_escape_string(trim($_POST['pw']));

		$sql = "inster into {$this->tableMaster} (uname, pw) values('{$username}', md5('{$pwNow}'))";

		if ($this->hookup->query($sql))
		{
			echo "Registration completed:";
		}
		elseif (($result = $this->hookup->query($sql)) === false )
		{
			printf("Invalid query: %s <br> Whole query: %s <br>", 
				$this->hookup->error, $sql
				)
			exit();
		}
		$this->hookup->close();
	}
}

$worker = new HashRegister;