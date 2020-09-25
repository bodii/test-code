<?php

// 命名空间：空对象设计模式
namespace design_patterns\nullObject;

/**
 * NullLogger class
 * 空日志记录者类
 */
class NullLogger implements LoggerInterface
{
    /**
     * Log function
     * 日志处理方法 这里是空
     *
     * @return void
     */
    public function log(string $str)
    {
        
    }
}