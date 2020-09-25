<?php

// 命名空间-服务定位模式
namespace design_patterns\serviceLocator;

/**
 * LogServiceInterface interface
 * 服务日志接口类
 */
interface LogServiceInterface
{
    /**
     * Log function
     * 记录日志方法
     *
     * @param string $class   类名
     * @param string $message 日志信息
     * 
     * @return void
     */
    public function log(string $class, string $message);
}