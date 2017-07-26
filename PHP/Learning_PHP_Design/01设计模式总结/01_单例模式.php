<?php
	
/*
 【 单例模式(Singleton Pattern)】
 单例模式，顾名思义就是只有一个实例。作为对象的创建模式，单例模式确保某一个
 类只有一个实例，而且自行实例化并向整个系统提供这个实例。
 */

class Singleton
{
    private static $instance = null;

    private function __construct()
    {
        echo '单例模式的实例被构造了';
    }

    private function __clone()
    {
        
    }

    public static function getInstance()
    {
        if (!self::$instance instanceof Singleton)
        {
            self::$instance = new Singleton();
        }
        return self::$instance;
    }
}

$singleton = Singleton::getInstance();


