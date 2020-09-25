<?php

// 命名空间：命令行设计模式
namespace design_patterns\command;

/**
 * CommandInterface interface
 * 命令行接口类
 */
interface CommandInterface
{
    /**
     * Execute function
     * 执行方法
     *
     * @return void
     */
    public function execute();
}