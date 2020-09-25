<?php

// singleton namespace
namespace singleton;

/**
 * Singleton class
 * 单例设计模式
 * 
 * @category Singleton
 * @package  Design_Patterns
 * @author   bodii <example@empty.no>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/tree/master/PHP/design_patterns/singleton/Singleton.php
 * @method   object getInstance()
 * @see      Singleton->test()
 */
class Singleton
{
    /**
     * $_singletonObj variable
     * 实例化后的单例类自身
     *
     * @var object
     */
    private static $_singletonObj = null;

    /**
     * 构造方法 function
     * 禁止外部访问，也就是new实例化对象
     */
    private function __construct()
    {
    }

    /**
     * 克隆方法 function
     * 当该对象被克隆时，提示警告错误信息
     * 
     * @return void
     */
    public function __clone()
    {
        trigger_error(
            'The '.__CLASS__.' class prohibits the use of clone!', 
            E_USER_WARNING
        );
    }

    /**
     * GetInstace function
     * 获取单例对象本身
     *
     * @return ojbect $_singletonObj 单例对象本身
     */
    public static function getInstace()
    {
        if (!(self::$_singletonObj instanceof self)) {
            self::$_singletonObj = new self;
        }

        return self::$_singletonObj;
    }

    /**
     * Test function
     * 公共测试方法
     * 
     * @return void
     */
    public function test()
    {
        echo 'Calling the ', __METHOD__.'() function.', PHP_EOL;
    }

}

// $singleton = Singleton::getInstace();
// $singleton->test();
// $s = clone $singleton;