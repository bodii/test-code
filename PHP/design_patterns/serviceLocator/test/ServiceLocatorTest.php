<?php

// 命名空间-服务定位模式-测试
namespace design_patterns\serviceLocator\test;

use design_patterns\serviceLocator\{
    ServiceLocatorInterface,
    ServiceLocator,
    LogService
};

// 目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

use PHPUnit\Framework\TestCase;

/**
 * ServiceLocatorTest class
 * 服务定位测试类
 */
class ServiceLocatorTest extends TestCase
{
    /**
     * ServiceLocator variable
     * 定位服务实例对象
     *
     * @var ServiceLocatorInterface
     */
    private $serviceLocator;

    /**
     * SetUp function
     * 每次调用一个测试方法，该方法会首先被调用
     *
     * @return void
     */
    public function setUp()
    {
        $this->serviceLocator = new ServiceLocator();
    }

    /**
     * TestHasServices function
     * 测试是否存LogService服务
     *
     * @return void
     */
    public function testHasServices()
    {
        $this->serviceLocator->add(LogService::class);

        $this->assertTrue($this->serviceLocator->has(LogService::class));
        $this->assertFalse($this->serviceLocator->has(self::class));
    }

    /**
     * TestGetWillInstantiateLogServiceIfNoInstanceHasBeenCreatedYet function
     * 测试日志服务是否创建成功
     *
     * @return void
     */
    public function testGetWillInstantiateLogServiceIfNoInstanceHasBeenCreatedYet()
    {
        $this->serviceLocator->add(LogService::class, []);
        $logger = $this->serviceLocator->get(LogService::class);

        $this->assertInstanceOf(LogService::class, $logger);
    }
}