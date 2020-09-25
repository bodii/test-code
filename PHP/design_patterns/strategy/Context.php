<?php

// 命名空间：策略模式
namespace design_patterns\strategy;

/**
 * Context class
 * 上下文类
 */
class Context
{
    /**
     * Comparator variable
     * 比较的方法对象
     *
     * @var ComparatorInterface
     */
    private $comparator;

    /**
     * __construct function
     * 实例化方法，传入一个比较实例
     *
     * @param ComparatorInterface $comparator 比较方法对象实例
     */
    public function __construct(ComparatorInterface $comparator)
    {
        $this->comparator = $comparator;
    }

    /**
     * ExecuteStrategy function
     * 根据不同的比较方法对象执行不同的排序策略
     *
     * @param array $elements 数组元素
     * 
     * @return array
     */
    public function executeStrategy(array $elements): array
    {
        uasort($elements, [$this->comparator, 'compare']);

        return $elements;
    }
}