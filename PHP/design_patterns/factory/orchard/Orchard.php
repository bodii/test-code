<?php
/**
 * 命名空间 factory\orchard
 */
namespace factory\orchard;

use factory\Factory;
/**
 * Orchard class
 * 果园类
 */
class Orchard implements Factory
{
    /**
     * __construct function
     * 构造方法，用于实例化一个果园类
     */
    public function __construct()
    {
        echo '开始创建果园:)', PHP_EOL;
    }

    /**
     * Produce function
     * 生产者方法， 用于给你什么水果
     * 
     * @param string $fruit 水果种类
     * 
     * @return object|void
     */
    public function produce(string $fruit='')
    {
        switch ($fruit) {
        case 'apple':
            return new Apple();
            break;
        case 'banana':
            return new Banana();
            break;
        default:
            echo '果园里没有种植[ '.$fruit.' ]这种水果!!', PHP_EOL;
            break;
        }   
    }
}