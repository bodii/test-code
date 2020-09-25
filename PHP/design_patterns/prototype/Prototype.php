<?php

/**
 * 命名空间 prototype
 */
namespace prototype;

use prototype\IPrototype;

/**
 * Prototype class
 * 原型类
 * 
 * @category Prototype
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/prototype/Prototype.php
 */
class Prototype implements IPrototype
{
    /**
     * Name variable
     *
     * @var string
     */
    private $_name;

    /**
     * __construct function
     * 构造方法
     *
     * @param string $name 给name属性名赋值
     */
    public function __construct($name='')
    {
        $this->_name = $name;
    }

    /**
     * __set function 
     * 类魔术方法
     * 
     * @param string $name  属性名
     * @param string $value 属性值
     * 
     * @return void 
     */
    public function __set(string $name=‘’, $value='')
    {
        $this->$name = $value;
    }

    /**
     * Copy function
     * 原型克隆方法
     *
     * @return object 类的克隆
     */
    public function copy()
    {
        return clone $this;
    }
}
