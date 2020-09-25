<?php

// 命名空间：状态设计模式-测试
namespace design_patterns\state\test;

use design_patterns\state\ContextOrder;
use design_patterns\state\CreateOrder;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用单元测试类
use PHPUnit\Framework\TestCase;

/**
 * StateTeset class
 * 状态测试类
 */
class StateTest extends TestCase
{
    /**
     * TestCanShipCreatedOrder function
     * 测试是否能运送创建的订单
     *
     * @return void
     */
    public function testCanShipCreatedOrder()
    {
        $order = new CreateOrder();
        $contextOrder = new ContextOrder();

        $contextOrder->setState($order);
        $contextOrder->done();

        $this->assertEquals('shipping', $contextOrder->getStatus());
    }

    /**
     * TestCanCompleteShiippedOrder function
     * 测试是否能完成购买的订单
     *
     * @return void
     */
    public function testCanCompleteShippedOrder()
    {
        $order = new CreateOrder();
        $contextOrder = new ContextOrder();

        $contextOrder->setState($order);
        $contextOrder->done();
        $contextOrder->done();

        $this->assertEquals('completed', $contextOrder->getStatus());
    }
}
