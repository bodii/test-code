<?php

// 命名空间：访问者设计模式
namespace design_patterns\visitor\test;

use design_patterns\visitor\{
    User,
    Group,
    RoleVisitor,
    Role,
    RoleVisitorInterface
};

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用单元测试类
use PHPUnit\Framework\TestCase;

/**
 * VisitorTest class
 * 访问者测试类
 */
class VisitorTest extends TestCase
{
    /**
     * Visitor variable
     *
     * @var RoleVisitorInterface
     */
    private $visitor;

    /**
     * SetUp function
     * 基境-模板方法，调用测试方法时这个方法被先调起
     *
     * @return void
     */
    protected function setUp()
    {
        $this->visitor = new RoleVisitor();
    }

    /**
     * ProvideRoles function
     * 测试提供数据
     *
     * @return void
     */
    public function provideRoles()
    {
        return [
            [new User('Dominik')],
            [new Group('Administrators')],
        ];
    }

    /**
     * TestVisitSomeRole function
     *
     * @param Role $role 角色对象
     * 
     * @dataProvider provideRoles 
     * 
     * @return void
     */
    public function testVisitSomeRole(Role $role)
    {
        var_dump($this->visitor);
        // $role->accept($this->visitor);
        // $this->assertSame($role, $this->visitor->getVisitod()[0]);
    }   
}
