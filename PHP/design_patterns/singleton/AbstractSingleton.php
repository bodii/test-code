<?php

namespace singleton;

/**
 * AbstractSingleton 
 * Singleton抽象类
 * 
 * @category AbstractSingleton
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/singleton/AbstractSingleton.php
 */
abstract class AbstractSingleton
{
    /**
     * $instance variable
     * 
     * @var object|null 
     */
    protected static $instance = null;

    /**
     * __construct function
     * 类构造(初始化)方法 禁止访问
     * final private
     * 
     * @retrun void
     */
    final private function __construct()
    {

    }

    /**
     * __clone function
     * 类克隆(复制)方法 禁止访问
     * final private
     *
     * @return void
     */
    final private function __clone()
    {

    }

    /**
     * GetInstance function
     *
     * @return void
     */
    final public static function getInstance()
    {
        if (null !== static::$instance) {
            return static::$instance;
        }
        static::$instance = new static();

        return static::$instance;
    }

    /**
     * Test function
     *
     * @return void string output 
     */
    public static function test()
    {
        echo 'Calling the '.get_called_class().'::test() function extends from abstract '.
        get_class().'::test() function.', PHP_EOL;
    }
}