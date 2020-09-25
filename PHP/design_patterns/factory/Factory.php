<?php
/**
 * 命名空间 factory
 */
namespace factory;

/**
 * Factory interface
 * Factory接口
 * 
 * @category Factory
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/factory/Factory.php
 */
interface Factory
{
    /**
     * Produce function
     * 定义生产方法
     * 
     * @return void
     */
    public function produce();
}

