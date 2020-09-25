<?php

// 设置目录分隔符常量:Win='\',Linux='/'
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 设置设计模式根目录常量: ../../patterns/
if (!defined('PATTERNS_PATH')) {
    define('PATTERNS_PATH', dirname(__DIR__).DIR_SEP);
}

// 匿名方法注册加载，引入AbstractSingleton
spl_autoload_register(
    function($class){
        require_once PATTERNS_PATH.str_replace('\\', '/', $class).'.php';
    }
);

// 使用命名空间
use singleton\AbstractSingleton;


/**
 * SingletonExtends测试类1
 */
class SingletonExtends extends AbstractSingleton
{

}

// 测试1
$s = SingletonExtends::getInstance();
$s::test();


/**
 * SingletonExtends测试类2
 */
class SingletonExtendsTwo extends AbstractSingleton
{
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

// 测试2
$s_two = SingletonExtendsTwo::getInstance();
$s_two::test();