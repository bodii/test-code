<?php

include_once '01.php';

// 建立登录注册 

class HashRegister
{
	public function __construct()
	{
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();
		$username = $this->hookup->real_escape_string(trim($_POST['uname']));
		$pwNow = $this->hookup->real_escape_string(trim($_POST['pw']));

		$sql = "insert into $this->tableMaster(uname, pw) values("$username",
				md5("$pwNow"))";
		
		if($this->hookup->query($sql))
		{
			echo "Registration completed:";
		}
		elseif (($result = $this->hookup->query($sql)) === false)
		{
			printf("Invalid query: %s <br> Whole query: %s <br>", 
					$this->hookup->error, $sql);
			exit();
		}
		$this->hookup->close();
	}
}

$worker = new HashRegister();

