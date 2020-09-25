<?php

// 命名空间：访问者设计模式
namespace design_patterns\visitor;

/**
 * User class
 * 用户类
 */
class User implements Role
{
    /**
     * Name variable
     * 用户名
     *
     * @var string
     */
    private $name;

    /**
     * __construct function
     *
     * @param string $name 用户名
     */
    public function __construct(string $name)
    {
        $this->name = $name;
    }

    /**
     * GetName function
     * 获取用户名
     *
     * @return string
     */
    public function getName(): string
    {
        return sprintf('User: %s', $this->name);
    }

    /**
     * Accept function
     * 接受访问的公共方法
     *
     * @param RoleVisitorInterface $visitor 访问者对象
     * 
     * @return void
     */
    public function accept(RoleVisitorInterface $visitor)
    {
        $visitor->visitUser($this);
    }
}