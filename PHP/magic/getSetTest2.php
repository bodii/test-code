<?php

class A
{
    protected $test_int = 2;
    protected $test_array = array('key' => 'test');
    protected $test_obj;

    public function __construct()
    {
        $this->test_obj = new stdClass();
    }

    public function __get($prop)
    {
        return $this->$prop;
    }

    public function __set($prop, $val)
    {
        $this->$prop = $val;
    }
}

$a = new A();
$a->test_array[] = 'asdf';
$a->test_array += array(1 => 'asdf');
$a->test_array = array('key'=>'asdf') + $a->test_array;
print_r($a->test_array);