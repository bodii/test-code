<?php

/*
  [ 组合设计模式 ]

  Composite pattern
  将对象组合成树形结构以表示“部分整体”的层次结构。组合模式使得用户对单个对象和组合
  对象的使用具有一致性。组合模式也叫合成模式，有时候又叫做部分整合模式。
*/




/* 实例

  如果我们做一个OA系统，公司的人事管理该如何设计呢。传统的就是树状结构。经理下面有部
  门主管，然后是员工。
*/


class Manager
{
  public $name;
  protected $c_nodes = []; // 存放子节点，部门经理，普通员工等

  public function __construct($name)
  {
    $this->name = $name;
  }

  // 添加部门经理
  public function addGm(Gm $gm)
  {
    $this->c_nodes[] = $gm;
  }

  // 添加普通员工
  public function addStaff(staff $staff)
  {
    $this->c_nodes[] = $staff;
  }

  // 获取全部子节占
  public function get_c_nodes()
  {
    return $this->c_nodes;
  }
}

// 部门经理 就用general manager简写GM
interface Gm
{
  public function add(Staff $staff);
  public function get_c_nodes();
}

// 销售经理
class Sgm implements Gm
{
  public $name;
  protected $c_nodes = [];

  public function __construct($name)
  {
    $this->name = $name;
  }

  // 添加员工
  public function add(Staff $staff)
  {
    $this->c_nodes[] = $staff;
  }

  // 获取子节点
  public function get_c_nodes()
  {
    return $this->c_nodes;
  }

  // 区别于其他经理，销售经理有一个销售方法
  public function sell()
  {
    echo '安利一下我司的产品';
  }
}


// 员工接口
interface staff
{
  public function work();
}

// 销售部员工
class Sstaff implements staff
{
  public $name;

  public function __construct($name)
  { 
    $this->name = $name;
  }

  public function work()
  {
    echo '在销售经理带领下，安利全世界';
  }
}

// 实例化
$manager = new Manager('总经理');
$sgm = new Sgm('销售经理');
$staff = new Sstaff('何在');

// 组装成树
$manager->addGm($sgm);
$sgm->add($staff);

echo 1;
var_dump($manager);
echo '<br>';
var_dump($manager->get_c_nodes());
echo '<br>';
var_dump($sgm->get_c_nodes());
echo $staff->work(), '<br>';
