<?php

// 命名空间-中介者模式-子派发类
namespace design_patterns\mediator\Subsystem;

use design_patterns\mediator\Colleague;

/**
 * Server class
 * 服务器类
 */
class Server extends Colleague
{
    /**
     * Process function
     * 处理方法
     *
     * @return void
     */
    public function process()
    {
        $data = $this->mediator->queryDb();
        $this->mediator->sendResponse(sprintf('Hello %s', $data));
    }
}
