<?php

// 命名空间：命令行设计模式
namespace design_patterns\command;

/**
 * Invoker class
 * 设置命令行公共输出类
 */
class Invoker
{
    /**
     * Command variable
     * 命令对象
     *
     * @var CommandInterface
     */
    private $command;

    /**
     * SetCommand function
     * 设置命令对象
     *
     * @param CommandInterface $cmd 命令对象
     * 
     * @return void
     */
    public function setCommand(CommandInterface $cmd)
    {
        $this->command = $cmd;
    }

    /**
     * Run function
     * 运行命令行对象的的方法
     *
     * @return void
     */
    public function run()
    {
        $this->command->execute();
    }
}