<?php 

// 命名空间: 代理设计模式
namespace design_patterns\proxy;

/**
 * RecordProxy class
 * 记录器代理类
 */
class RecordProxy extends Record
{
    /**
     * IsDirty variable
     * 是一个标记，记录数据是否有数据内容的
     *
     * @var boolean
     */
    private $isDirty = false;

    /**
     * IsInitialized variable
     * 标记，是否数据被初始化过
     *
     * @var boolean
     */
    private $isInitialized = false;

    /**
     * __construct function
     * 构造方法
     * 用于在初始化时，设置data属性的值
     *
     * @param array $data
     */
    public function __construct(array $data)
    {
        // 将数据交给父类去操作
        parent::__construct($data);

        // 设置标记的值
        if (count($data) > 0) {
            $this->isInitialized = true;
            $this->isDirty = true;
        }
    }

    /**
     * __set function
     * 魔术方法
     * 用于设置属性的值，这里设置交给父类来处理
     *
     * @param string $name
     * @param string $value
     */
    public function __set(string $name, string $value)
    {
        $this->isDirty = true;
        parent::__set($name, $value);
    }

    /**
     * isDirty function
     * 获取$data是否有数据
     * 
     * @return boolean
     */
    public function isDirty(): bool
    {
        return $this->isDirty;
    }
}