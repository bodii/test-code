<?php

// 命名空间-服务定位模式
namespace design_patterns\serviceLocator;

/**
 * ServiceLocatorInterface interface
 * 服务定位接口类
 */
interface ServiceLocatorInterface
{
    /**
     * Has function
     * 是否存在该服务类
     * 
     * @param string $interface 类名
     * 
     * @return boolean
     */
    public function has(string $interface): bool ;

    /**
     * Get function
     * 获取该类名的服务类实例对象
     *
     * @param string $interface 类名
     * 
     * @return void
     */
    public function get(string $interface);
}