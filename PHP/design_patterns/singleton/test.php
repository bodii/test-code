<?php
 
// 设置目录分隔符常量:Win='\',Linux='/'
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 设置设计模式根目录常量: ../../patterns/
if (!defined('PATTERNS_PATH')) {
    define('PATTERNS', dirname(__DIR__).DIR_SEP);
}

// 第一种自加载方法：外部声明
// function autoLoad($className)
// {
//     require_once PATTERNS.str_replace('\\', DIR_SEP, $className).'.php';
// }

// 第一种自加载方法： 注册外部声明方法
// spl_autoload_register('autoLoad');

// 第二种自加载方法：匿名函数
spl_autoload_register(
    function ($className) {
        // $className = 'singleton\Singleton';所以需要将'\'，转为目录分隔符
        require_once PATTERNS.str_replace('\\', DIR_SEP, $className).'.php';
    }
);

// 第三种自加载方法：类方法
// class AutoLoad
// {
//     public static function load($className)
//     {
//         require_once PATTERNS.str_replace('\\', DIR_SEP, $className).'.php';
//     }
// }

// spl_autoload_register(array('AutoLoad', 'load'));

use singleton\Singleton;

$singletonObj = Singleton::getInstace();
$singletonObj->test();
clone $singletonObj;