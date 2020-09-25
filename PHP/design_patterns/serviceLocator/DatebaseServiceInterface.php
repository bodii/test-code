<?php

// 命名空间-服务定位模式
namespace design_patterns\serviceLocator;

/**
 * DatebaseServiceInterface interface
 * 数据库服务接口类
 */
interface DatebaseServiceInterface
{
    /**
     * ConnectionDb function
     * 连接数据库方法
     *
     * @param array $config 数据库连接配置集合
     * 
     * @return void
     */
    public function connectionDb(array $config);
}