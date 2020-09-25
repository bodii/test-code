<?php

// 命名空间：备忘录模式
namespace design_patterns\memento;

/**
 * State class
 * 状态类
 */
class State
{
    const STATE_CREATED = 'created';
    const STATE_OPENED = 'opened';
    const STATE_ASSIGNED = 'assigned';
    const STATE_CLOSED = 'closed';

    /**
     * State variable
     *
     * @var string
     */
    private $state;

    /**
     * ValidStates variable
     *
     * @var string[]
     */
    private static $validStates = [
        self::STATE_CREATED,
        self::STATE_OPENED,
        self::STATE_ASSIGNED,
        self::STATE_CLOSED,  
    ];

    /**
     * __construct function
     *
     * @param string $state 状态值
     */
    public function __construct(string $state)
    {
        self::ensureIsValidState($state);

        $this->state = $state;
    }

    /**
     * EnsureIsValidState function
     * 测试状态值是否存在
     *
     * @param string $state 状态值
     * 
     * @return void
     */
    private static function ensureIsValidState(string $state)
    {
        if (!in_array($state, self::$validStates)) {
            throw new \InvalidArgumentException('Invalid state given');
        }
    }

    /**
     * __toString function
     *
     * @return string
     */
    public function __toString(): string
    {
        return $this->state;
    }
}