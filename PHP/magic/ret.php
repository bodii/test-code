<?php

class Test
{
    protected $state = null;
    public function __construct($state=null)
    {
        $this->state = $state;
    }

    public function __toString()
    {
        if ($this->state) {
            return '1';
        } else {
            return '0';
        }
    }
}

$ret = new Test(true);

var_dump((bool)(string)$ret);
var_dump($ret);
$ret = null;
$ret = new Test();
var_dump((bool)(string)$ret);
var_dump($ret);