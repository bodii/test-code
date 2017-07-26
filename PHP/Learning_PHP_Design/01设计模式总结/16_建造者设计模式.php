<?php


/*
    [ 建造者设计模式 ]

    Builder Pattern
    将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

    建造者模式是一步一步创建一个复杂的对象，它允许用户只通过指定复杂对象的类型
    和内容就可以创建它们，用户不需要知道内部的具体构建细节。建造者模式属于对象
    创建型模式。根据中文翻译的不同，建造者模式又可称为生成器模式

    为什么需要建造者模式
    1. 对象的生产需要复杂的初始化 ，比如给一大堆类成员属性赋初值，设置一下其他
    的系统环境变量。使用建造者模式可以将这些初始化的工作封装起来。
    2. 对象的生成时可根据初始化的顺序或数据不同，而生成不同角色。
*/

// 实例
/*
    在很多游戏软件中，地图包括天空、地面、背景等组成，人物角色包括人体、服装、
    装备等组成部分，可以使用建造者模式对其进行设计，通过不同的具体建造者创建
    不同类型的地图或人物
*/

class Person
{
    public $age;
    public $speed;
    public $knowledge;
}

// 抽象建造者类
abstract class Builder
{
    public $_person;
    public abstract function setAge();
    public abstract function setSpeed();
    public abstract function setKnowledge();

    public function __construct(Person $person)
    {
        $this->_person = $person;
    }

    public function getPerson()
    {
        // 指挥者去创建这个人。
        new Director($this);
        return $this->_person;
    }
}

// 长者建造者
class OlderBuiler extends Builder
{
    public function setAge()
    {
        $this->_person->age = 70;
    }

    public function setSpeed()
    {
        $this->_person->speed = 'low';
    }

    public function setKnowledge()
    {
        $this->_person->knowledge = 'more';
    }
}

// 小孩建造者
class ChildBuiler extends Builder
{
    public function setAge()
    {
        $this->_person->age = 10;
    }

    public function setSpeed()
    {
        $this->_person->speed = 'fast';
    }

    public function setKnowledge()
    {
        $this->_person->knowledge = 'little';
    }
}

// 建造指挥者
class Director
{
    private $_builder;

    public function __construct(Builder $builder)
    {
        $this->_builder = $builder;
        //指挥者创建 Person
        $this->built();
    }

    public function built()
    {
        $this->_builder->setAge();
        $this->_builder->setSpeed();
        $this->_builder->setKnowledge();
    }
}

// 实例化一个长者建造者 长者模型
$oldB = new OlderBuiler(new Person);

// 实例化一个建造指挥者 并告诉指挥者，我要创建一个老人的样子
// $director = new Director($oldB);

// 得到长者 对象
$older = $oldB->getPerson();

// 实例化一个小孩 小孩模型
$child = new ChildBuiler(new Person);
$childer = $child->getPerson();

var_dump($older);
var_dump($childer);

/*

使用建造者模式时，我们把创建一个person实例的过程分为了两步.
一步是先交给对应角色的建造者，如长者建造者。这样的好处就把角色的属性设置封装了起来，我们不用
在new一个person时，因为要得到一个older角色的实例，而在外面写了一堆$older->age=70。

另一步是交给了一个建造指挥者，调了一个built方法，通过先设置age，再设置Speed的顺序，初始化
这个角色。当然在这个例子中，初始化的顺序，是无所谓的。但是如果对于一个建造汉堡，或是地图，初
始化的顺序不同，可能就会得到不同的结果。
也许，你会说，我直接设置也很方便呀。是的，对于某些情况是这样的。但是如果你考虑，我现在想增加
一个青年人角色呢？如果我现在想让建造有初始化有三种不同的顺序呢？

如果你使用了建造者模式，这两个问题就简单了，增加一个青年人角色，那就增加一个青年年建造者类。
初始化三种不同的顺序，那么就在指挥建造者中增加两种建造方法。
*/
