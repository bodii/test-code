<?php

/*
    [ 空对象设计模式 ]

    Null Object Pattern
    用一个空对象取代NULL，减少对实例的检查。这样的空对象可以在数据不可用时提供默认
    行为。

    为什么需要空对象模式
    解决在需要一个对象时返回一个null值，使其调用函数出错的情况。
*/

/*
// 测试类
class Person
{
    public function code()
    {
        echo 'code makes me happy', '<br>';
    }
}

// 定义一个生成对象函数，只有PHPer才允许生成对象
function getPerson($name)
{
    if ($name == 'PHPer')
    {
        return new Person;
    }
}

$phper = getPerson('PHPer');
$phper->code();
*/
/* 上面代码会输出"code make me happy"。如果有时候这个函数是别人调用的，它
并没有传入合适的参数呢 */
/*
    $javaer = getPerson('javaer');
    $javaer->code();

    // Uncaught Error: Call to a member function code()
*/

/*
$javaer = getPerson('javaer');
if (!is_null($javaer))
{
    $javaer->code();
}

或

if (is_object($javaer))
{
    $javaer->code();
}

*/

// 使用NullObject模式的话，我们就可以让函数没有返回值时返回一个nullobjec对象。
// 而不是一个null值（没有return默认null值)

// 测试类
class Person
{
    public function code()
    {
        echo 'code makes me happy', '<br>';
    }
}

// 空对象模式
class NullObject
{
    public function __call($methodName, $arg)
    {
        echo 'this is NullObject';
    }
}

// 定义一个生成对象函数，只有PHPer才允许生成对象
function getPerson($name)
{
    if ($name == 'PHPer')
    {
        return new Person;
    }
    else
    {
        return new NullObject;
    }
}

$javaer = getPerson('javaer');
$javaer->code();

/*
    我们可以通过返回一个NullObject对象来取代返回Null，这样就不用在调用
    方法时判断是否为null,而且只要你实现了__call方法，不管真正的对象它原
    来是调用那个方法，NullObject都可以去调用而且不报错（实际上是隐式调用
    了魔术方法__call）.当然，如果你原本的逻辑是返回对象是null的话什么都
    不做，那么你可以让__call什么都不做。或者你也可以让它抛出一个异常。
*/

