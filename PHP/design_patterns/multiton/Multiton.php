<?php

// 命名空间
namespace design_patterns\multiton;

/**
 * Multiton class
 * 多例模式类 
 * 此类不可被继承
 * 
 * @category Multiton
 * @package  Design_Patterns
 * @author   bodii <1401097251@qq.com>
 * @license  MIT https://www.github.com/bodii/programming_knowledge_base
 * @link     https://www.github.com/bodii/programming_knowledge_base/blob/master/PHP/design_patterns/multiton/Multiton.php
 */
final class Multiton
{

    const INSTANCE_1 = '1';
    const INSTANCE_2 = '2';

    /**
     * _instances variable
     * 被实例的数组
     *
     * @var array
     */
    private static $_instances = [];

    /**
     * __construct function
     * 构造方法
     * 不可被访问
     */
    private function __construct()
    {

    }

    /**
     * GetInstance function
     * 获取实例
     *
     * @param string $instanceName
     * 
     * @return Multiton
     */
    public static function getInstance(string $instanceName): Multiton
    {
        if (!isset(self::$_instances[$instanceName])) {
            self::$_instances[$instanceName] = new self();
        }

        return self::$_instances[$instanceName];
    }

    /**
     * __clone function
     * 类的clone
     * 不可被使用
     */
    private function __clone()
    {

    }

    /**
     * __wakeup function
     * 类的反序列化
     * 不可被使用
     */
    private function __wakeup()
    {

    }
}
