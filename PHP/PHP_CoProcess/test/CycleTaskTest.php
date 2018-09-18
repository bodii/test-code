<?php

// 命名空间: PHP协程-测试
namespace PHP_CoProcess\test;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

spl_autoload_register(
    function ($callName) {
        require dirname(dirname(__DIR__)).DIR_SEP.str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\Scheduler;

/**
 * EchoTimes function
 *
 * @param string $msg 内容
 * @param int    $max 最大数
 * 
 * @return void
 */
function echoTimes($msg, $max) 
{
    for ($i = 1; $i <= $max; ++$i) {
        echo "$msg iteration $i\n";
        yield;
    }
}

/**
 * Task function
 * 任务
 *
 * @return void
 */
function task() 
{
    echoTimes('foo', 10);
    echo "---\n";
    echoTimes('bar', 5);
    yield;
}

$scheduler = new Scheduler;
$scheduler->newTask(task()); //　无法调用子协程
$scheduler->run();  // output: ---