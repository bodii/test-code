<?php

/* 【 工厂模式 】
    Factor Pattern ，就是负责生成其他对象的类或方法，也叫工厂方法模式。

    抽象工厂模式(Abstract Factor Pattern),可简单理解为工厂模式的升级版。

    为什么需要工厂模式
    1.工厂模式可以将对象的生产从直接new一个对象，改变通过调用一个工厂方法
    生产。这样的封装，代码若需修改new的对象时，不需修改多处new语句，只需要
    改生产对象方法。
    2.若所需实例化的对象可选择来自不同的类，可省略if-else多层判断，给工厂
    方法传入对应的参数，利用多态性，实例化对应的类。

 */

// 工厂类
class Factor
{
    static function createDB()
    {
        echo '我生产了一个DB实例';
        return new DB;
    }
}

class DB
{
    public function __construct()
    {
        echo __CLASS__.PHP_EOL;
    }
}

$db = Factor::createDB();


/*
    实现一个运算器
 */

abstract class Operation
{
    abstract public function getVal($i, $j);
}

// 加法
class OperationAdd extends Operation
{
    public function getVal($i, $j)
    {
        return $i + $j;
    }
}

// 减法
class OperationSub extends Operation
{
    public function getVal($i, $j)
    {
        return $i - $j; 
    }
}

// 计数器工厂
class CountFactor
{
    private static $operation;

    // 工厂生产特定类对象方法
    static function createOperation(string $operation)
    {
        switch($operation)
        {
            case '+' :
                self::$operation = new OperationAdd;
                break;
            
            case '-':
                self::$operation = new OperationSub;
                break;
        }
        return self::$operation;
    }
}

$counter = CounterFactor::createOperation('+');
echo $counter->getVal(1,2);

/*
    缺点：若是再增加一个乘法运算，除了增加一个乘法运算之外，还得去工厂方法里面
    添加对就的case代码，违反了开放-封闭原则。
 */

// 解决方法1；通过传入指定类名
// 计算器工厂
class CountFactor
{
    static function createOperation(string $operation)
    {
        return new $operation;
    }
}

class OperationMul extends Operation
{
    public function getval($i, $j)
    {
        return $i * $j;
    }
}

$counter = CounterFactor::createOperation('OperationMul');


// 解决方法2：通过抽象工厂模式
abstract class Operation
{
    abstract public function getVal($i, $j);
}

// 加法类
class OperationAdd extends Operation
{
    public function getVal($i, $j)
    {
        return $i + $j;
    }
}

// 减法类
class OperationSub extends Operation
{
    public function getVal($i, $j)
    {
        return $i - $j;
    }
}

// 乘法
class OperationMul extends Operation
{
    public function getVal($i, $j)
    {
        return $i * $j;
    }
}

// 抽象工厂类
abstract class Factor
{
    abstract static function getInstance();
}

// 加法器生产工厂
class AddFactor extends Factor
{
    static function getInstance()
    {
        return new OperationAdd;
    }
}

// 减法器生产工厂
class SubFactor extends Factor
{
    static function getInstance()
    {
        return new OperationSub;
    }
}

// 乘法器生产工厂
class MulFactor extends Factor
{
    static function getInstance()
    {
        return new OperationMul;
    }
}

// 文体输入器生产工厂
class TextFactor extends Factor
{
    static function getInstance(){}
}

$mul = MulFactor::getInstance();
echo $mul->getVal(1, 2);