<?php
/**
 * 命名空间 factory\farm
 */
namespace factory\farm;

use factory\Factory;

/**
 * Farm class
 * 农场类 实现工厂类
 */
class Farm implements Factory
{

    /**
     * __construct function
     * 构造方法 创建农场
     */
    public function __construct()
    {
        echo '开始建造农场:)', PHP_EOL;
    }

    /**
     * Produce function
     * 生产者方法
     * 通过生产者方法，判断生产农畜的类型，然后实例化对应的农畜类
     *
     * @param string $type 生产农畜类型[chicken、duck、goose]
     * 
     * @return void
     */
    public function produce(string $type='')
    {
        switch ($type) {
        case 'chicken':
            return new Chicken();
            break;
        case 'duck':
            return new Duck();
            break;
        case 'goose':
            return new Goose();
            break;

        default:
            echo '农场不喂养[ '.$type.' ]这种农畜!!', PHP_EOL;
            break;
        }
    }
}