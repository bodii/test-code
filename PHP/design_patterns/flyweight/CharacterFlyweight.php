<?php

// 命名空间: 享元设计模式
namespace design_patterns\flyweight;

/**
 * CharacterFlyweight class
 * 字符串享元类
 * 用于将字符串格式成指定的格式内容
 */
class CharacterFlyweight implements FlyweightInterface
{
    /**
     * Name variable
     * 名称
     *
     * @var string
     */
    private $name;

    /**
     * __construct function
     * 初始化方法
     * 用于在实例化类时，设置name的初始值
     *
     * @param string $name
     */
    public function __construct(string $name)
    {
        $this->name = $name;
    }

    /**
     * Render function
     * 渲染方法
     * 用于将name和传入的font格式化成指定的格式的内容
     *
     * @param string $font 定义的字体
     * 
     * @return string 返回格式化后的内容
     */
    public function render(string $font): string
    {
        return sprintf('Character %s with font %s', $this->name, $font);
    }
}