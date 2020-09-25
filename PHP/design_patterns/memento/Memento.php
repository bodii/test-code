<?php

// 命名空间-备忘录模式
namespace design_patterns\memento;

/**
 * Memento class
 * 备忘录类
 */
class Memento
{
    /**
     * State variable
     *
     * @var string
     */
    private $state;

    /**
     * __construct function
     * 初始化一个状态值
     *
     * @param State $state 状态值
     */
    public function __construct(State $state)
    {
        $this->state = $state;
    }

    /**
     * GetState function
     * 获取状态
     *
     * @return void
     */
    public function getState()
    {
        return $this->state;
    }
}