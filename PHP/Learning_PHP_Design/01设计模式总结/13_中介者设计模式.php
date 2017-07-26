<?php

/*
    [ 中介者模式 ]

    Mediator Pattern
    用于一个中介对象来封装一系列的对象交互，中介者使各对象不需要显示地相互引用，
    从而使其耦合松散，而且可以独立地改变它们之间的交互。中介者模式又称为调停者，
    它是一种对象行为型模式。

    为什么需要中介者模式
    1. 中介者模式可以使对象之间的关系数量急剧减少。
    2. 中转作用（结构性）：通过中介者提供的中转作用，各个同事对象就不再需要显
    式引用其他同事，当需要和其他同事进行通信时，通过中介者即可。该中转作用属于
    中介者在结构上的支持。
    3. 协调作用（行为性）：中介者可以更进一步的对同事之间的关系进行封装，同事
    可以一致地和中介者进交互，而不需要指明中介者需要具体怎么做，中介者根据封装
    在自身内部的协调逻辑，对同事的请求进行进一步处理，将同事成员之间的关系行为
    进行分离和封装。该协调作用属于中介者在行为上的支持。
*/


// 抽象同事类 电话机
abstract class Colleague
{
    protected $mediator; // 用于存放中介者

    abstract public function sendMsg($num, $msg);

    abstract public function receiveMsg($num, $msg);

    // 设置中介者
    final public function setMediator(Mediator $mediator)
    {
        $this->mediator = $mediator;
    }
}

// 具体同事类  座机
class Phone extends Colleague
{
    public function sendMsg($num, $msg)
    {
        $self_num = array_search($this, $this->mediator->colleagues);
        echo $self_num . ' -- 发送声音： '. $msg, '<br>';
        $this->mediator->opreation($num, $msg);
    }

    public function receiveMsg($num, $msg)
    {
        echo $num . ' -- 接收声音： '. $msg, '<br>';
    }
}

// 具体同事类  手机
class TelePhone extends Colleague
{
    private $hasOne = true;
    public function sendMsg($num, $msg)
    {
        $self_num = array_search($this, $this->mediator->colleagues);
        echo $self_num . ' -- 发送声音： '. $msg, '<br>';
        $this->mediator->opreation($num, $msg);
    }

    public function receiveMsg($num, $msg)
    {
        if ($this->hasOne)
        {
            echo '手机来电铃声 ----- [接听，开始通话]', '<br>';
        }
        $this->hasOne = false;
        echo $num . ' -- 接收声音 ：' . $msg, '<br>';
    }
}

// 抽象中介者类
abstract class Mediator 
{
    abstract public function opreation($id, $message);
    abstract public function register($id, Colleague $colleague);
}

// 具体中介者类  交换机
class Switches extends Mediator
{
    public  $colleagues = [];

    // 交换机业务处理
    public function opreation($num, $message)
    {
        if (!array_key_exists($num, $this->colleagues))
        {
            echo __class__ . ' -- 交换机内没有此号码信息，无法通话', '<br>';
        } 
        else
        {
            $this->colleagues[$num]->receiveMsg($num, $message);
        }
        
    }

    // 注册号码
    public function register($num, Colleague $colleague)
    {
        if (!in_array($colleague, $this->colleagues))
        {
            $this->colleagues[$num] = $colleague;
        }
        $colleague->setMediator($this);
    }
}


// 实例化座机
$phone = new Phone;

// 实例化手机
$telePhone = new TelePhone;

// 实例化交换机
$switches = new Switches;

$phoneNum = '010-123456';
$telePhoneNum = '18512345678';
// 注册号码
$switches->register($phoneNum, $phone);
$switches->register($telePhoneNum, $telePhone);

// 通话
$phone->sendMsg($telePhoneNum, 'hello world');
$telePhone->sendMsg($phoneNum, '请说普通话');
$telePhone->sendMsg($phoneNum, '你好');