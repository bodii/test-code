<?php
/**
 * 命名空间 factory\farm
 */
namespace factory\farm;


/**
 * Duck class
 * 鸭子
 */
class Duck implements IAnimal
{
    /**
     * __construct function
     * 构造方法 一头猪的实例
     */
    public function __construct()
    {
        echo '生产了一只鸭子', PHP_EOL;
    }
}