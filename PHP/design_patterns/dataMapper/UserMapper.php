<?php

// 命名空间
namespace design_patterns\dataMapper;


/**
 * UserMapper class
 * 用户关系映射类
 */
class UserMapper
{
    /**
     * Adapter variable
     * 适配对象实例
     *
     * @var StorageAdapter
     */
    private $adapter;

    /**
     * __construct function
     * 初始化方法
     * 设置适配对象实例
     *
     * @param StorageAdapter $storage 储存实例
     */
    public function __construct(StorageAdapter $storage)
    {
        $this->adapter = $storage;
    }

    /**
     * findById function
     * 根据id查找储存实例中对应key的用户信息
     *
     * @param integer $id user key
     * 
     * @return User 返回用户信息[username, email]
     */
    public function findById(int $id): User
    {
        $result = $this->adapter->find($id);

        if ($result === null) {
            throw new \InvalidArgumentException("User #$id not found");
        }

        return $this->mapRowToUser($result);
    }

    /**
     * mapRowToUser function
     * 设置用用户类映射
     *
     * @param array $row 一个数据集[username, email]
     * 
     * @return User 返回用户信息[username, email]
     */
    private function mapRowToUser(array $row): User
    {
        return User::fromState($row);
    }
}

