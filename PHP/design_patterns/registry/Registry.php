<?php

// 命名空间: 注册设计模式
namespace design_patterns\registry;

/**
 * Registry abstract class
 * 注册登记抽象类
 */
abstract class Registry
{
    const LOGGER = 'logger';

    /**
     * StoredValues variable
     * 定义存储数组
     *
     * @var array
     */
    private static $storedValues = [];

    /**
     * AllowedKeys variable
     * 定义合法键名数组
     *
     * @var array
     */
    private static $allowedKeys = [self::LOGGER,];

    /**
     * Set function
     * 设置storedValues的键值
     *
     * @param string $key
     * @param mixed  $value
     * 
     * @return void
     */
    public static function set(string $key, $value)
    {
        if (!in_array($key, self::$allowedKeys)) {
            throw new \InvalidArgumentException('Invalid key given');
        }

        self::$storedValues[$key] = $value;
    }

    /**
     * Get function
     * 获取storedValues的键对应的值
     *
     * @param string $key
     * 
     * @return mixed
     */
    public static function get(string $key)
    {
        if (!in_array($key, self::$allowedKeys) || !isset(self::$storedValues[$key])) {
            throw new \InvalidArgumentException('Invalid key given');
        }

        return self::$storedValues[$key];
    }
}