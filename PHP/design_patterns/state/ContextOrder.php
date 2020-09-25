<?php

// 命名空间：状态设计模式
namespace design_patterns\state;

/**
 * ContextOrder class
 * 订单内容类
 */
class ContextOrder extends StateOrder
{
    /**
     * GetState function
     *
     * @return StateOrder
     */
    public function getState(): StateOrder
    {
        return static::$state;
    }

    /**
     * SetState function
     *
     * @param StateOrder $state 状态
     * 
     * @return void
     */
    public function setState(StateOrder $state)
    {
        static::$state = $state;
    }

    /**
     * Done function
     * 完成
     *
     * @return void
     */
    public function done()
    {
        static::$state->done();
    }

    /**
     * GetStatus function
     *
     * @return string
     */
    public function getStatus(): string
    {
        return static::$state->getStatus();
    }
}