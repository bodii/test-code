<?php

// 命名空间：空对象设计模式
namespace design_patterns\nullObject;

/**
 * Service class
 * 服务类
 */
class Service
{
    /**
     * Logger variable
     * 日志记录对象
     *
     * @var LoggerInterface
     */
    private $logger;

    /**
     * __construct function
     *
     * @param LoggerInterface $log 日志记录对象
     */
    public function __construct(LoggerInterface $log)
    {
        $this->logger = $log;
    }

    /**
     * DoSomething function
     * 使用方法
     *
     * @return void
     */
    public function doSomething()
    {
        $this->logger->log('We are in '.__METHOD__);
    }
}