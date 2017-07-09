<?php 
	
/*
	MySQL和PHP设计模式

	代理(Proxy)
	策略(Strategy)
	职责链(Chain of Responsibility)
	观察者(Observer)
 */

/*
	% 通用类负责连接，代理模式保证安全 %
 */

interface IConnectInfo
{
	const HOST = 'Localhost';
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
			// 调试时删除下一行前面的两个斜线
			// echo 'Successful connection to MySQL:';
		}
		elseif (mysqli_connect_error(self::$hookup))
		{
			echo 'Here is why it failed: '. mysqli_connect_error();
		}
		return self::$hookup;
	}

	/*
		除了为连接变量指定静态值，UniversalConnect类中的mysqli变量($hookup)
		也是静态的。这样一来，Client类做出请求时，就可以在一个更大的静态上
		下文中使用这个变量。
	 */
}

/*
	如果正确地实现单例设计模式，它就相当于全局变量。
	之所以避免使用单例，主要原因就在于它们相当于全局变量。
	如果关心全局变量可能带来的影响，可以从所有代码中删除static关键字，另
	外有private可见性的变量都使用$this->格式。

	当然，使用静态变量最大的好处可能是可以轻松地使用这些静态变量的值，而不必
	每次程序中某个部分需要创建一新连接时都实例化类。
 */

class ConnectClient
{
	private $hookup;

	public function __construct()
	{
		$this->hookup = UniversalConnect::doConnect();
	}
}

$worker = new ConnectClient;