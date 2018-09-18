<?php

// 命名空间：PHP协程-测试
namespace PHP_CoProcess\test;

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 匿名方法：注册函数
spl_autoload_register(
    function ($callName) {
        require_once dirname(dirname(__DIR__)).DIR_SEP.
        str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\SystemScheduler as SScheduler;
require_once __DIR__.DIR_SEP.'NonBlockingTest.php';

/**
 * SocketServer function
 * socket server 测试
 *
 * @param integer $port 端口号
 * 
 * @return void
 */
function socketServer($port) 
{
    echo "Starting socket server at prot $port...\n";

    $socket = @stream_socket_server("tcp://localhost:$port", $errNo, $errStr);
    if (!$socket) throw new \Exception($errStr, $errNo);
    
    stream_set_blocking($socket, 0);

    while (true) {
        yield waitForRead($socket);
        $clientSocket = stream_socket_accept($socket, 0);
        yield newTask(handleClient($clientSocket));
    }
}

/**
 * NewTask function
 *
 * @param \Iterator $coroutine 迭代生成器
 * 
 * @return void
 */
function newTask(\Iterator $coroutine)
{
    return new SScheduler(
        function (Task $task, SScheduler $scheduler) use ($coroutine) {
            $task->setSendValue($scheduler->newTask($coroutine));
            $scheduler->schedule($task);
        }
    );
}

/**
 * HandleClient function
 * 处理客户端连接
 *
 * @param mixed $socket socket对象
 * 
 * @return void
 */
function handleClient($socket) 
{
    yield waitForWrite($socket);
    $data = fread($socket, 8192);

    $msg = "Received following request:\n\n$data";
    $msgLength = strlen($msg);

    $response = <<<RES
HTTP/1.1 200 OK\r
Content-Type: text/plain\r
Content-Length: $msgLength\r
Connection: close\r
\r
$msg
RES;

    yield waitForWrite($socket);
    fwrite($socket, $response);
    fclose($socket);
}

$scheduler = new SScheduler;
$scheduler->newTask(socketServer(8000));
$scheduler->run();