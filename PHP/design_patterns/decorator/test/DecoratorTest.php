<?php

// 命名空间
namespace design_patterns\decorator\test;

use design_patterns\decorator\{
    RendererInterface,
    RenderInXml,
    RenderInJson,
    Webservice,
};

// 定义目录分隔符，兼容win|linux
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入 composer vendor autoload 文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用 vendor 下的phpunit测试类
use PHPUnit\Framework\TestCase;

/**
 * DecoratorTest class
 * 装饰器测试类
 */
class DecoratorTest extends TestCase
{
    /**
     * Service variable
     * 被装饰的渲染服务类实例
     *
     * @var RendererInterface
     */
    private $service;

    /**
     * SetUp function
     * 设置渲染服务实例
     *
     * @return void
     */
    protected function setUp()
    {
        $this->service = new Webservice('foobar');
    }

    /**
     * TestJsonDecorator function
     * 测试渲染生成JSON内容字符串的方法
     *
     * @return void
     */
    public function testJsonDecorator()
    {
        $JsonDecorator = new RenderInJson($this->service);

        $this->assertEquals('"foobar"', $JsonDecorator->renderData());
    }

    /**
     * TestXmlDecorator function
     * 测试渲染生成XML内容字符串的方法
     *
     * @return void
     */
    public function testXmlDecorator()
    {
        $XmlDecorator = new RenderInXml($this->service);

        $this->assertXmlStringEqualsXmlString(
            '<?xml version="1.0"?><content>foobar</content>',
            $XmlDecorator->renderData()
        );
    }
}