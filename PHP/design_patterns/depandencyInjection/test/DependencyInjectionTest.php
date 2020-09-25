<?php

// 命名空间
namespace design_patterns\depandencyInjection\tets;

use design_patterns\depandencyInjection\DatabaseConfiguration;
use design_patterns\depandencyInjection\DatabaseConnection;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vender autload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的phpunit测试类
use PHPUnit\Framework\TestCase;

/**
 * DependencyInjectionTest class
 * 依赖注入测试类
 */
class DependencyInjectionTest extends TestCase
{
    /**
     * TestDependencyInjection function
     * 测试数据库配置的依赖注入
     *
     * @return void
     */
    public function testDependencyInjection()
    {
        $config = new DatabaseConfiguration('localhost', 3306, 'test', '123456');
        $connection = new DatabaseConnection($config);

        $this->assertEquals('test:123456@localhost:3306', $connection->getDsn());
    }
}