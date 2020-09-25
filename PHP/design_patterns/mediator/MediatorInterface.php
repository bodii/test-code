<?php

// 命名空间：中介者模式
namespace design_patterns\mediator;

/**
 * MediatorInterface interface
 */
interface MediatorInterface
{
    /**
     * SendResponse function
     * 发送Resonse内容的方法
     *
     * @param mxied $content 发送的内容
     * 
     * @return void
     */
    public function sendResponse($content);

    /**
     * MakeRequest function
     * 生成Request
     *
     * @return void
     */
    public function makeRequest();

    /**
     * QueryDb function
     * 查询数据库
     *
     * @return void
     */
    public function queryDb();
}