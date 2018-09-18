<?php

// 命名空间:PHP协程- 测试
namespace PHP_CoProcess\test;

// 非阻塞IO

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

spl_autoload_register(
    function ($callName) {
        require_once dirname(dirname(__DIR__)).DIR_SEP.
        str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\SystemScheduler as SScheduler;
use PHP_CoProcess\Task;

/**
 * WaitForRead function
 *
 * @param mixed $socket socket对象
 * 
 * @return void
 */
function waitForRead($socket)
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) use ($socket) {
            $scheduler->waitForRead($socket, $task);
        }
    );
}

/**
 * WaitForWrite function
 *
 * @param mixed $socket socket对象
 * 
 * @return void
 */
function waitForWrite($socket) 
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) use ($socket) {
            $scheduler->waitForWrite($socket, $task);
        }
    );
}
