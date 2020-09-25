<?php

// 命名空间：策略模式
namespace design_patterns\strategy;

/**
 * IdComparator class
 * id比较类
 */
class IdComparator implements ComparatorInterface
{
    /**
     * Compare function
     * 比较方法
     *
     * @param mixed $a 包含尖id索引的数组
     * @param mixed $b 包含尖id索引的数组
     * 
     * @return integer -1|0|1
     */
    public function compare($a, $b): int
    {
        return $a['id'] <=> $b['id'];
    }
}