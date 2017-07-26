<?php

/*
    [ 命令设计模式 ]

    Command Pattern
    将一个请求封装为一个对象，从而使我们可用不同的请求对客户进行参数化;对请求排除
    或者记录请求日志，以及支持可撤销的操作。命令模式是一种对象行为型模式，其别名为
    动作模式（Action Pattern）事务模式（Transaction Pattern）。

    为什么需要命令模式
    1. 使用命令模式，能够让请求发送者与请求接收者消除彼此之间的耦合，让对象之间的
    调用关系更加灵活。

    2. 使用命令模式可以比较容易地设计一个命令队列和宏命令（组合命令），而且新的命令
    可以很容易地加入系统。

*/


// 例子： 电视机和遥控器。电视机是请求的接收者，遥控器是请求的发送者。遥控器上有一
// 些按钮，不同的按钮对应电视机的不同操作。

// 抽象命令角色
abstract class Command
{
    protected $receiver;

    function __construct(TV $receiver)
    {
        $this->receiver = $receiver;
    }

    abstract public function execute();
}

// 具体命令角色 开机命令
class CommandOn extends Command
{
    public function execute()
    {
        echo '执行开机命令：', '<br>';
        $this->receiver->action();
    }
}

// 具体命令角色 关机
class CommandOff extends Command
{
    public function execute()
    {
        echo '执行关机命令：', '<br>';
        $this->receiver->action();
    }
}


// 命令的发送者  遥控
class Invoker
{
    protected $command;

    public function setCommand(Command $command)
    {
        $this->command = $command;
    }

    public function send()
    {
        $this->command->execute();
    }

}

class TV
{
    public $invoker; // 遥控，随购买电视时附赠
    public function __construct()
    {
        $this->invoker = new Invoker;
    }
    public function action()
    {
        echo '接收到命令，执行成功', '<br>';
    }
}

// 实例一个电视
$receiver = new TV;

$invoker = $receiver->invoker;

$commandOn = new CommandOn($receiver);
$commandOff = new CommandOff($receiver);

// 设置命令 --- 按下开机按钮
$invoker->setCommand($commandOn);
// 发送命令
$invoker->send();

// 设置命令 --- 按下关机按钮
$invoker->setcommand($commandOff);
$invoker->send();
