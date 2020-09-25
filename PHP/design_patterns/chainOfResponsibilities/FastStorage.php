<?php

// 命名空间： 责任链模式
namespace design_patterns\chainOfResponsibilities;

use design_patterns\chainOfResponsibilities\Handler;
use Psr\Http\Message\RequestInterface;

/**
 * FastStorge class
 * 快速数据存储类
 */
class FastStorage extends Handler
{
    /**
     * Data variable
     *
     * @var array
     */
    private $data;

    /**
     * __construct function
     * 构造方法
     * 用于实现父类初始化的传参数和data的设置
     *
     * @param array $data 
     * 
     * @param Handler $successor 
     */
    public function __construct(array $data, Handler $successor = null)
    {
        parent::__construct($successor);

        $this->data = $data;
    }

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
        $key = sprintf(
            '%s?%s',
            $request->getUri()->getPath(),
            $request->getUri()->getQuery()
        );

        if ($request->getMethod() == 'GET' && isset($this->data[$key])) {
            return $this->data[$key];
        }

        return null;
    }
}