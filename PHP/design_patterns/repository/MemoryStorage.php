<?php

// 命名空间-资源库模式
namespace design_patterns\repository;

/**
 * MemoryStorage class
 * 内存存储类
 */
class MemoryStorage implements StorageInterface
{
    /**
     * Data variable
     *
     * @var array
     */
    private $data = [];

    /**
     * LastId variable
     *
     * @var integer
     */
    private $lastId = 0;

    /**
     * Persist function
     *
     * @param array $data data
     *
     * @return integer
     */
    public function persist(array $data): int
    {
        $this->lastId++;

        $data['id'] = $this->lastId;
        $this->data[$this->lastId] = $data;

        return $this->lastId;
    }

    /**
     * Retrieve function
     *
     * @param integer $id id
     *
     * @return array
     */
    public function retrieve(int $id): array
    {
        if (!isset($this->data[$id])) {
            throw new \OutOfRangeException(sprintf('No data found for ID %d', $id));
        }

        return $this->data[$id];
    }

    /**
     * Delete function
     *
     * @param integer $id id
     *
     * @return void
     */
    public function delete(int $id)
    {
        if (!isset($this->data[$id])) {
            throw new \OutOfRangeException(sprintf('No data found for Id %d', $id));
        }

        unset($this->data[$id]);
    }
}