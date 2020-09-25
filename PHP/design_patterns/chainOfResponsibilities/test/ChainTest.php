<?php

// 命名空间: 责任链设计模式测试
namespace design_patterns\chainOfResponsibilities\test;

use design_patterns\chainOfResponsibilities\Handler;
use design_patterns\chainOfResponsibilities\FastStorage;
use design_patterns\chainOfResponsibilities\SlowStorage;

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入 composer vendor autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的PHPUnit测试类
use PHPUnit\Framework\TestCase;


/**
 * ChainTest class
 * 链测试类
 */
class ChainTest extends TestCase
{
    /**
     * Chain variable
     * 定义一个责任链对象实例
     *
     * @var [type]
     */
    private $chain;

    /**
     * SetUp function
     * 当调用其它测试方法时，会先调用，生成责任链对象
     *
     * @return void
     */
    protected function setUp()
    {
        $this->chain = new FastStorage(
            ['/foo/bar?index=1' => 'Hello In Memory!'],
            new SlowStorage()
        );
    }

    /**
     * TestCanRequestKeyInFastStorage function
     * 测试请求的key是否存在于快速存储对象中
     *
     * @return void
     */
    public function testCanRequestKeyInFastStorage()
    {
        $uri = $this->createMock('Psr\Http\Message\UriInterface');
        $uri->method('getPath')->willReturn('/foo/bar');
        $uri->method('getQuery')->willReturn('index=1');

        $request = $this->createMock('Psr\Http\Message\RequestInterface');
        $request->method('getMethod')
            ->willReturn('GET');
        $request->method('getUri')->willReturn($uri);

        $this->assertEquals('Hello In Memory!', $this->chain->handle($request));

    }

    /**
     * TestCanRequestKeyInSlowStorage function
     * 测试请求的key是否存在于慢速存储对象中
     *
     * @return void
     */
    public function testCanRequestKeyInSlowStorage()
    {
        $uri = $this->createMock('Psr\Http\Message\UriInterface');
        $uri->method('getPath')->willReturn('/foo/baz');
        $uri->method('getQuery')->willReturn('');

        $request = $this->createMock('Psr\Http\Message\RequestInterface');
        $request->method('getMethod')
            ->willReturn('GET');
        $request->method('getUri')->willReturn($uri);

        $this->assertEquals('Hello World!', $this->chain->handle($request));
    }
}