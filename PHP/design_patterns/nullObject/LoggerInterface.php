<?php

// 命名空间：空对象设计模式
namespace design_patterns\nullObject;

/**
 * LoggerInterface interface
 * 记录者接口类
 */
interface LoggerInterface
{
    /**
     * Log function
     * 日志方法
     *
     * @param string $str 日志内容
     * 
     * @return void
     */
    public function log(string $str);
}