<?php

// 命名空间：代理设计模式
namespace design_patterns\proxy;

/**
 * Record class
 * 记录器类
 */
class Record
{
    /**
     * Data variable
     *
     * @var array
     */
    private $data;
    
    /**
     * __construct function
     * 构造方法
     * 用于初始化，设置data属性值
     *
     * @param array $data
     */
    public function __construct(array $data = [])
    {
        $this->data = $data;
    }

    /**
     * __set function
     * 魔术方法
     * 当试图设置一个属性时，调用
     *
     * @param string $name
     * @param string $value
     */
    public function __set(string $name, string $value)
    {
        $this->data[$name] = $value;
    }

    /**
     * __get function
     * 魔术方法
     * 当试图获取一个属性值时，调用 
     *
     * @param string $name
     * @return string
     */
    public function __get(string $name): string
    {
        if (!isset($this->data[$name])) {
             throw new \OutOfRangeException('Invalid name given');
        }

        return $this->data[$name];
    }
}
