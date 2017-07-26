<?php

/*
	建立一个简单的CMS
 */

// CMS表
interface IConnectInfo
{
	const HOST = 'localhost';
	const UNAME = 'uname';
	const PW = 'password';
	const DBNAME = 'dataBase';

	public function doConncet();
}

class UniversalConnect implements IConnectInfo
{
	private static $server = IConnectInfo::HOST;
	private static $currentDB = IConnectInfo::DBNAME;
	private static $user = IConnectInfo::UNAME；
	private static $pass = IConnectInfo::PW;
	private static $hookup;

	public function doConncet()
	{
		self::$hookup = mysqli_connect(self::$server, self::$user, self::$pass,
			self::$currentDB);

		if (self::$hookup)
		{
			// 注释，用于调试
		}
		elseif (mysqli_connect_error(self::$hookup))
		{
			exit('Here\'s why it failed: ' . mysqli_connect_error());
		}
		return self::$hookup;
	}
}

// 创建表
class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'cms';
		$this->hookup = UniversalConnect::doConncet();

		$drop = "drop table if exists {$this->tableMaster}";

		if ($this->hookup->query($drop) === true)
		{
			printf("Old table %s has been dropped.<br>", $this->tableMaster);
		}

		$sql = "create table {$this->tableMaster} (
			id  int(10),
			dpHeader varchar(50),
			textBlock text,
			imageURL varcahr(60),
			primary key (id))
		";

		if ($this->hookup->query($sql) === true)
		{
			printf("Table {$this->tableMaster} has been created successfully.<br>");
		}

		$this->hookup->close();
	}
}

$worker = new CreateTable;

/* CMS数据输入和更新 */
class DataEntry
{
	private $tableMaster;
	private $hookup;
	private $sql;

	public function __construct()
	{
		$this->tableMaster = 'cms';
		$this->hookup = UniversalConnect::doConncet();

		if ($_POST['dpHeader'])
			$dpHeader = $this->hookup->real_escape_string($_POST['dpHeader']);
		if ($_POST['textBlock'])
			$textBlock = $this->hookup->real_escape_string($_POST['textBlock']);
		if ($_POST['imageURL'])
			$imageURL = $this->hookup->real_escape_string($_POST['imageURL']);

		$this->sql = "inster into {$this->tableMaster} 
			(dpHeader, textBlock, imageURL) values
			('$dpHeader', '$textBlock', '$imageURL')
		";

		if ($this->hookup->query($this->sql) === true)
		{
			printf("Successful data entry for table : %s <br>", $this->tableMaster);
		}
		elseif ( ($result = $this->hookup->query($this->sql)) === false)
		{
			printf("Invalid query %s <br> Whole query: %s<br>", 
				$this->hookup->error, $this->sql
				);
			exit();
		}
		$this->hookup->close();
	}
}

$worker = new DataEntry;

class DataUpdate
{
	private $tableMaster;
	private $hookup;
	private $sql;

	public function __construct()
	{
		$this->tableMaster = 'cms';
		$this->hookup = UniversalConnect::doConncet();

		if ($_POST['dpHeader'])
			$dpHeader = $this->hookup->real_escape_string($_POST['dpHeader']);
		if ($_POST['newData'])
			$newData = $this->hookup->real_escapee_string($_POST['newData']);

		$changeFiled = 'textBlock';

		$this->sql = "update {$this->tableMaster} set $changeFiled='$newData'
		 where dpHeader='$dpHeader'";

		 if ($result = $this->hookup->query($this->sql))
		 {
		 	echo "$changeFiled changed to:<br> $newData";
		 }
		 else
		 {
		 	echo "Change filed: {$this->hookup->error}";
		 }
	}
}

$worker = new DataUpdate;

// 监听者客户
class SniffClient
{
	private $userAgent;
	private $mobile = false;
	private $deviceObserver;
	private $dpNow;
	private $sub;

	public function __construct()
	{
		if (isset($_POST['dp']))
			$this->dpNow = $_POST['dp'];
		$this->sub = new ConcreteSubject;
		$this->userAgent = $_SERVER['HTTP_USER_AGENT'];
		if (stripos($this->userAgent, 'iphone'))
		{
			$this->mobile = true;
			$this->deviceObserver = new ConcreteObserverPhone;
		}
		if (stripos($this->userAgent, 'android'))
		{
			$this->mobile = true;
			$this->deviceObserver = new ConcreteObserverPhone;
		}

		if (stripos($this->userAgent, 'blackberry'))
		{
			$this->mobile = true;
			$this->deviceObserver = new ConcreteObserverPhone;
		}

		if (stripos($this->userAgent, 'ipad'))
		{
			$this->mobile = true;
			$this->deviceObserver = new ConcreteObserverTable;
		}

		if (!$this->mobile)
		{
			$this->deviceObserver = new ConcreteObserverDT;
		}

		$this->sub->attachObser($this->deviceObserver);
		$this->sub->setState($this->dpNow);
	}
}

$worker = new SniffClient;