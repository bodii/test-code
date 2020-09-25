<?php

// 命名空间:观察者设计模式-测试
namespace design_patterns\observer\test;

use design_patterns\observer\User;
use design_patterns\observer\UserObserver;

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
 * ObserverTest class
 * 观察者测试类
 */
class ObserverTest extends TestCase
{
    /**
     * TestChangeInUserLeadsToUserObserverBeingNotified function
     * 测试修改用户并通知用户观察者
     *
     * @return void
     */
    public function testChangeInUserLeadsToUserObserverBeingNotified()
    {
        // 实例化一个观察者
        $observer = new UserObserver();

        // 实例化一个用户
        $user = new User();

        // 给这个用户添加观察者
        $user->attach($observer);

        // 修改这个用户的邮箱地址
        $user->changeEmail('foo@bar.com');

        // 测试-目前用户观察者实例中是否存在有当前这个用户
        $this->assertCount(1, $observer->getChangedUsers());
    }
}

