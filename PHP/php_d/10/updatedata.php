<?php

include_once "02.php";

class UpdateData
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		// 制定表和链接
		$this->tableMaster = 'helpdesk';
		$this->hookup = UniversalConnect::doConnect();

		// 从HTML表单
		$chain = $this->hookup->real_escape_string($_POST['chain']);
		$response = $this->hookup->real_escape_string($_POST['response']);

		// 创建MySQL语句
		$sql = "UPDATE $this->tableMaster SET response='$response' WHERE
			chain='$chain'";
		if($this->hookup->query($sql))
		{
			printf("Chain query: %s <br> Resqonse %s <br> have been changed
			   and set into %s.", $chain, $response, $this->talbeMaster);
		}
		elseif(($result = $this->hookup->query($sql)) === false)
		{
			printf("Invalid query: %s <br> Whole query: %s <br>", 
					$this->hookup->error, $sql);
			exit();
		}
		$this->hookup->close();
	}
}
$worker = new UpdateDate();
