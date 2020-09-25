<?php

// 命名空间-委托设计模式
namespace design_patterns\delegation;

/**
 * JniorDeveloper class
 * 初级开发者类
 */
class JuniorDeveloper
{
    /**
     * WriteBadCode function
     * 编写一般代码方法
     *
     * @return string 编写的代码字符串
     */
    public function writeBadCode(): string
    {
        return 'Some jnior developer generated code...';
    }
}