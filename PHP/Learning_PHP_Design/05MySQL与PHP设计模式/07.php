<?php 
	
/*
	Context类和Strategy接口
	在状态模式设计中，Context类相当于一个“跟踪者”(track keeper), 它会跟踪当
	前的状态。在策略模式设计中，Context则有完全不同的功能，用于将请求与具体
	策略分离，使用策略和请求可以独立地工作。这体现了请求证书后果之间的另外一
	种松绑定。与此同时，它还有利于从Client发出请求。
*/

/*
	Context不是一个接口(既不是抽象类也不是接口)，不过它与Strategy接口有聚合关
	系。“四人帮”指定了以下特征：

	用一个具体策略对象来配置。
	维护Strategy对象的一个引用。
	可以定义一个接口，允许Strategy访问其数据。
 */

class Context
{
	private $strategy;

	public function __construct(IStrategy $strategy)
	{
		$this->strategy = $strategy;
	}

	public  function algorithm()
	{
		$this->strategy->algorithm();
	}
}

interface IStrategy
{
	public function algorithm();
}

class DataEntry implements IStrategy
{
	public function algorithm()
	{
		$hookup = UniversalConnect::doConnect();
		$test = $hookup->real_escape_string($_POST['data']);
		echo "This data has been entered: " . $test .'<br>';
	}
}

class DisplayData implements IStrategy
{
	public function algorithm()
	{
		$hookup = UniversalConnect::doConnect();
		$test = 'Here\'s all the data!!';
		echo $test . '<br>';
	}
}

class SearchData implements IStrategy 
{
	public function algorithm()
	{
		$hookup = UniversalConnect::doConnect();
		$test = $hookup->real_escape_string($_POST['data']);
		echo 'Here\'s what you were looking for ' . $test . '<br>';
	}
}

class UpdateData implements IStrategy 
{
	public function algorithm()
	{
		$hookup = UniversalConnect::doConnect();
		$test = $hookup->real_escape_string($_POST['data']);
		echo 'Your new data is now ' . $test . '<br>';
	}
}

class DeleteRecord implements IStrategy
{
	public function algorithm()
	{
		$hookup = UniversalConnect::doConnect();
		$test = $hookup->real_escape_string($_POST['data']);
		echo 'The record ' . $test .' has been deleted.<br>';
	}
}


// 连接接口和类

interface IConnectInfo
{
	const HOST = 'localhost';
	const UNAME = 'alpha';
	const PW = 'beta';
	const DBNAME = 'gamma';

	public function doConnect();
}

class UniversalConnect implements IConnectInfo
{
	private static $server = IConnectInfo::HOST;
	private static $currentDB = IConnectInfo::DBNAME;
	private static $user = IConnec,tInfo::UNAME;
	private static $pass = IConnectInfo::PW;
	private static $hookup;

	public function doConnect()
	{
		self::$hookup = mysqli_connect(self::$server, self::$user,
		 self::$pass, self::$currentDB);

		if (self::$hookup)
		{
			echo 'Successful connect to MySQl:<br>';
		}
		elseif (mysqli_connect_error(self::$hookup))
		{
			echo 'Here is why it failed: ' . mysqli_connect_error();
		}

		return self::$hookup;
	}
}

// 辅助类
class SecureData
{
	private $changeField;
	private $company;
	private $device;
	private $disappear;
	private $hookup;
	private $lang;
	private $newData;
	private $oldData;
	private $plat;
	private $style;
	private $term;

	// $dataPack将是一个数组
	private $dataPack = [];

	public function enterDate()
	{
		$this->hookup = UniversalConnect::doConnect();
		$this->company = $this->hookup->real_escape_string($_POST['company']);
		$this->devdes = $this->hookup->real_escape_string($_POST['devdes']);
		$this->lang = $this->hookup->real_escape_string($_POST['plat']);
		$this->style = $this->hookup->real_escape_string($_POST['style']);
		$this->device = $this->hookup->real_escape_string($_POST['device']);
		$this->dataPack = [
			$this->company,
			$this->devdes,
			$this->lang,
			$this->plat,
			$this->style,
			$this->device,
		];

		$this->hookup->close();
	}

	public function conductSearch()
	{
		$this->hookup = UniversalConnect::doConnect();
		$this->field = $this->hookup->real_escape_string($_POST['field']);
		$this->term = $this->hookup->real_escape_stirng($_POST['term']);
		$this->dataPack = [
			$this->field,
			$this->term,
		];
		$this->hookup->close();
	}

	public function makeChange()
	{
		$this->hookup = UniversalConnect::doConnect();
		$this->changeField = $this->hookup->real_escape_string($_POST['update']);
		$this->oldData =$this->hookup->real_escape_string($_POST['old']);
		$this->newData = $this->hookup->real_escape->string($_POST['new']);	
		$this->dataPack = [
			$this->changeField,
			$this->oldDate,
			$this->newData
		];
		$this->hookup->close();
	}

	public function removeRecord()
	{
		$this->hookup = UniversalConnect::doConnect();
		$this->disappear = $this->hookup->real_escape_string($_POST['delete']);
		$this->dataPack = [$this->disappear];
		$this->hookup->close();
	}

	// 将安全数据作为数组返回给请求客户
	public function setEntry()
	{
		return $this->dataPack;
	}
}

// 为算法增强参数
interface IStrategy
{
	const TABLENOW = 'survey';
	public function algorithm(Array $dataPack);
}

// 调查表
class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'survey';
		$this->hookup = UniversalConnect::doConnect();

		$drop = 'dorp table if exists ' . $this->tableMaster;

		if ($this->hookup->query($drop) === true)
		{
			printf('Old table %s has been dropped.<br>', $this->tableMaster);
		}

		$sql = 'create table if not exists ' . $this->tableMaster .' (
			id serial,
			company varchar(40),
			devdes varchar(10),
			lang varchar(15),
			plat varchar(15),
			style varchar(20),
			device varchar(10),
			primary key (id)
		)';

		if ($this->hookup->query($sql) === true )
		{
			printf("Table {$this->tableMaster} has been created successfully.<br>");
		}

		$this->hookup->close();
	}
}

$worker = new CreateTable;
