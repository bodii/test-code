<?php

// 命名空间: 享元设计模式
namespace design_patterns\flyweight;

/**
 * FlyweightFactory class
 * 享元工厂类
 */
class FlyweightFactory implements \Countable
{
    /**
     * Pool variable
     * 实例池
     *
     * @var array
     */
    private $pool = [];

    /**
     * Get function
     * 获取一个字符享元实例
     *
     * @param string $name 名称
     * 
     * @return CharacterFlyweight 返回一个字符串享元实例
     */
    public function get(string $name): CharacterFlyweight
    {
        if (!isset($this->pool[$name])) {
            $this->pool[$name] = new CharacterFlyweight($name);
        }

        return $this->pool[$name];
    }

    /**
     * Count function
     * 统计享元实例的个数
     *
     * @return integer
     */
    public function count(): int
    {
        return count($this->pool);
    }
}