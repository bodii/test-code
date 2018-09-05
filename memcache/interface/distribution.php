<?php
namespace distribution;

/**
 * IDistribution interface
 * 分配器接口，主要用于分配置虚拟节点
 * 
 */
interface IDistribution
{
    /**
     * Lockup function
     * 查找最近的虚拟节点
     * 
     * @param string $key mem
     * 
     * @return int 返回最近的节点
     */
    public function lookup($key='');
}
