<?php

// 命名空间： 责任链模式
namespace design_patterns\chainOfResponsibilities;

use design_patterns\chainOfResponsibilities\Handler;

use Psr\Http\Message\RequestInterface;

/**
 * SlowStorage class
 * 慢速处理存储类
 */
class SlowStorage extends Handler
{
    /**
     * Processing function
     * 处理方法
     *
     * @param RequestInterface $request Request对象
     * 
     * @return void
     */
    protected function processing(RequestInterface $request)
    {
        return 'Hello World!';
    }
}