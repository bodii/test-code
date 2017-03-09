<?php

/*
1--2--3
|  |  |
4--5--6
|  |  |
7--8--9
 */

interface IMatrix
{
	public function goUp();
	public function goDown();
	public function goLeft();
	public function goRight();
}

class Context
{
	private $cell1;
	private $cell2;
	private $cell3;
	private $cell4;
	private $cell5;
	private $cell6;
	private $cell7;
	private $cell8;
	private $cell9;

	private $currentState;

	public function __construct()
	{
		$this->cell1 = new Cell1State($this);
		$this->cell2 = new Cell2State($this);
		$this->cell3 = new Cell3State($this);
		$this->cell4 = new Cell4State($this);
		$this->cell5 = new Cell5State($this);
		$this->cell6 = new Cell6State($this);
		$this->cell7 = new Cell7State($this);
		$this->cell8 = new Cell8State($this);
		$this->cell9 = new Cell9State($this);

		// 开始状态由开发人员来定
		$this->currentState = $this->cell5;
	}

	// 调用状态方法
	public function doUp()
	{
		$this->currentState->goUp();
	}
	
	public function doDown()
	{
		$this->currentState->goDown();
	}

	public function doLeft()
	{
		$this->currentState->goLeft();
	}	

	public function doRight()
	{
		$this->currentState->goRight();
	}

	// 设置当前状态
	public function setState(IMatrix $state)
	{
		$this->currentState = $state;
	}

	// 获得状态
	public function getCell1State()
	{
		return $this->cell1;
	}

	public function getCell2State()
	{
		return $this->cell2;
	}

	public function getCell3State()
	{
		return $this->cell3;
	}

	public function getCell4State()
	{
		return $this->cell4;
	}
	
	public function getCell5State()
	{
		return $this->cell5;
	}

	public function getCell6State()
	{
		return $this->cell6;
	}

	public function getCell7State()
	{
		return $this->cell7;
	}

	public function getCell8State()
	{
		return $this->cell8;
	}
	
	public function getCell9State()
	{
		return $this->cell9;
	}
}

class Cell1State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		// 不合法的移动
	}

	public function goRight()
	{
		echo "<img src='cells/two.png' alt='2'>";
		$this->context->setState($this->context->getCell2State());
	}

	public function goUp()
	{
		// 不合法的移动
	}

	public function goDown()
	{
		echo "<img src='cells/four.png' alt='4'>";
		$this->context->setState($this->context->getCell4State());
	}
}

class cell2State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/one.png' alt='1'>";
		$this->context->setState($this->context->getCell1State());
	}

	public function goRight()
	{
		echo "<img src='cells/three.png' alt='3'>";
		$this->context->setState($this->context->getCell3State());
	}

	public function goUp()
	{
		// 不合法的移动
	}

	public function goDown()
	{
		echo "<img src='cells/five.png' alt='5'>";
		$this->context->setState($this->context->getCell5State());
	}

}

class Cell3State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/two.png' alt='2'>";
		$this->context->setState($this->context->getCell2State());
	}

	public function goRight()
	{
		// 不合法的移动
	}

	public function goUp()
	{
		// 不合法的移动
	}

	public function goDown()
	{
		echo "<img src='cells/six.png' alt='6'>";
		$this->context->setState($this->context->getCell6State());
	}

}

class Cell4State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		// 不合法的移动
	}

	public function goRight()
	{
		echo "<img src='cells/five.png' alt='5'>";
		$this->context->setState($this->context->getCell5State());
	}

	public function goUp()
	{
		echo "<img src='cells/one.png' alt='1'>";
		$this->context->setState($this->context->getCell1State());
	}

	public function goDown()
	{
		echo "<img src='cells/seven.png' alt='7'>";
		$this->context->setState($this->context->getCell7State());
	}
}

class Cell5State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/four.png' alt='4'>";
		$this->context->setState($this->context->getCell4State());
	}

	public function goRight()
	{
		echo "<img src='cells/six.png' alt='6'>";
		$this->context->setState($this->context->getCell6State());
	}

	public function goUp()
	{
		echo "<img src='cells/two.png' alt='2'>";
		$this->context->setState($this->context->getCell2State());
	}

	public function goDown()
	{
		echo "<img src='cells/eight.png' alt='8'>";
		$this->context->setState($this->context->getCell8State());
	}
}

class Cell6State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/five.png' alt='5'>";
		$this->context->setState($this->context->getCell5State());
	}

	public function goRight()
	{
		// 不合法的移动
	}

	public function goUp()
	{
		echo "<img src='cells/three.png' alt='3'>";
		$this->context->setState($this->context->getCell3State());
	}

	public function goDown()
	{
		echo "<img src='cells/nine.png' alt='9'>";
		$this->context->setState($this->context->getCell9State());
	}
}

class Cell7State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		//  不合法的移动
	}

	public function goRight()
	{
		echo "<img src='cells/night.png' alt='8'>";
		$this->context->setState($this->context->getCell8State());
	} 

	public function goUp()
	{
		echo "<img src='cells/four.png' alt='4'>";
		$this->context->setState($this->context->getCell4State());
	}

	public function goDown()
	{
		// 不合法的移动
	}
}

class Cell8State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/seven.png' alt='7'>";
		$this->context->setState($this->context->getCell7State());
	}

	public function goRight()
	{
		echo "<img src='cells/nine.png' alt='9'>";
		$this->context->setState($this->context->getCell9State());
	}

	public function goUp()
	{
		echo "<img src='cells/five.png' alt='5'>";
		$this->context->setState($this->context->getCell5State());
	}

	public function goDown()
	{
		// 不合法的移动
	}
}

class Cell9State implements IMatrix
{
	private $context;

	public function __construct(Context $contextNow)
	{
		$this->context = $contextNow;
	}

	public function goLeft()
	{
		echo "<img src='cells/night.png' alt='8'>";
		$this->context->setState($this->context->getCell8State());
	}

	public function goRight()
	{
		// 不合法的移动
	}

	public function goUp()
	{
		echo "<img src='cells/six.png' alt='6'>";
		$this->context->setState($tihs->context->getCell6State());
	}

	public function goDown()
	{
		// 不合法的移动
	}
}


class Client
{
	private $context;

	public function __construct()
	{
		$this->context = new Context();
		$this->context ->doUp();
		$this->context->doRight();
		$this->context->doDown();
		echo '<br>';
		$this->context->doDown();
		$this->context->doLeft();
		$this->context->doUp();
	}
}

$worker = new Client();
	
