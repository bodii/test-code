<?php

// 命名空间：代理模式测试
namespace design_patterns\proxy\test;

use design_patterns\proxy\RecordProxy;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor 下PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * RecordProxyTest class
 * 记录器代理测试类
 */
class RecordProxyTest extends TestCase
{
    /**
     * RecordProxy variable
     * RecordProxy的实例对象
     *
     * @var RecordProxy
     */
    private $recordProxy;

    /**
     * setUp function
     * 设置RecordProxy的测度对象实例
     *
     * @return void
     */
    public function setUp()
    {
        $this->recordProxy = new RecordProxy(['test1'=>'value1', 'test2'=>'value2']);
        $this->recordProxy->test3 = 'value3';
    }

    /**
     * TestProxyGetValue function
     * 测试记录代理对值的获取
     *
     * @return void
     */
    public function testProxyGetValue()
    {
        $this->assertEquals('value1', $this->recordProxy->test1);
    }

    /**
     * TestProxyStatus function
     * 测试记录代理对象的属性isDirty状态
     *
     * @return void
     */
    public function testProxyStatus()
    {
        // $recordProxy = $this->createMock(RecordProxy::class);
        // $recordProxy->method('isDirty')
        //     ->willReturn(true);

        // $this->assertEquals(true, $recordProxy->isDirty());
        $this->assertEquals(true, $this->recordProxy->isDirty());
    }
}