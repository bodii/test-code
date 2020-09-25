<?php

// 命名空间
namespace design_patterns\facade;

/**
 * OsInterface interface
 * 系统接口类
 */
interface OsInterface
{
    /**
     * Halt function
     * 关闭系统的方法 
     *
     * @return void
     */
    public function halt();

    /**
     * GetName function
     * 获取系统名的方法
     *
     * @return string
     */
    public function getName(): string;
}