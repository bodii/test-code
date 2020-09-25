<?php

// 命名空间：注册登记测试
namespace design_patterns\registry\test;

use design_patterns\registry\Registry;
use stdClass;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入 composer vendor autolaod文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor 下的PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * RegistryTest class
 * 注册登记测试类
 */
class RegistryTest extends TestCase
{
    /**
     * TestSetAndGetLogger function
     * 测试注册和获取记录
     *
     * @return void
     */
    public function testSetAndGetLogger()
    {
        // 获取key
        $key = Registry::LOGGER;
        // 一个类的实例
        $logger = new stdClass();

        // 注册这个实例
        Registry::set($key, $logger);
        // 获取刚注册的实例
        $storedLogger = Registry::get($key);

        $this->assertSame($logger, $storedLogger);
        $this->assertInstanceOf(stdClass::class, $storedLogger);
    }

    /**
     * TestThrowsExecptionWhenTryingToSetInvalidKey function
     * 测试当注册一个非法的key value时是否抛出异常
     *
     * @return void
     */
    public function testThrowsExecptionWhenTryingToSetInvalidKey()
    {
        Registry::set('foobar', new stdClass());
    }

    /**
     * TestThrowsExceptionWhenTryingToGetNotSetKey function
     * 测试获取一个不存在的value时是否抛出异常
     *
     * @return void
     */
    public function testThrowsExceptionWhenTryingToGetNotSetKey()
    {
        Registry::get(Registry::LOGGER);
    }
}
