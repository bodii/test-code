<?php
/**
 * 命名空间 abstractFactory
 */
namespace abstractFactory;

use abstractFactory\garmentFactory\GarmentType;
use abstractFactory\garmentFactory\SalesPatterns;

/**
 * Factory abstract
 * 工厂抽象类
 * 
 * @category Factory
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/abstractFactory/Factory.php
 */
abstract class Factory
{
    /**
     * Production function
     * 生产产品
     * 
     * @return void
     */
    abstract public static function production(GarmentType $type);

    /**
     * Sales function
     * 销售产品
     *
     * @return void
     */
    abstract public static function sales(SalesPatterns $pattern);
}