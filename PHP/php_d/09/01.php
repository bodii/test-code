<?php

# 通用类负责连接，代理模式保证安全

# 重要的接口
# 接口并不是一个用来组织操作的抽象结构，它包含一些具体的数据，实现这个
# 接口的类都可以使用这些数据。

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
		self::$hookup = mysqli_connect(self::$server, self::$user, self::$pass,
			self::$currentDB);

		if(self::$hookup)
		{
			// 调试时删除下一行前面的两个斜线
			// echo "Successful connection to MySQL:";
		}
		elseif (mysqli_connect_error(self::$hookup))
		{
			echo('Here is why it failed: ' . mysqli_connect_error());
		}
		return self::$hookup;
	}
}


// 在程序设计中应尽量避免使用全局变量，而单例就相当于合局变量
// 如果可能应尽量从程序代码中册除static关键字。

// 简单客户

class ConnectClient
{
	private $hookup;

	public function __construct()
	{
		// 一行完成整个连接操作
		$this->hookup = UniversalConnect::doConnect();
	}
}

$worker = new ConnectClient();



