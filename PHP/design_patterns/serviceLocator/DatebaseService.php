<?php

// 命名空间-定位服务模式
namespace design_patterns\serviceLocator;

/**
 * DatebaseService class
 * 数据库服务类
 */
class DatebaseService implements DatebaseServiceInterface
{
    /**
     * ConnectionDb function
     * 连接数据方法
     *
     * @param array $config 连接数据需要的配置集合
     * 
     * @return void
     */
    public function connectionDb(array $config)
    {

    }
}