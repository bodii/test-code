<?php

// 命名空间
namespace design_patterns\pool\test;

use design_patterns\pool\WorkerPool;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.'vendor'.DIR_SEP.'autoload.php';
use PHPUnit\Framework\TestCase;

/**
 * PoolTest class
 * 对象池测试类
 */
class PoolTest extends TestCase
{
    /**
     * TestCanGetNewInstancesWithGet function
     * 测试获取两个工作者
     *
     * @return void
     */
    public function testCanGetNewInstancesWithGet()
    {
        $pool = new WorkerPool();
        $worker1 = $pool->get();
        $worker2 = $pool->get();

        $this->assertCount(2, $pool);
        $this->assertNotSame($worker1, $worker2);
    }

    /**
     * TestCanGetSameInstanceTwiceWhenDisposingItFirst function
     * 测试获取一个工作者，并退还工作池，再获取另一个工作者
     *
     * @return void
     */
    public function testCanGetSameInstanceTwiceWhenDisposingItFirst()
    {
        $pool = new WorkerPool();
        $worker1 = $pool->get();
        $pool->dispose($worker1);
        $worker2 = $pool->get();

        $this->assertCount(1, $pool);
        $this->assertSame($worker1, $worker2);
    }
}
