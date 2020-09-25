<?php

// 命名空间-资源库模式
namespace design_patterns\repository;

/**
 * StorageInterface interface
 * 数据处理接口类
 */
interface StorageInterface
{
    /**
     * Persist function
     * 持久化方法
     *
     * @param array $data 要储存的数据数组
     * 
     * @return integer 储存后的id
     */
    public function persist(array $data): int;

    /**
     * Retrieve function
     * 取回数据的方法
     *
     * @param integer $id 存储时的id
     * 
     * @return array 返回存储的数据数组
     */
    public function retrieve(int $id): array;

    /**
     * Delete function
     * 删除数据的方法
     *
     * @param integer $id 要删除数据的id
     * 
     * @return void
     */
    public function delete(int $id);
}