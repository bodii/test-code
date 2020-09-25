<?php

// 命名空间 责任链模式
namespace design_patterns\chainOfResponsibilities;

use Psr\Http\Message\RequestInterface;
use Psr\Http\Message\ResponseInterface;

/**
 * Handler class
 * 处理器抽象方法
 */
abstract class Handler
{
    /**
     * Successor variable
     * 定义继承处理器的实例对象
     * 
     * @var Handler|null
     */
    private $successor = null;

    /**
     * __construct function
     * 构造方法
     * 用于设置处理器对象
     *
     * @param Handler $handler 处理方法
     */
    public function __construct(Handler $handler = null)
    {
        $this->successor = $handler;
    }

    /**
     * Handle function
     * 定义好的处理器方法
     *
     * @param RequestInterface $request Request对象
     * 
     * @return void
     */
    final public function handle(RequestInterface $request)
    {
        $processed = $this->processing($request);

        if ($processed === null) {
            if ($this->successor !== null) {
                $processed = $this->successor->handle($request);
            }
        }

        return $processed;
    }

    /**
     * Processing function
     * 抽象方法，处理方法，用于将传入的request内容做下一步的处理
     *
     * @param RequestInterface $request Request对象
     * 
     * @return void
     */
    abstract protected function processing(RequestInterface $request);
}
