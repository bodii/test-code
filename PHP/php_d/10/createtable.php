<?php

include_once "02.php";


class CreateTable
{
	private $tableMaster;
	private $hookup;

	public function __construct()
	{
		$this->tableMaster = 'helpdesk';
		$this->hookup = UniversalConnect::doConnect();

		$drop = "DROP TABLE IF EXISTS $this->tableMaster";

		if($this->hookup->query($drop) === true)
		{
			printf("old table %s has been dropped.<br>", $this->tableMater);
		}

		$sql = "CREATE TABLE $this->tableMaster (
				id INT NOT NULL AUTO_INCREMENT,
				chain VARCHAR(3), 
				response TEXT,
				PRIMARY KEY (id)
				)";
		if($this->hookup->query($sql) === true)
		{
			printf("Table $this->tableMaster has been created successfully.
					<br>");
		}
		$this->hookup->close();
	}
}
$worker = new CreateTable(); 
