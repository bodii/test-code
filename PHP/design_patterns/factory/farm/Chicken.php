<?php
/**
 * 命名空间 factory\farm
 */
namespace factory\farm;

/**
 * Chicken class
 * 鸡
 */
class Chicken implements IAnimal
{
    /**
     * __construct function
     * 构造方法 一只鸡的实例
     */
    public function __construct()
    {
        echo '生产了一只鸡', PHP_EOL;
    }
}