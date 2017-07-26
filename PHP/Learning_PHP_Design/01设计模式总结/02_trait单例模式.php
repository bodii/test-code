<?php


/* 使用Trait关键字实现类似于继承单例类的功能 */

trait Singleton
{
    private static $instance = null;

    private function __clone()
    {
        
    }

    public static function getInstance()
    {
        $class = __CLASS__;

        if (!self::$instance instanceof $class)
        {
            self::$instance = new $class;
        }
    }
}

class DB 
{
    private function __construct()
    {
        echo __CLASS__.PHP_EOL;
    }
}

class DBhandle extends DB
{
    use Singleton;

    private function __construct()
    {
        echo '单例模式的实例被构造了';    
    }
}

$handle = DBhandle::getInstance();
