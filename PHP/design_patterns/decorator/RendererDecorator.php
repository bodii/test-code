<?php

// 命名空间
namespace design_patterns\decorator;

/**
 * RendererDecorator class
 * 抽象渲染类
 */
abstract class RendererDecorator implements RendererInterface
{
    /**
     * Wrapped variable
     * 实例化后要装饰器的对象
     *
     * @var RendererInterface
     */
    protected $wrapped;

    /**
     * __construct function
     * 初始化这个装饰器对象
     *
     * @param RendererInterface $wrapped
     */
    public function __construct(RendererInterface $wrapped)
    {
        $this->wrapped = $wrapped;
    }
}