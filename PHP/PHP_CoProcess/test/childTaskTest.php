<?php

// 命名空间：PHP协程-测试-子协程
namespace PHP_CoProcess\test;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

spl_autoload_register(
    function ($callName) {
        require_once dirname(dirname(__DIR__)).DIR_SEP.str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\SystemScheduler as SScheduler;
use PHP_CoProcess\Task;

/**
 * GetTaskId function
 *
 * @return void
 */
function getTaskId()
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) {
            $task->setSendValue($task->getTaskId());
            $scheduler->schedule($task);
        }
    );
}

/**
 * NewTask function
 *
 * @param \Generator $coroutine 迭代生成器
 * 
 * @return void
 */
function newTask(\Generator $coroutine) 
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) use ($coroutine) {
            $task->setSendValue($scheduler->newTask($coroutine));
            $scheduler->schedule($task);
        }
    );
}

/**
 * KillTask function
 *
 * @param integer $tid 任务号
 * 
 * @return void
 */
function killTask($tid)
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) use ($tid) {
            $task->setSendValue($scheduler->killTask($tid));
            $scheduler->schedule($task);
        }
    );
}

/**
 * ChildTask function
 *
 * @return void
 */
function childTask()
{
    $tid = yield getTaskId();
    while (true) {
        echo "Child task $tid still alive!\n";
        yield;
    }
}

/**
 * Task function
 *
 * @return void
 */
function task()
{
    $tid = yield getTaskId();
    $childTid = yield newTask(childTask());

    for ($i = 1; $i <= 6; ++$i) {
        echo "Parent task $tid iteration $i.\n";
        yield;

        if ($i == 3) yield killTask($childTid);
    }
}


$scheduler = new SScheduler;
$scheduler->newTask(task());
$scheduler->run();