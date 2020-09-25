<?php

// 命名空间-服务定位模式
namespace design_patterns\serviceLocator;

/**
 * LogService class
 * 服务日志类
 */
class LogService implements LogServiceInterface
{
    /**
     * Log function
     *
     * @param string $class   服务类名
     * @param string $message 日志信息
     * 
     * @return void
     */
    public function log(string $class, string $message)
    {

    }
}