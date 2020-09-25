<?php

// 命名空间：访问者模式
namespace design_patterns\visitor;

/**
 * Group class
 */
class Group implements Role
{
    /**
     * Name variable
     * 组名
     *
     * @var string
     */
    private $name;

    /**
     * __construct function
     *
     * @param string $name 组名
     */
    public function __construct(string $name)
    {
        $this->name = $name;
    }

    /**
     * GetName function
     * 获取组名
     *
     * @return string
     */
    public function getName(): string
    {
        return sprintf('Group: %s', $this->name);
    }

    /**
     * Accept function
     * 接受访问的公共方法
     *
     * @param RoleVisitorInterface $visitor 访问者
     * 
     * @return void
     */
    public function accept(RoleVisitorInterface $visitor)
    {
        $visitor->visitGroup($this);
    }
}