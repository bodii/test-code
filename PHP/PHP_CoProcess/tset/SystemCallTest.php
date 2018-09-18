<?php

// 命名空间：PHP协程-测试
namespace PHP_CoProcess\test;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

spl_autoload_register(
    function ($callName) {
        require_once dirname(dirname(__DIR__)).DIR_SEP.str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\Task;
use PHP_CoProcess\SystemScheduler as SScheduler;

/**
 * SystemCallTest class
 * 系统调用测试类
 */
class SystemCallTest
{
    /**
     * Callback variable
     * 调用对象
     *
     * @var callable
     */
    protected $callback;

    /**
     * __construct function
     * 实例化一个调用对象
     *
     * @param callable $callback
     */
    public function __construct(callable $callback)
    {
        $this->callback = $callback;
    }

    /**
     * __invoke function
     * 当使用调用当前类的实例作为方法调用时，调用此方法
     *
     * @param Task      $task      任务
     * @param Scheduler $scheduler 调用器
     * 
     * @return mixed
     */
    public function __invoke(Task $task, SScheduler $scheduler)
    {
        $callback = $this->callback;
        return $callback($task, $scheduler);
    }
}


/**
 * GetTaskId function
 * 将当前的任务id设置成发送值，并调度这个任务（将这个任务放入队列）
 *
 * @return void
 */
function getTaskId() 
{
    return new SystemCallTest(
        function (Task $task, SScheduler $scheduler) {
            $task->setSendValue($task->getTaskId());
            $scheduler->schedule($task);
        }
    );
}

/**
 * Task function
 * 创建任务
 *
 * @param integer $max 最大任务数
 * 
 * @return void
 */
function task($max)
{
    $tid = (yield getTaskId()); // 这里的括号是为了和php兼容
    for ($i = 1; $i <= $max; ++$i) {
        echo "This is task $tid iteration $i.\n";
        yield;
    }
}

$scheduler = new SScheduler;

$scheduler->newTask(task(10));
$scheduler->newTask(task(5));

$scheduler->run();