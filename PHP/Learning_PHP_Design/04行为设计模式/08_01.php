<?php 

/*
	测试 - 游戏
	1--2--3
	|  |  |
	4-[5]-6
	|  |  |
	7--8--9

*/
class Context
{
	private $moveStup;
	public $currentLocation;

	public function __construct()
	{
		$this->moveStup = new Move($this);
	}

	public function up()
	{
		echo "up the results after the move: <br>";
		$this->moveStup->boundary();
		$this->moveStup->upMove();
	}

	public function down()
	{
		echo "down the results after the move: <br>";
		$this->moveStup->boundary();
		$this->moveStup->downMove();
	}

	public function left()
	{
		echo "left the results after the move: <br>";
		$this->moveStup->boundary();
		$this->moveStup->leftMove();
	}

	public function right()
	{
		echo "right the results after the move: <br>";
		$this->moveStup->boundary();
		$this->moveStup->rightMove();
	}

	public function getLocation()
	{
		$this->moveStup->getLocation();
	}

	public function setLocation($locationNow)
	{
		$this->moveStup->setLocation($locationNow);
	}

	public function initLocation($locationNow = 5)
	{
		echo "currentLocation: $locationNow <br>";
		$this->moveStup->setLocation($locationNow);
	}
}

abstract class IMove 
{
	public $okUp;
	public $okDown;
	public $okLeft;
	public $okRight;
	public $currentLocation;
	abstract public function upMove();
	abstract public function downMove();
	abstract public function leftMove();
	abstract public function rightMove();

	public function boundary()
	{
		switch($this->currentLocation)
		{
			case 1:
				$this->okUp = false;
				$this->okLeft = false;
				break;

			case 2:
				$this->okUp = false;
				break;

			case 3:
				$this->okUp = false;
				$this->okRight = false;
				break;

			case 4:
				$this->okLeft = false;
				break;

			case 5:
				$this->okLeft = true;
				$this->okRight = true;
				$this->okUp = true;
				$this->okDown = true;
				break;

			case 6:
				$this->okRight = false;
				break;

			case 7:
				$this->okLeft = false;
				$this->okDown = false;
				break;

			case 8:
				$this->okDown = false;
				break;

			case 9:
				$this->okRight = false;
				$this->okDown = false;
				break;

			default:
				$this->okLeft = false;
				$this->okRight = false;
				$this->okUp = false;
				$this->okDown = false;
				break;
		}
	}

	public function moveErr(){
		echo '<p><em>不正确的移动<em><br><p/>';
	}
}

class Move extends IMove 
{
	private $stup;

	public function __construct()
	{
		$this->boundary();
	}

	public function upMove()
	{
		if ($this->okUp) {
			$this->currentLocation -= 3;
			echo $this->currentLocation , '<br>';
		} 
		else 
		{
			$this->moveErr();
		}
	}

	public function downMove()
	{
		if ($this->okDown) {
			$this->currentLocation += 3;
			echo $this->currentLocation, '<br>';
		} 
		else 
		{
			$this->moveErr();
		}
	}

	public function leftMove()
	{
		if ($this->okLeft) {
			$this->currentLocation -= 1;
			echo $this->currentLocation, '<br>';
		} 
		else 
		{
			$this->moveErr();
		}
	}

	public function rightMove()
	{
		if ($this->okRight) {
			$this->currentLocation += 1;
			echo $this->currentLocation, '<br>';
		} 
		else 
		{
			$this->moveErr();
		}
	}

	public function getLocation()
	{
		return $this->currentLocation;
	}

	public function setLocation($locationNow)
	{
		$this->currentLocation = $locationNow;
	}
}

class Client
{
	private $context;

	public function __construct()
	{
		$this->context = new Context;
		$this->context->setLocation(5);
		$this->context->up();
		$this->context->up();
		$this->context->left();
		$this->context->down();
		$this->context->left();
		$this->context->setLocation(1);
		$this->context->down();
		$this->context->right();
		$this->context->right();
		$this->context->right();
	}
}

$work = new Client;