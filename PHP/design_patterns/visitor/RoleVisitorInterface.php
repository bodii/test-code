<?php

// 命名空间：访问者设计模式
namespace design_patterns\visitor;

/**
 * RoleVisitorInterface interface
 * 角色访问接口类
 */
interface RoleVisitorInterface
{
    /**
     * VisitUser function
     * 访问用户的方法
     *
     * @param User $role 用户对象
     * 
     * @return void
     */
    public function visitUser(User $role);

    /**
     * VisitGroup function
     * 访问用户组方法
     * 
     * @param Group $role 用户组对象
     * 
     * @return void
     */
    public function visitGroup(Group $role);
}