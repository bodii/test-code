<?php

// 命名空间：策略模式
namespace design_patterns\strategy;

/**
 * DateComparator class
 * 日期比较类
 */
class DateComparator implements ComparatorInterface
{
    /**
     * Compare function
     *
     * @param mixed $a 包含date索引的数组
     * @param mixed $b 包含date索引的数组
     * 
     * @return integer  -1|0|1
     */
    public function compare($a, $b): int
    {
        $aDate = new \DateTime($a['date']);
        $bDate = new \DateTime($b['date']);

        return $aDate <=> $bDate;
    }
}