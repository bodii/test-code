<?php

// 命名空间：访问者设计模式
namespace design_patterns\visitor;

/**
 * RoleVisitor class
 * 角色访问者类
 */
class RoleVisitor implements RoleVisitorInterface
{
    /**
     * Visited variable
     * 访问者容器
     *
     * @var Role[]
     */
    private $visited = [];

    /**
     * VisitorUser function
     * 访问用户
     *
     * @param User $role 用户对象
     * 
     * @return void
     */
    public function visitUser(User $role)
    {
        $this->visited[] = $role;
    }

    /**
     * VisitorGroup function
     * 访问用户组
     *
     * @param Group $role 用户组对象
     * 
     * @return void
     */
    public function visitGroup(Group $role)
    {
        $this->visited[] = $role;
    }

    /**
     * GetVisited function
     * 获取访问者有哪些
     *
     * @return array
     */
    public function getVisited(): array
    {
        return $this->visited;
    }

}