<?php

// 命名空间:PHP协同程序(协程)-测试
namespace PHP_CoProcess\test;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 匿名方法引入
spl_autoload_register(
    function ($callName) {
        if ($callName != 'Cenerator') {
            require_once dirname(dirname(__DIR__)).DIR_SEP.str_replace('\\', DIR_SEP, $callName).'.php';
        }
    }
);

use PHP_CoProcess\Task;
use PHP_CoProcess\Scheduler;

/**
 * Task function
 * 任务测试方法一
 *
 * @return void
 */
function task1() 
{
    for ($i = 1; $i <= 10; ++$i) {
        echo "This is task 1 iteration $i.\n";
        yield;
    }   
}

/**
 * Task2 function
 * 任务测试方法二
 *
 * @return void
 */
function task2()
{
    for ($i = 1; $i <= 5; ++$i) {
        echo "This is task 2 iteration $i.\n";
        yield;
    }
}

// 实例化一个任务调度器
$scheduler = new Scheduler;

$scheduler->newTask(task1());
$scheduler->newTask(task2());

$scheduler->run();