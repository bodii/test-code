<?php

// 命名空间
namespace design_patterns\decorator;

/**
 * Webservice class
 * web服务类
 * 用于将字符串数据通过装饰器类，转换成web格式
 */
class Webservice implements RendererInterface
{
    /**
     * Data variable
     * 要渲染的字符串数据
     *
     * @var string
     */
    protected $data;

    /**
     * __construct function
     * 初始化类方法
     *
     * @param string $data 传入要渲染的字符串数据
     */
    public function __construct(string $data) 
    {
        $this->data = $data;
    }

    /**
     * RenderData function
     * 渲染数据的方法
     *
     * @return string 返回通过装饰器类渲染后的数据
     */
    public function renderData(): string
    {
        return $this->data;
    }
}