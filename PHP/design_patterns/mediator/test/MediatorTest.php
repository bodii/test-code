<?php

// 命名空间-中介者模式-测试
namespace design_patterns\mediator\test;

use design_patterns\mediator\Mediator;
use design_patterns\mediator\Subsystem\{
    Client,
    Server,
    Database,
};

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * MediatorTest class
 * 中介者测试类
 */
class MediatorTest extends TestCase
{
    /**
     * TestOutputHelloWorld function
     * 测试输出Hello World
     *
     * @return void
     */
    public function testOutputHelloWorld()
    {
        $client = new Client();
        new Mediator(new Database(), $client, new Server());

        $this->expectOutputString('Hello World');
        $client->request();
    }
}