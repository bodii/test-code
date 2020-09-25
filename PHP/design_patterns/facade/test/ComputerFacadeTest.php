<?php

// 命名空间
namespace design_patterns\facade\test;

use design_patterns\facade\{
    ComputerFacade,
    BiosInterface,
    OsInterface
};

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autload
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的phpunit测试类
use PHPUnit\Framework\TestCase;

/**
 * ComputerFacadeTest class
 * 电脑门面测试类
 */
class ComputerFacadeTest extends TestCase
{
    /**
     * TestOn function
     * 测试开机
     *
     * @return void
     */
    public function testOn()
    {
        // createMock 是PHPUnit中提供的用于自动生成对象，
        // 此对象可以充当任意指定原版类型（接口或类名）的测试替身
        // 创建computer OS 测试对象
        $os = $this->createMock(OsInterface::class);

        // $os->method('getName')->will($this->returnValue('Linux'));
        // will($this->returnValue('Linux')) 等价如下：
        $os->method('getName')->willReturn('Linux');

        // 使用getMockBuilder方法可以在测试中生成过程序
        $bios = $this->getMockBuilder(BiosInterface::class)
            ->setMethods(['launch', 'execute', 'waitForKeyPress']) // 要模拟的方法
            ->disableAutoload() // 可用于测试替身类的生成期间禁用__autolaod()
            ->getMock(); // 获取这个虚拟类实例

        $bios->expects($this->once()) // expects用来指明此交互应该是什么样的
            ->method('launch') // 使用方法
            ->with($os); // 对于$os实例对象使用此方法

        $computerFacade = new ComputerFacade($bios, $os);

        $computerFacade->turnOn();

        $this->assertEquals('Linux', $os->getName());
    }
}


