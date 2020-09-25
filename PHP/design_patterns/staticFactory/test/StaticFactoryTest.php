<?php

namespace design_patterns\staticFactory\test;

use design_patterns\staticFactory\StaticFactory;

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
 * StaticFactoryTest class
 * 静态工厂测试类
 */
class StaticFactoryTest extends TestCase
{
    /**
     * TsetCanCreateNumberFormatter function
     * 测试 创建 NumberFormater 方法
     *
     * @return void
     */
    public function tsetCanCreateNumberFormatter()
    {
        $this->assertInstanceOf(
            'design_patterns\staticFactory\FormatNumber',
            StaticFactory::factory('number')
        );
    }

    /**
     * TestCanCreateStringFormater function
     * 测试 创建 StringFormater 方法
     *
     * @return void
     */
    public function testCanCreateStringFormatter()
    {
        $this->assertInstanceOf(
            'design_patterns\staticFactory\FormatString',
            StaticFactory::factory('string')
        );
    }

    /**
     * TestException function
     * 测试其它非 number和string的触发异常 方法
     *
     * @return void
     */
    public function testException()
    {
        StaticFactory::factory('object');
    }
}
