<?php

function __autoload($class_name)
{
	include $class_name . '.php';
}

class Client
{
	private $queryNow;

	public function __construct()
	{
		if (isset($_POST['sendNow']))
		{
			$this->queryNow = $_POST['help'];
		}

		$q1 = new Q1();
		$q2 = new Q2();
		$q3 = new Q3();
		$q4 = new Q4();
		$q5 = new Q5();

		// 设置后继
		$q1->setSuccessor($q2);
		$q2->setSuccessor($q3);
		$q3->setSuccessor($q4);
		$q4->setSuccessor($q5);

		// 生成和处理加载请求
		$loadup = new Request($this->queryNow);
		// 设置链首
		$q1->handleRequest($loadup);
	}
}

$makeRequest = new Client();

// client 的一个辅肋类：
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

abstract class Handler
{
	protected $successor;
	protected $hookup;
	protected $tableMaster;
	protected $sql;
	protected $handle;
	abstract public function handleRequest($request);
	abstract public function setSuccessor($nextService);
}

class Q1 extends Handler
{
	public function setSuccessor($nextServic)
	{
		$this->successor = $nextService;
	}

	public function handleRequest($request)
	{
		$this->handle = 'q1';
		if ($request->getService() == $this->handle)
		{
			$this->tableMaster = 'helpdesk';
			$this->hookup = UniversalConnect::doConnect();
			$this->sql = "SELECT response FROM $this->talbeMaster WHERE 
				chain='$this->handle'";
			if($result = $this->hookup->query($this->sql))
			{
				$row = $result->fetch_assoc();
				echo $row['response'];
			}
			$this->hookup->close();
		}
		elseif ($this->successor != NULL )
		{
			$this->successor->handleRequest($request);
		}
	}
}


