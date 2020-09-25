<?php
// 命名空间
namespace design_patterns\multiton\test;

// 使用多多例模式类
use design_patterns\multiton\Multiton;

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autoload文件
require dirname(dirname(dirname(__DIR__))).DIR_SEP.'vendor'.DIR_SEP.'autoload.php';
// 使用PHPUnit单无测试类
use PHPUnit\Framework\TestCase;

/**
 * MultitonTest class
 * 多例模式测试类
 */
class MultitonTest extends TestCase
{
    /**
     * TestException function
     * 测试多例类是否有异常错误
     *
     * @return void
     */
    public function testException()
    {
        // Multiton::getInstance('test1');
    }

    public function testCanCreateTestInstance()
    {
        $this->assertInstanceOf(
            'design_patterns\multiton\Multiton',
            Multiton::getInstance('test')
        );
    }
}

// 实例MultitonTest 多例模式测试类