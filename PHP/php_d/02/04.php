<?php

include_once '../display_errors.php';

# 封装
/*
    为了保持封装，同时提供可访问性，OOP设计建议使用获取方法（getters）和
    设置方法（setters），也分别称为存取方法（accessors)和修改方法(mutators)
    不建议直接访问一个类，通过直接赋值来得到或修改属性值，这些工作完全可以
    由获取方法/设置方法来完成。滥用获取方法和设置方法会破坏封装。
*/

class GetSet
{
    private $dataWarehouse;

    function __construct()
    {
        $this->setter(200);
        $got = $this->getter();
        echo $got;
    }

    private function getter()
    {
        return $this->dataWarehouse;
    }

    private function setter($value)
    {
        $this->dataWarehouse = $value;
    }
}

$worker = new GetSet();

# 不要直接请求完成一个工作所需的信息，而应当请求拥有这个信息的对象为你完成工作。
/*
    在上面这个例子中，通过实例化类，
    $worker = new GetSet();
    就很好地做到了这一点。它没有暴露实现细节。不过单独来看，GetSet类似乎
    没有太大用处，因为要指定一个值，唯一的途径就是在类中硬编码设置。
*/


// 从某种程度上讲，设计模式的目的是建立对象之间的通信链路。
// 允许公开访问获取方法和设置方法只会破坏封装
// 保持封装同时保持对象（类）之间通信的过程正是设计模式的一个工作。
