<?php

// 命名空间-委托设计模式
namespace design_patterns\delegation;

/**
 * TeamLead class
 * 团队领导者类
 */
class TeamLead
{
    /**
     * Junior variable
     * 初级开发者对象
     *
     * @var JniorDeveloper
     */
    private $junior;

    /**
     * __construct function
     *
     * @param JuniorDeveloper $junior 初级开发者对象
     */
    public function __construct(JuniorDeveloper $junior)
    {
        $this->junior = $junior;
    }

    /**
     * WriteCode function
     * 编写一般代码方法 
     *
     * @return string 返回编写成功的代码字符串
     */
    public function writeCode(): string
    {
        return $this->junior->writeBadCode();
    }
}