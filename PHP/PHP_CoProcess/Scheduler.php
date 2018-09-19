<?php

// 命名空间：PHP协程 - 调度
namespace PHP_CoProcess;

use PHP_CoProcess\Task;

/**
 * Scheduler class
 * 任务调度器类
 */
class Scheduler
{
    /**
     * MaxTaskId variable
     * 一个自增的任务id
     *
     * @var integer
     */
    protected $maxTaskId = 0;

    /**
     * TaskMap variable
     * 任务容器
     *
     * @var array
     */
    protected $taskMap = [];

    /**
     * TaskQueue variable
     * 任务调用队列
     *
     * @var SplQueue
     */
    protected $taskQueue;

    /**
     * __construct function
     * 实例化一个任务队列调试
     */
    public function __construct()
    {
        $this->taskQueue = new \SplQueue();
    }

    /**
     * NewTask function
     * 添加一个新的任务
     *
     * @param Generator $coroutine 迭代生成器对象
     * 
     * @return integer 任务id 
     */
    public function newTask(\Generator $coroutine)
    {
        $taskId = ++$this->maxTaskId;
        $task = new Task($taskId, $coroutine);
        $this->taskMap[$taskId] = $task;
        $this->schedule($task);

        return $taskId;
    }

    /**
     * Schedule function
     * 任务调度
     *
     * @param Task $task 将任务添加到队列
     * 
     * @return void
     */
    public function schedule(Task $task)
    {
        $this->taskQueue->enqueue($task);
    }

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
            $task->run();

            if ($task->isFinished()) { // 如果当前迭代对象完成
                unset($this->taskMap[$task->getTaskId()]); // 销毁这个任务
            } else {
                $this->schedule($task); // 否则将迭代断点(yield)添加到任务调用度
            }
        }
    }
}