<?php 

// 命名空间-中间者模式
namespace design_patterns\mediator;

/**
 * Colleague class
 * 关联协同类
 */
abstract class Colleague
{
    protected $mediator;

    /**
     * SetMediator function
     * 设置中介者对象
     *
     * @param MediatorInterface $mediator 中介者对象
     * 
     * @return void
     */
    public function setMediator(MediatorInterface $mediator)
    {
        $this->mediator = $mediator;
    }
}