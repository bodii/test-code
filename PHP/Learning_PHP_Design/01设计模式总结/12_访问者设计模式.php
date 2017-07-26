<?php

/* 
     [ 访问者设计模式 ]

     Visitor Pattern
     表示一个作用于某个对象结构中的各元素的操作。它使你可以不改变各元素类的前提下
     定义作用于这些元素的新操作。

     为什么需要访问者模式
     访问者模式能够在不更改对象的情况下为该对象添加新的功能性

*/

// 具体元素
class Superman
{
    public $name;
    public $age = 18;

    public function doSomething()
    {
        echo '我是超人，我会飞', '<br>';
    }

    public function accept(Visitor $visitor)
    {
        $visitor->doSomething($this);
    }

}

// 具体访问者
class Visitor
{
    public function doSomething($object)
    {
        echo '我可以返老还童到' . $object->age , '<br>';
    }
}

$man = new Superman;
$man->doSomething();

// 通过添加访问者，把访问者能力扩展成自已的
$man->accept(new Visitor);