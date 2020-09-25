<?php

// 命名空间：状态设计模式
namespace design_patterns\state;

/**
 * CreateOrder class
 * 创建订单
 */
class CreateOrder extends StateOrder
{
    /**
     * __construct function
     * 实例化方法，实例化一个创建订单对象
     */
    public function __construct()
    {
        $this->setStatus('created');
    }

    /**
     * Dome function
     * 完成创建订单
     *
     * @return void
     */
    protected function done()
    {
        // 将订单状态设置为订单运输中
        static::$state = new ShippingOrder();
    }
}