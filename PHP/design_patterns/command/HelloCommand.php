<?php

// 命名空间：命令行设计模式
namespace design_patterns\command;

/**
 * HelloCommand class
 * Hello命令行类
 */
class HelloCommand implements CommandInterface
{
    /**
     * Output variable
     * 输出者对象
     *
     * @var Receiver
     */
    private $output;

    /**
     * __construct function
     * 实例化一个接收对象
     *
     * @param Receiver $console 接收对象
     */
    public function __construct(Receiver $console)
    {
        $this->output = $console;
    }

    /**
     * Execute function
     * 执行接收者方法
     *
     * @return void
     */
    public function execute()
    {
        $this->output->write('Hello World');
    }
}