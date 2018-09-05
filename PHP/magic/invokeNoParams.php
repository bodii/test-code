<?php

class InvokeNoParams
{
    function __invoke()
    {
        print __METHOD__ . PHP_EOL;
        $i = 1;
        foreach (func_get_args() as $arg) {
            print 'The value of $param'.$i. ' is: '. $arg . PHP_EOL;
            ++$i;
        }
        print PHP_EOL;
    }
}

class InvokeSingleParam
{
    function __invoke($param1)
    {
        print __METHOD__ . PHP_EOL;
        print 'Value of $param1 is: ' . $param1 . PHP_EOL . PHP_EOL;
    }
}

class InvokeMultiParams
{
    function __invoke($param1, $param2, $param3)
    {
        print __METHOD__ . PHP_EOL;
        print 'Value of $param1 is: ' . $param1 . PHP_EOL;
        print 'Value of $param2 is: ' . $param2 . PHP_EOL;
        print 'Value of $param3 is: ' . $param3 . PHP_EOL;
        print PHP_EOL;
    }
}


$no = new InvokeNoParams;
$single = new InvokeSingleParam;
$multi = new InvokeMultiParams;

$no(1, 2, 3);
$single('one param');
$multi('param 1', 'param 2', 'param 3');