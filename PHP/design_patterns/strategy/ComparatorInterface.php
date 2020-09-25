<?php

// 命名空间：策略模式
namespace design_patterns\strategy;

/**
 * ComparatorInterface interface
 * 比较接口类
 */
interface ComparatorInterface
{
    /**
     * Compare function
     * 比较方法
     *
     * @param mixed $a 比较值a
     * @param mixed $b 比较值b
     * 
     * @return integer
     */
    public function compare($a, $b): int;
}