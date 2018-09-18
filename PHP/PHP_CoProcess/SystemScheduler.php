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
}