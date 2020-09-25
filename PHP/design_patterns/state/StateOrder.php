<?php

// 命名空间：状态设计模式
namespace design_patterns\state;

/**
 * StateOrder class
 * 订单状态操作抽象类
 */
abstract class StateOrder
{
    /**
     * Details variable
     * 订单详情内容
     *
     * @var array
     */
    private $details;

    /**
     * State variable
     *
     * @var string
     */
    protected static $state;

    /**
     * Done function
     *
     * @return void
     */
    abstract protected function done();

    /**
     * SetStatus function
     * 设置订单的状态值
     *
     * @param string $status 订单的状态值
     * 
     * @return void
     */
    protected function setStatus(string $status)
    {
        $this->details['status'] = $status;
        $this->details['updatedTime'] = time();
    }

    /**
     * GetStatus function
     * 获取订单的状态值
     * 
     * @return string
     */
    protected function getStatus(): string
    {
        return $this->details['status'];
    }
    
}