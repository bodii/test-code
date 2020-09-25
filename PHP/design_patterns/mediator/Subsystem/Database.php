<?php
// 命名空间：中介者设计模式-子派发
namespace design_patterns\mediator\Subsystem;

use design_patterns\mediator\Colleague;

/**
 * Database class
 * 数据库类
 */
class Database extends Colleague
{
    /**
     * GetData function
     *
     * @return string
     */
    public function getData(): string
    {
        return 'World';
    }
}