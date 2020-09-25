<?php
/**
 * 命名空间 factory
 */
namespace factory;

use factory\Farm;

/**
 * SampleFactory class
 * 简单工厂类
 * 
 * @category SompleFactory
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/factory/SampleFactory.php
 */
class SampleFactory
{
    /**
     * Produce function
     * 生产方法
     *
     * @param string $type 被生产者的类型，然后根据不同类型，返回实例化不同对象
     * 
     * @return object|void
     */
    public static function produce(string $type='')
    {
        switch ($type) {
        case 'chicken':
            return new Farm\Chicken();
            break;
        case 'pig':
            return new Farm\Pig();
            break;

        default:
            echo '该农场不生产该家畜！！', PHP_EOL;
            break;
        }
    }
}