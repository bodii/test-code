<?php

// 命名空间: 数据库-流接口设计模式测试
namespace design_patterns\fluentInterface\test;

use design_patterns\fluentInterface\Sql;

// 目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入 composer vendor autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用 vendor 下 PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * SqlFluentInterfaceTest class
 * 数据库-流接口测试类
 */
class SqlFluentInterfaceTest extends TestCase
{
    /**
     * TestBuildSQL function
     * SQL测试方法
     *
     * @return void
     */
    public function testBuildSQL()
    {
        $query = (new Sql())
            ->select(['foo', 'bar'])
            ->from('foobar', 'f')
            ->where('f.bar = ?');
        
        $this->assertEquals(
            'SELECT foo, bar FROM foobar AS f WHERE f.bar = ?', 
            (string) $query
        );
    }
}