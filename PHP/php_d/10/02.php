<?php

# 构建和加载响应表
interface IConnectInfo
{
	const HOST = 'localhost';
	const UNAME = 'phpWorker';
	const PW = 'easyWay';
	const DBNAME = 'dpPatt';

	public function doConnect();
}

class UniversalConnect Implements IConnectInfo
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
			// 你可能想删除这个消息
			echo 'Successful connection to MySQL:<br>';
		}
		elseif(mysqli_connect_error(self::$hookup))
		{
			echo 'Here is why it failed: ' . mysqli_connect_error();
		}

		return self::$hookup;
	}
}

