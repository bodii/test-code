<?php

// 命名空间-委托设计模式-测试
namespace design_patterns\delegation\test;

use design_patterns\delegation\{
    TeamLead,
    JuniorDeveloper
};

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload.php
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用PHPUnit的TestCase测试类
use PHPUnit\Framework\TestCase;

/**
 * DelegationTest class
 * 委托模式测试类
 */
class DelegationTest extends TestCase
{
    /**
     * TestHowTeamLeadWriteCode function
     * 团队领导如何写代码方法
     *
     * @return void
     */
    public function testHowTeamLeadWriteCode()
    {
        $junior = new JuniorDeveloper();
        $teamLead = new TeamLead($junior);

        $this->assertEquals($junior->writeBadCode(), $teamLead->writeCode());
    }
}