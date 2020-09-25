<?php

// 命名空间:备忘录设计模式
namespace design_patterns\memento;

use design_patterns\memento\State;

/**
 * Ticket class
 * 通往不同状态的凭证类
 * 
 */
class Ticket
{
    /**
     * CurrentState variable
     *
     * @var string
     */
    private $currentState;

    /**
     * __construct function
     * 实例化一个状态类
     */
    public function __construct()
    {
        $this->currentState = new State(State::STATE_CREATED);
    }

    /**
     * Open function
     * 将当前状态设置成 opened
     *
     * @return void
     */
    public function open()
    {
        $this->currentState = new State(State::STATE_OPENED);
    }

    /**
     * Assign function
     * 将当前状态设置成 assigned
     *
     * @return void
     */
    public function assign()
    {
        $this->currentState = new State(State::STATE_ASSIGNED);
    }

    /**
     * Close function
     * 将当前状态设置成 closed
     *
     * @return void
     */
    public function close()
    {
        $this->currentState = new State(State::STATE_CLOSE);
    }

    /**
     * SaveToMemento function
     * 将当前状态保存到备忘录
     *
     * @return Memento
     */
    public function saveToMemento(): Memento
    {
        return new Memento(clone $this->currentState);
    }

    /**
     * RestoreFromMemento function
     * 获取指定备忘录的状态，并设置成当前状态
     *
     * @param Memento $memento 备忘录
     * 
     * @return void
     */
    public function restoreFromMemento(Memento $memento)
    {
        $this->currentState = $memento->getState();
    }

    /**
     * GetState function
     * 获取当前状态
     *
     * @return string
     */
    public function getState(): string
    {
        return $this->currentState;
    }

}