<?php

include_once '../display_errors.php';

/*
    封装
    封装就是划分一个抽象的诸多元素的过程，这些元素构成该抽象的结构和行为;
    封装的作用就是将抽象的契约接口与其实现分离。
    在编程中，正是因为封装，对象才成为一个对象。

    可见性 是指对类属性的存取（或访问）。
    private(私有)、protected(保护) 和public(公共)。

*/

# private example
class PrivateVis
{
    private $money;

    public function __construct()
    {
        $this->money = 200;
        $this->secret();
    }

    private function secret(){
        echo $this->money;
    }
}

$worker = new PrivateVis();

