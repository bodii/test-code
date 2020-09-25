<?php

// 命名空间
namespace design_patterns\dataMapper;

/**
 * StorageAdapter class
 * 数据存储适配器类
 */
class StorageAdapter
{
    /**
     * Data variable
     * 数据集
     *
     * @var array
     */
    // private $data = [];
    public $data = [];

    /**
     * __construct function
     * 构造方法
     * 初始化赋值数据
     *
     * @param array $data 数据信息初始化
     * 
     * @return void
     */
    public function __construct(array $data)
    {
        $this->data = $data;
    }

    /**
     * Find function
     * 根据key 查找对应的信息
     *
     * @param integer $id key
     * 
     * @return array|null 返回数据信息 或 null 
     */
    public function find(int $id)
    {
        if (isset($this->data[$id])) {
            return $this->data[$id];
        }
        
        return null;
    }
}