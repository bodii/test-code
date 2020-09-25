<?php

// 命名空间:中介者-子派发
namespace design_patterns\mediator\Subsystem;

use design_patterns\mediator\Colleague;

/**
 * Client class
 * 客户端连接类
 */
class Client extends Colleague
{
    /**
     * Request function
     * Request方法
     *
     * @return void
     */
    public function request()
    {
        $this->mediator->makeRequest();
    }

    /**
     * Output function
     * 输出内容
     *
     * @param string $content 内容
     * 
     * @return void
     */
    public function output(string $content)
    {
        echo $content;
    }
}
