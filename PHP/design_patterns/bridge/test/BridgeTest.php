<?php

// 命名空间
namespace design_patterns\bridge\test;

use design_patterns\bridge\{
    HtmlFormatter,
    PlainTextFormatter,
    HelloWorldService,
};

// 定义目录分隔符，兼容win|linux
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autoload 文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的phpunit的测试类
use PHPUnit\Framework\TestCase;

/**
 * BridgeTeset class
 * 桥梁模式没魔类
 */
class BridgeTest extends TestCase
{
    /**
     * TestCanPrintUsingThePlainTextPrinter function
     * 测试字符串格式化类的打印内容方法
     * 
     * @return void
     */
    public function testCanPrintUsingThePlainTextPrinter()
    {
        $service = new HelloWorldService(new PlainTextFormatter());

        $this->assertEquals('Hello World', $service->get());
    }

    /**
     * TestCanPrintUsingTheHtmlPrinter function
     * 测试html格式化类的打印内容方法
     *
     * @return void
     */
    public function testCanPrintUsingTheHtmlPrinter()
    {
        $service = new HelloWorldService(new HtmlFormatter());

        $this->assertEquals('<p>Hello World</p>', $service->get());
    }
}
