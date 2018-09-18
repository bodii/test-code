<?php

// 命名空间：PHP协同程序(协程) - 任务
namespace PHP_CoProcess;

/**
 * Task class
 * 任务类
 * 使用队列的形式将多个协程对象依次调用
 */
class Task 
{
    /**
     * TaskId variable
     * 任务id
     * 每次调用的断点都会生成一个任务id
     *
     * @var integer
     */
    protected $taskId;

    /**
     * Coroutine variable
     * 一个可迭代的生成器对象
     *
     * @var Generator
     */
    protected $coroutine;

    /**
     * SendValue variable
     * 要发送的协程断点内容
     *
     * @var mixed
     */
    protected $sendValue = null;

    /**
     * BeforeFirstYield variable
     * 标记，用于标识当前任务的第一个断点是否已经rewind
     * 一般默认当调用一个协和对象时，将会调用迭代的rewind()，并接着调用第一个yield
     *
     * @var boolean
     */
    protected $beforeFirstYield = true;

    /**
     * Exception variable
     * 抛出异常的内容
     *
     * @var Exception
     */
    protected $exception = null;

    /**
     * __construct function
     * 实例化一个迭代任务
     *
     * @param integer  $taskId    任务id, 每一个迭代对象的每一次迭代累加
     * @param Iterator $coroutine 迭代生成器对象
     */
    public function __construct($taskId, \Iterator $coroutine)
    {
        $this->taskId = $taskId;
        $this->coroutine = $coroutine;
    }

    /**
     * GetTaskId function
     * 获取当前迭代对象的任务id
     *
     * @return integer 
     */
    public function getTaskId() 
    {
        return $this->taskId;
    }

    /**
     * SetSendValue function
     * 设置要发送的内容
     *
     * @param mixed $sendValue 发送内容
     * 
     * @return void
     */
    public function setSendValue($sendValue)
    {
        $this->sendValue = $sendValue;
    }

    /**
     * SetException function
     * 设置一个异常错误
     *
     * @param Exception $exception 异常内容
     * 
     * @return void
     */
    public function setException(\Exception $exception) 
    {
        $this->exception = $exception;
    }
    /**
     * Run function
     * 执行一个任务
     *
     * @return mixed 迭代断点内容
     */
    public function run() 
    {
        if ($this->beforeFirstYield) {
            $this->beforeFirstYield = false;
            return $this->coroutine->current();
        } elseif ($this->exception) {
            $retval = $this->coroution->throw($this->exception);
            $this->excetpion = null;
            return $retval;
        } else {
            $retval = $this->coroutine->send($this->sendValue);
            $this->sendValue = null;
            return $retval;
        }
    }

    /**
     * IsFinished function
     * 当前迭代生成器对象是否已迭代完毕
     *
     * @return boolean
     */
    public function isFinished() 
    {
        return !$this->coroutine->valid();
    }
}
