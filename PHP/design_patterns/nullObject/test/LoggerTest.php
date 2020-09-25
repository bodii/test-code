<?php

// 命名空间：空对象设计模式
namespace design_patterns\nullObject\test;

use design_patterns\nullObject\{
    NullLogger,
    PrintLogger,
    Service
};

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用单元测试类
use PHPUnit\Framework\TestCase;

/**
 * LoggerTest class
 * 测试类
 */
class LoggerTest extends TestCase
{
    /**
     * TestNullObject function
     * 测试空记录对象方法
     *
     * @return void
     */
    public function testNullObject()
    {
        $service = new Service(new NullLogger);
        $this->expectOutputString('');
        $service->doSomething();
    }

    /**
     * TestStandarLogger function
     * 测试标准记录对象的方法
     *
     * @return void
     */
    public function testStandarLogger()
    {
        $service = new Service(new PrintLogger);
        $this->expectOutputString('We are in design_patterns\nullObject\Service::doSomething');
        $service->doSomething();
    }
}
