<?php

// 重新设计登录代理

include_once '01.php';
interface ISubject
{
	public function request();
}

class Proxy implements ISubject
{
	private $tableMater;
	private $hookup;
	private $logGood;
	private $realSubject;

	public function login($uNow, $pNow)
	{
		// 由客户过滤; 对口令掩码
		$uname = $uNow;
		$pw = md5($pNow);
		$this->logGood = false;

		// 选择表和连接
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();

		// 创建MySQL语句
		$sql = "select pw from $this->tableMatster where uname='$uname'";

		if($result= $this->hookup->query($sql))
		{
			$row = $result->fetch_array(MYSQLI_ASSOC);
			if($row['pw'] == $pw)
			{
				$this->logGood = true;
			}
			$result->close();
		}
		elseif ( ($result = $this->hookup->query($sql)) === false )
		{
			printf("Failed: %s <br>", $this->hookup->error);
			exit();
		}
		$this->hookup->close();
		if($this->logGood)
		{
			$this->request();
		}
		else
		{
			echo "Username and/or Password not on record.";
		}
	}

	public function request()
	{
		$this->realSubject = new RealSubject();
		$this->realSubject->request();
	}
}


class RealSubject implements ISubject
{
	public function request()
	{
		$practice = <<<REQUEST
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<link type="text/css" href="proxy.css" >
	</head>
	<body>
	<header>PHP Tip Sheet:<br>
	<span class="subhead"> For OOP Developers</span>
	</header>
	<ol>
		<li>Program to the interface and not the implementation.</li>
		<li>Encapsulate your objects.</li>
		<li>Favor composition over class inheritance.</li>
		<li>A class Should only have a single responsibility.</li>
	</ol>
	</body>
</html>
REQUEST;
		echo $practice;
	}
}


class Client
{
	private $proxy;
	private $un;
	private $pw;

	public function __construct()
	{
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();
		$this->un = $this->hookup->real_escape_string(trim($_POST['uname']));
		$this->pw = $this->hookup->real_escape_string(trim($_POST['pw']));
		$this->getIface($this->proxy = new Proxy());
	}

	private function getIface(ISubject $proxy)
	{
		$proxy->login($this->un, $this->pw);
	}

}

$worker = new Client();

