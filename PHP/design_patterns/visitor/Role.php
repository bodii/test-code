<?php

// 命名空间：访问者模式
namespace design_patterns\visitor;

/**
 * Role class
 * 角色抽象类
 */
interface Role
{
    /**
     * Accept function
     * 接受
     *
     * @param RoleVisitorInferface $visitor 访问者
     * 
     * @return void
     */
    public function accept(RoleVisitorInferface $visitor);
}