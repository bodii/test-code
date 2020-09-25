<?php

// 命名空间
namespace design_patterns\bridge;

/**
 * Service class
 * 服务抽象类
 */
abstract class Service
{
    /**
     * Implementation variable
     * 要格式化内容的实例
     *
     * @var FormatterInterface
     */
    protected $implementation;

    /**
     * __construct function
     * 实例化服务类
     *
     * @param FormatterInterface $printer 传入要格式化的内容的实例
     */
    public function __construct(FormatterInterface $printer)
    {
        $this->implementation = $printer;
    }

    /**
     * SetImplementation function
     * 设置要格式化内容的实例
     *
     * @param FormatterInterface $printer 传入要格式化的内容的实例
     * 
     * @return void
     */
    public function setImplementation(FormatterInterface $printer)
    {
        $this->implementation = $printer;
    }

    /**
     * Get function
     * 抽象方法 获取格式化后的内容
     *
     */
    abstract public function get();

}