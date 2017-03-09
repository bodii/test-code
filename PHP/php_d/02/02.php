<?php


# protected example
abstract class ProtectVis
{
    abstract protected function countMoney();
    protected $wage;

    protected function setHourly($hourly)
    {
        $money = $hourly;
        return $money;
    }
}

class ConcreteProtect extends ProtectVis
{
    function __construct()
    {
        $this->countMoney();
    }

    protected function countMoney()
    {
        $this->wage = 'Your hourly wage is $';
        echo $this->wage . $this->setHourly(36);
    }
}

$worker = new ConcreteProtect();