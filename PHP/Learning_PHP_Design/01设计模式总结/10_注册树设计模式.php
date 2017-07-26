<?php

/*
    [ 注册树设计模式 ]

    Registry Pattern
    注册树模式为应用中经常使用的对象创建一个中央存储器来存放这些对象————通常
    通过一个只包含静态方法的抽象类来实现（或者通过单例模式）。也叫做注册器模式

    为什么需要注册树模式
    解决常用对象的存放问题，实现类似于全局变量的功能。
*/

// 测试类
class User{}


// 注册树类
class Rigistry
{
    protected static $objects; // 用于存放实例

    public static function set($key, $object)
    {
        self::$objects[$key] = $object;
    }

    public static function get($key)
    {
        if (!isset(self::$objects))
        {
            return false;
        }
    }

    // 删除
    public static function _unset($key)
    {
        unset(self::$objects[$key]);
    }
}

$user = new User;
// 存入实例
Registry::set('User', $user);
// 查看实例
var_dump(Registry::get('User'));
// 删除实例
Registry::_unset('User');

/* 
    注册树经常与单例模式一起使用，先查看注册树上是否有该实例，有就直接使用，
    没有就生成一个实例，并挂在对上。有些时候我们还可以这样做，让get方法如果
    get不到实例时就自动new一存放起来，这样我们使用时就不用管有没有存放这个
    实例了，反正没有的话get方法会帮我们存放。
*/
/*
public static function get($key)
{
    if (!isset(self::$objects[$key]))
    {
        self::$objects[$key] = new $key;
    }
    return self::$objects[$key];
}
*/