<?php

// 命名空间：PHP协程 - 调度
namespace PHP_CoProcess;

use PHP_CoProcess\Task;
use PHP_CoProcess\Scheduler;

/**
 * Scheduler class
 * 系统任务调度器类
 */
class SystemScheduler extends Scheduler
{
    /**
     * WaitingForRead variable
     * 等待读取容器
     * 
     * @var array
     */
    protected $waitingForRead = [];

    /**
     * WaitingForWrite variable
     * 等待写入容器
     *
     * @var array
     */
    protected $waitingForWrite = [];
    /**
     * Run function
     * 执行任务调度
     *
     * @return void
     */
    public function run()
    {
        while (!$this->taskQueue->isEmpty()) {
            $task = $this->taskQueue->dequeue();
            $retval = $task->run();

            if ($retval instanceof \PHP_CoProcess\test\SystemCallTest) {
                try {
                    $retval($task, $this);
                } catch (\Exception $e) {
                    $task->setException($e);
                    $this->schedule($task);
                }
                continue;
            }

            if ($task->isFinished()) { // 如果当前迭代对象完成
                unset($this->taskMap[$task->getTaskId()]); // 销毁这个任务
            } else {
                $this->schedule($task); // 否则将迭代断点(yield)添加到任务调用度
            }
        }
    }

    /**
     * KillTask function
     * 杀死一个任务的方法
     *
     * @param int $tid 任务号
     * 
     * @return bool
     */
    public function killTask(int $tid): bool
    {
        // 如果这个任务号不存在，直接返回
        if (!isset($this->taskMap[$tid])) {
            return false;
        }

        // 否则从任务容器中将这个任务销毁
        unset($this->taskMap[$tid]);

        // 并且将任务队列中的这个任务号的任务销毁
        foreach ($this->taskQueue as $i => $task) {
            if ($task->getTaskId() === $tid) {
                unset($this->taskQueue[$i]);
                break;
            }
        }

        return true;
    }

    /**
     * WaitForRead function
     * 等待读方法
     *
     * @param mixed $socket socket对象
     * @param Task  $task   task任务对象
     * 
     * @return void
     */
    public function waitForRead($socket, Task $task)
    {
        if (isset($this->waitingForRead[(int)$socket])) {
            $this->waitingForRead[(int)$socket][1][] = $task;
        } else {
            $this->waitingForRead[(int)$socket] = [$socket, [$task]];
        }
    }

    /**
     * WaitForWrite function
     * 等待写方法
     *
     * @param mixed $socket socket对象
     * @param Task  $task   任务对象
     * 
     * @return void
     */
    public function waitForWrite($socket, Task $task)
    {
        if (isset($this->waitingForWrite[(int)$socket])) {
            $this->waitingForWrite[(int)$socket][1][] = $task;
        } else {
            $this->waitingForWrite[(int)$socket] = [$socket, [$task]];
        }
    }

    /**
     * IoPoll function
     * IO轮询的方法
     *
     * @param integer $itemout 过期时间
     * 
     * @return void
     */
    protected function ioPoll($itemout)
    {
        $rSocks = [];
        foreach ($this->waitingForRead as list($socket)) {
            $rSocks[] = $socket;
        }

        $wSocks = [];
        foreach ($this->waitingForWrite as list($socket)) {
            $wSocks[] = $socket;
        }

        $eSocks = [];
        if (!stream_select($rSocks, $wSocks, $eSocks, $timeout)) {
            return;
        }

        foreach ($rSocks as $socket) {
            list(, $tasks) = $this->waitingForRead[(int)$socket];

            unset($this->waitingForRead[(int)$socket]);

            foreach ($tasks as $task) {
                $this->schedule($task);
            }
        }

        foreach ($wSocks as $socket) {
            list(, $tasks) = $this->waitingForWrite[(int)$socket];

            unset($this->waitingForWrite[(int)$socket]);

            foreach ($tasks as $task) {
                $this->schedule($task);
            }
        }
    }

    /**
     * IoPollTask function
     * 轮询任务
     *
     * @return void
     */
    protected function ioPollTask() 
    {
        while (true) {
            if ($this->taskQueue->isEmpty()) {
                $this->ioPoll(null);
            } else {
                $this->ioPoll(0);
            }

            yield;
        }
    }

}