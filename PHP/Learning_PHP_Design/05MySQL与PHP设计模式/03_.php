<?php 
	
/*
	代理的工作
	由于Proxy和RealSubject类有一个共同的接口，它们必须分别实现ISubject接口:
*/

interface ISubject
{
	public function request();
}

class Proxy implements ISubject
{
	private $tableMaster;
	private $hookup;
	private $logGood;
	private $realSubject;

	public function login($uNow, $pNow)
	{
		// 由客户过滤；对口令掩码
		$uname = $uNow;
		$pw = md5($pNow);
		$this->logGood = false;

		// 选择表和连接
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();

		// 创建mysql语句
		$sql = "select pw from {$this->tableMaster} where uname='{$uname}'";

		if ($result = $this->hookup->query($sql))
		{
			$row = $result->fetch_array(MYSQL_ASSOC);
			if ($row['pw'] == $pw)
			{
				$this->logGood = true;
			}
			$result->close();
		}
		elseif (($reuslt = $this->hookup->query($sql)) === false)
		{
			printf('Failed: %s <br>', $this->hookup->error);
			exit;
		}
		$this->hookup->close();

		if ($this->logGood)
		{
			$this->request();
		}
		else
		{
			echo 'Username and/or Password not on record.';
		}
	}

	public function request()
	{
		$this->realSubject = new RealSubject;
		$this->realSubject->request();
	}
}

/* 在代理模式中，真实主题(请求的目标)的代理是Proxy类，所以它的任务是掩护真实
	主题，而不是重复真实主题的request()方法。
 */

/*
	真实主题
 */

class RealSubject implements ISubject 
{
	public function request()
	{
		$practice = <<<REQUEST 
		<!Doctype html>
		<html>
		<body>
		<body>
		</html>
REQUEST;

	}
}