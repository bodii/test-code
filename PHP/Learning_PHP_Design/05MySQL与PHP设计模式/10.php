<?php 
	
/*
	构建和加载响应表
*/

interface IConnectInfo
{
	const HOST = 'localhost';
	const UNAME = 'phpWorker';
	const PW = 'easyWay';
	const DBNAME = 'dpPatt';

	public function doConnect();
}

class UniversalConnect implements IConnectInfo
{
	private static $server = IConnectInfo::HOST;
	private static $currentDB = IConnectInfo::DBNAME;
	private static $user = IConnectInfo::UNAME;
	private static $pass = IConnectInfo::PW;
	private static $hookup;

	public function doConnect()
	{
		self::$hookup = mysqli_connect(self::$server, self::$user, self::$pass, self::$currentDB);

		if (self::$hookup) 
		{
			echo 'Successful connection to MySQL:<br>';
		}
		elseif (mysqli_connect_error(self::$hookup))
		{
			die('Here is why it failed: ' . mysqli_connect_error()); 
		}
		return self::$hookup;
	}
}

class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function  __construct()
	{
		$this->tableMaster = 'helpdesk';
		$this->hookup = UniversalConnect::doConnect();

		$dorp = "drop table if exists {$this->tableMaster}";

		if($this->hookup->query($drop) === true)
		{
			printf("Old table %s has been dropped.<br>", $this->tableMaster);
		}

		$sql = "create table $this->tableMaster (id int not null auto_incremant primary key, chain varchar(3), response text)engine=Innodb default charset=utf8";

		if ($this->hookup->query($sql) === true)
		{
			printf("Table {$this->tableMaster} has been created successfully.<br>");
		}

		$this->hookup->close();
	}
}

$worker = new CreateTable;


class InsertData
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'helpdesk';
		$this->hookup = UniversalConnect::doConnect();

		// 从HTML表单获取数据
		$chain = $this->hookup->real_escape_string($_POST['chain']);
		$response = $this->hookup->real_escape_string($_POST['response']);

		// 创建MySQL语句
		$sql = "insert into {$this->tableMaster} (chain, response) values"
		."('$chain', '$response')";

		if ($this->hookup->query($sql))
		{
			printf("Chain query: %s <br> Response %s<br> have been"
				." insterted into %s.", $chain, $response, $this->tableMaster);
		} 
		elseif ( ($result = $this->hookup->query($sql)) === false)
		{
			printf("Invalid query: %s <br> Whole query: %s <br?", 
				$this->hookup->error, $sql);
			exit();
		}
		$this->hookup->close();
	}
}

$worker = new InsertData;


class UpdaeData
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		// 制定表和链接
		$this->tableMaster = 'helpdesk';
		$this->hookup = UniversalConnect::doConnect();

		// 从HTML表单获取的数据
		$chain = $this->hookup->real_escape_string($_POST['chain']);
		$response = $this->hookup->real_escape_string($_POST['response']);

		// 创建MySQL语句
		$sql = "update {$this->tableMaster} set response='$response' where"
		. " chain='$chain'";

		if ($this->hookup->query($sql))
		{
			printf("Chain query: %s <br> Response %s <br> have been changed " 
				. " and set into %s.", $chain, $response, $this->tableMaster);
		}
		elseif ( ($result = $this->hookup->query($sql)) === false)
		{
			printf("Invalid query: %s<br> Whole query: %s <br>", 
				$this->hookup->error(), $sql);
			exit();
		}

		$this->hookup->close();
	}
}

$worker = new UpdateData;

class Client
{
	private $queryNow;

	public function __construct()
	{
		if (isset($_POST['sendNow']))
		{
			$this->queryNow = $_POST['help'];
		}

		$q1 = new Q1;
		$q2 = new Q2;
		$q3 = new Q3;
		$q4 = new Q4;
		$q5 = new Q5;

		// 设置后继
		$q1->setSuccessor($q2);
		$q2->setSuccessor($q3);
		$q3->setSuccessor($q4);
		$q4->setSuccessor($q5);

		// 生成和处理加载请求
		$loadup = new Requset($this->queryNow);
		// 设置链首
		$q1->handleRequest($loadup);
	}
}

$makeRequest = new Client;

class Request
{
	private $value;

	public function __construct($service)
	{
		$this->value = $service;
	}

	public function getService()
	{
		return $this->value;
	}
}

/* 处理器接和具体处理器 */

abstract class Handler
{
	protected $successor;
	protected $hookup;
	protected $tableMaster;
	protected $sql;
	protected $handle;

	// 处理请求方法非常简单。handleRequest()方法传递请求(将请求作为参数)
	// 由Client首先发起请求，启动这个链--就像是点燃导火索。
	// 如果$handle变量与Client通过Request辅助类传递的$request匹配，就会
	// 由这个具体处理器来处理这个查询。否则，它会将请求继续传递给链中的后继
	abstract public function handleRequest($request);
	abstract public function setSuccessor($nextService);
}

class Q1 extends handler 
{
	public function setSuccessor($nextService)
	{
		$this->successor = $nextService;
	}

	public function handleRequest ($request)
	{
		$this->handle = 'q1';
		if ($reqeust->getService() == $this->handle)
		{
			$this->tableMaster = 'helpdesk';
			$this->hookup = UniversalConnect::doConnect();

			$this->sql = "select response from {$this->tableMaster} where "
			. "chain='$this->handle'";

			if ($result = $this->hookup->qeury($this->sql))
			{
				$row = $result->fetch_assoc();
				echo $row['response'];
			}
			$this->hookup->close();
		} 
		elseif ($this->successor != null)
		{
			$this->successor->handleRequest($request);
		}
	}
}