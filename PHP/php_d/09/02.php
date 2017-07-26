<?php

# 代理模式是一种结构型设计模式。
# 4种代理模式：
# 远程代理
# 虚拟代理
# 保护代理
# 智能引用
#
# 代理模式有两个主要的参与者：一个代理主题和一个真实主题。
# 客户通过主题接口向代理提交请求，只有请求通过代理之后，才有可能访问真实
# 主题。


// 建立登录注册

include_once '01.php';

class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'proxyLog';
		$this->hookup = UniversalConnect::doConnect();

		$drop = 'DROP TABLE IF EXISTS ' . $this->tableMaster . ';';

		if($this->hookup->query($drop) === true )
		{
			printf("Old table %s has been dropped.<br>", $this->tableMaster);
		}

		$sql = "CREATE TABLE $this->tableMaster (uname NVARCHAR(15), 
				pw NVARCHAR(120)";
		if($this->hookup->query($sql) === true)
		{
			echo "Table $this->tableMaster has been created successfully.
				<br>";
		}
		$this->hookup->close();
	}
}

$worker = new CreateTable();

