<?php

# 类的继承和组合之间的区别

include_once '../display_errors.php';

class DoMath
{
	private $sum;
	private $quotient;

	public function simpleAdd($first, $second)
	{
		$this->sum = ($first + $second);
		return $this->sum;
	}

	public function simpleDivide($dividend, $divisor)
	{
		$this->quotient = ($dividend / $divisor);
		return $this->quotient;
	}
}

class InheritMath extends DoMath
{
	private $textOut;
	private $fullFace;

	public function numToText($num)
	{
		$this->textOut = (string)$num;
		return $this->textOut;
	}

	public function addFace($face, $msg)
	{
		$this->fullFace = "<strong>" . $face . "</strong>: " . $msg;
		return $this->fullFace;
	}

}

class ClientInherit
{
	private $added;
	private $divided;
	private $textNum;
	private $output;

	public function __construct()
	{
		$family = new InheritMath();
		$this->added = $family->simpleAdd(40, 60);
		$this->divided = $family->simpleDivide($this->added, 25);
		$this->textNum = $family->numToText($this->divided);
		$this->output = $family->addFace('Your results', $this->textNum);
		echo $this->output;
	}
}

$worker = new ClientInherit();
