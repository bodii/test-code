<?php

// 命名空间
namespace design_patterns\staticFactory;

/**
 * StaticFactory class
 * 静态工厂类
 * 禁止继承
 */
final class StaticFactory
{
    /**
     * Factory function
     * 静态工厂方法
     *
     * @param string $type 要实例的类的类型名称
     * 
     * @return FormatInterface 返回一个实例后的 FormatInterface 实现的类的实例
     */
    public static function factory(string $type): FormatInterface 
    {
        if ($type == 'number') {
            return new FormatNumber();
        }

        if ($type == 'string') {
            return new FormatString();
        }

        throw new \Exception('Unknow format given');
    }
}