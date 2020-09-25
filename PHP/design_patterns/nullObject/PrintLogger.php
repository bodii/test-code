<?php

// 命名空间：空对象设计模式
namespace design_patterns\nullObject;

/**
 * PrintLogger class
 * 打印记录者内容类
 */
class PrintLogger implements LoggerInterface
{
    /**
     * Log function
     * 打印日志方法
     *
     * @param string $str 日志内容
     * 
     * @return void
     */
    public function log(string $str)
    {
        echo $str;
    }
}