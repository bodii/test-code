<?php
// 命名空间
namespace design_patterns\pool;

/**
 * WorkerPool class
 * 工作池类
 */
class WorkerPool implements \Countable
{
    /**
     * _occupiedWorkers variable
     * 被使用的工作组
     *
     * @var array
     */
    private $occupiedWorkers = [];

    /**
     * _freeWorkers variable
     * 未使用的工作组
     *
     * @var array
     */
    private $freeWorkers = [];


    /**
     * Count function
     * 获取工作组对象的总数量 = 使用中的工作组数量 + 未使用的工作组数量
     *
     * @return integer 工作组对象的总数量
     */
    public function count(): int
    {
        return count($this->occupiedWorkers) + count($this->freeWorkers);
    }

    /**
     * Get function
     * 获取工作对象
     *
     * @return StringReverseWorker
     */
    public function get(): StringReverseWorker
    {
        // 如果未使用工作组为空
        if (count($this->freeWorkers) == 0) {
            // 实例化创建一个工作者对象
            $worker = new StringReverseWorker();
        } else { // 否则
            // 从未使用工作组数组中取出一个工作者对象
            $worker = array_pop($this->freeWorkers);
        }

        // spl_object_hash 获取一个对象的唯一标识符
        //　并将这个工作者对象加入到已使用工作组数据中。
        $this->occupiedWorkers[spl_object_hash($worker)] = $worker;

        return $worker;
    }

    /**
     * Dispose function
     * 退还工作者对象
     *
     * @param StringReverseWorker $worker 工作者对象
     * 
     * @return void
     */
    public function dispose(StringReverseWorker $worker)
    {
        $key = spl_object_hash($worker);

        // 如果这个工作者对象存在于已使用工作组中
        if (isset($this->occupiedWorkers[$key])) {
            // 将已使用工作组中的这个工作者对象销毁
            unset($this->occupiedWorkers[$key]);
            // 并将这个工作者添加到未使用工作组中
            $this->freeWorkers[$key] = $worker;
        }
    }

}
