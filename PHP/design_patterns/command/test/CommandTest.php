<?php

// 命名空间：命令行设计模式-测试
namespace design_patterns\command\test;

use design_patterns\command\HelloCommand;
use design_patterns\command\Invoker;
use design_patterns\command\Receiver;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor 下的 autoload 文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor 下的PHPUnit 单元测试类
use PHPUnit\Framework\TestCase;

/**
 * CommandTest class
 * 命令行测试类
 */
class CommandTest extends TestCase
{
    /**
     * TestHelloCommand function
     * 测试Hello命令行类方法
     *
     * @return void
     */
    public function testHelloCommand()
    {
        $invoker = new Invoker();
        $receiver = new Receiver();

        $invoker->setCommand(new HelloCommand($receiver));
        $invoker->run();
        $this->assertEquals('Hello World', $receiver->getOutput());
    }

}
