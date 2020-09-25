<?php

// 命名空间
namespace design_patterns\dataMapper\test;

use design_patterns\dataMapper\StorageAdapter;
use design_patterns\dataMapper\User;
use design_patterns\dataMapper\UserMapper;

// 定义目录分隔符，兼容win|linux
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的phpunit TestCase测试类
use PHPUnit\Framework\TestCase;

/**
 * UserDateMapperTest class
 * 用户数据映射测试类
 */
class UserDateMapperTest extends TestCase
{
    /**
     * TestCanMapUserFromStorage function
     * 测试用数据存储的映射关系
     *
     * @return void
     */
    public function testCanMapUserFromStorage()
    {
        $storage = new StorageAdapter(
            [
                1 => ['username'=>'tester', 'email'=>'test@email.com']
            ]
        );

        $mapper = new UserMapper($storage);

        $user = $mapper->findById(1);

        $this->assertInstanceOf(User::class, $user);
    }

    /**
     * TestWillNotMapInvalidData function
     * 测试错误数据的映射方法
     * 输出 InvalidArgumentException: User #1 not found
     *
     * @return void
     */
    public function testWillNotMapInvalidData()
    {
        $storage = new StorageAdapter([]);
        $mapper = new UserMapper($storage);

        $mapper->findById(1);
    }
}
