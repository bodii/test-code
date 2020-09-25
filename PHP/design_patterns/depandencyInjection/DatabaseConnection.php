<?php

// 命名空间
namespace design_patterns\depandencyInjection;

/**
 * DatabaseConnection class
 * 数据数据库连接类
 */
class DatabaseConnection
{
    /**
     * Configuration variable
     * 数据库配置项实例
     *
     * @var DatabaseConfiguration
     */
    private $configuration;

    /**
     * __construct function
     * 初始化一个数据库配置项实例
     *
     * @param DatabaseConfiguration $config
     */
    public function __construct(DatabaseConfiguration $config)
    {
        $this->configuration = $config;
    }

    /**
     * GetDsn function
     * 获取将数据库配置格式化成配置内容的方法
     *
     * @return string
     */
    public function getDsn(): string
    {
        return sprintf(
            '%s:%s@%s:%d',
            $this->configuration->getUsername(),
            $this->configuration->getPassword(),
            $this->configuration->getHost(),
            $this->configuration->getPort()
        );
    }
}