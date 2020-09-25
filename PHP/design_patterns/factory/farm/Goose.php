<?php
/**
 * 命名空间 factory\farm
 */
namespace factory\farm;


/**
 * Goose class
 * 鹅
 */
class Goose implements IAnimal
{
    /**
     * __construct function
     * 构造方法 一头猪的实例
     */
    public function __construct()
    {
        echo '生产了一只鹅', PHP_EOL;
    }
}