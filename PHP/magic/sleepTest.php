<?php
class A
{
    private $a;

    public function __construct()
    {
        $this->a = 1;
    }

}


class B extends A
{
    protected $b;

    public function __construct()
    {
        parent::__construct();
        $this->b = 2;
    }

    public function __sleep()
    {
        return array('a', 'b');
    }
}

echo serialize(new B);