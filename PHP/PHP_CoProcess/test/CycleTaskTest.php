<?php

// 命名空间: PHP协程-测试
namespace PHP_CoProcess\test;

/*
    协程堆栈
*/

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

spl_autoload_register(
    function ($callName) {
        require dirname(dirname(__DIR__)).DIR_SEP.str_replace('\\', DIR_SEP, $callName).'.php';
    }
);

use PHP_CoProcess\Scheduler;
use PHP_CoProcess\SystemScheduler as SScheduler;

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


/**
 * CoroutineReturnValue class
 * 返回值的调用
 */
class CoroutineReturnValue 
{
    /**
     * Value variable
     *
     * @var mixed
     */
    protected $value;

    /**
     * __construct function
     *
     * @param mixed $value 值
     */
    public function __construct($value) 
    {
        $this->value = $value;
    }

    /**
     * GetValue function
     * 获取值
     *
     * @return void
     */
    public function getValue()
    {
        return $this->value;
    }
}


/**
 * Retval function
 *
 * @param mixed $value 值
 * 
 * @return void
 */
function retval($value) 
{
    return new CoroutineReturnValue($value);
}


/**
 * StackedCoroutine function
 * 堆栈生成器
 *
 * @param \Generator $gen 迭代生成器对象
 * 
 * @return void
 */
function stackedCoroutine(\Generator $gen)
{
    $stack = new \SplStack;
    $exception = null;

    for (;;) {
        try {
            if ($exception) {
                $gen->throw($exception);
                $exception = null;
                continue;
            }


            $value = $gen->current();

            if ($value instanceof \Generator) {
                $stack->push($gen);
                $gen = $value;
                continue;
            }

            $isReturnValue = $value instanceof CoroutineReturnValue;
            if (!$gen->valid() || $isReturnValue) {
                if ($stack->isEmpty()) {
                    return;
                }

                $gen = $stack->pop();
                $gen->send($isReturnValue ? $value->getValue(): null);
                continue;
            }

            try {
                $sendValue = yield $gen->key() => $value;
            } catch (\Exception $e) {
                $gen->throw($e);
                continue;
            }

        } catch (\Exception $e) {
            if ($stack->isEmpty()) {
                throw $e;
            }

            $gen = $stack->pop();
            $exception = $e;
        }

    }

}


/**
 * CoSocket class
 * socket 处理类
 */
class CoSocket
{
    /**
     * Socket variable
     *
     * @var mixed
     */
    protected $socket;

    /**
     * __construct function
     *
     * @param mixed $socket socket对象
     */
    public function __construct($socket)
    {
        $this->socket = $socket;
    }

    /**
     * Accept function
     *
     * @return void
     */
    public function accept()
    {
        yield waitForRead($this->socket);
        yield retval(
            new CoSocket(stream_socket_accept($this->socket, 0))
        );
    }

    /**
     * Read function
     * 读取
     *
     * @param integer $size 长度
     * 
     * @return void
     */
    public function read($size)
    {
        yield waitForRead($this->socket);
        yield retval(fread($this->socket, $size));
    }

    /**
     * Write function
     * 写入
     *
     * @param string $string 字符串
     * 
     * @return void
     */
    public function write($string)
    {
        yield waitForWrite($this->socket);
        fwriet($this->socket, $string);
    }

    /**
     * Close function
     * 关闭
     *
     * @return void
     */
    public function close()
    {
        @fclose($this->socket);
    }
}

/**
 * Server function
 *
 * @param integer $port 端口号
 * 
 * @return void
 */
function server($port)
{
    echo "Start socket  server at port $port...\n";

    $socket = @stream_socket_server("tcp://localhost:$port", $errNo, $errStr);
    if (!$socket) throw new \Exception($errstr, $errNo);

    stream_set_blocking($socket, 0);

    $socket = new CoSocket($socket);

    while (true) {
        yield newTask(
            handleClient(yield $socket->accept())
        );
    }
}


/**
 * HandleClient function
 * 处理客户端请求
 *
 * @param CoSocket $socket socket对象
 * 
 * @return void
 */
function handleClient(CoSocket $socket)
{
    $data = yield $socket->read(8192);

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

    yield $socket->write($response);
    yield $socket->close();
}

function generator()
{
    echo "Foo\n";
    try {
        yield;
    } catch (\Exception $e) {
        echo "Exception: {$e->getMessage()}\n";
    }

    echo "Bar\n";
}

$gen = generator();
$gen->rewind();
$gen->throw(new \Exception('Test'));
