<?php

// 命名空间
namespace design_patterns\facade;

/**
 * ComputerFacade class
 * 计算机门面类
 */
class ComputerFacade
{
    /**
     * Os variable
     * 操作系统的实例对象
     *
     * @var OsInterface
     */
    private $os;

    /**
     * Bios variable
     * 输入输出系统的实例对象
     *
     * @var BiosInterface
     */
    private $bios;

    /**
     * __construct function
     * 初始化方法 
     *
     * @param BiosInteface $bios 输入输出对象实例
     * @param OsInterface  $os   系统对象实例
     */
    public function __construct(BiosInterface $bios, OsInterface $os)
    {
        $this->os = $os;
        $this->bios = $bios;
    }

    /**
     * TurnOn function
     * 启动电脑后进入BIOS到OS的过程方法
     *
     * @return void
     */
    public function turnOn()
    {
        $this->bios->execute();
        $this->bios->waitForKeyPress();
        $this->bios->launch($this->os);
    }

    /**
     * TurnOff function
     * 关闭电脑系统后断电BIOS的过程方法
     *
     * @return void
     */
    public function turnOff()
    {
        $this->os->halt();
        $this->bios->powerDown();
    }
}