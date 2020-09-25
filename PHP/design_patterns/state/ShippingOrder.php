<?php

// 命名空间：状态设计模式
namespace design_patterns\state;

/**
 * ShippingOrder class
 * 加入订单
 */
class ShippingOrder extends StateOrder
{
    /**
     * __construct function
     * 实例化方法，加入一个订单
     */
    public function __construct()
    {
        $this->setStatus('shipping');
    }

    /**
     * Dome function
     * 订单添加完成
     *
     * @return void
     */
    protected function done()
    {
        $this->setStatus('completed');
    }
}