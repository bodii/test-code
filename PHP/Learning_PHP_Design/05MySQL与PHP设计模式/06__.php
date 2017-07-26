<?php 
	
/*
	策略设计模式
*/

class Client
{
	public function insertData()
	{
		$context = new Context(new DataEntry());
		$context->algorithm();
	}

	public function findData()
	{
		$context = new Context(new SearchData());
		$context->algorithm();
	}

	public function showAll()
	{
		$context = new Context(new DisplayData());
		$context->algorithm();
	}

	public function changeData()
	{
		$context = new Context(new UpdateData());
		$context->algorithm();
	}

	public function killer()
	{
		$context = new Context(new DeleteRecord());
		$context->algorithm();
	}
}

namespace insert;
$trigger = new Client;
$trigger->insertData();

namespace display;
$trigger = new Client;
$trigger->showAll();

namespace find;
$trigger = new Client;
$trigger->findData();

namespace update;
$trigger = new Client;
$trigger->changeData();

namespace kill;
$trigger = new Client;
$trigger->killer();