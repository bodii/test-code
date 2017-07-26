<?php

include_once '../display_errors.php';

/*
    抽象类
    抽象类不能实例化，只能由具体类（也就是可以实例化的类）继承抽象类的接口
    以及它的所有具体属性。
*/
abstract class OneTrickAbstract{
  public $storeHere;

  abstract public function trick($whatever);
}


class OneTrickConcrete extends OneTrickAbstract{
    public function trick($whatever){
        $this->storeHere = 'An abstract property';
        return $whatever . $this->storeHere;
    }
}

$worker = new OneTrickConcrete();
echo $worker->trick('From an abstract origin...');
