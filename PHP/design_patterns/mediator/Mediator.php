<?php

// 命名空间：中介者设计模式
namespace design_patterns\mediator;

/**
 * Mediator class
 * 中介者类 
 */
class Mediator implements MediatorInterface
{
    /**
     * Server variable
     * 服务器对象
     *
     * @var Subsystem\Server
     */
    private $server;

    /**
     * Database variable
     * 数据库对象
     *
     * @var Subsystem\Database
     */
    private $database;

    /**
     * Client variable
     * 连接者对象
     *
     * @var Subsystem\Client
     */
    private $client;

    /**
     * __construct function
     *
     * @param Subsystem\Database $database 数据库对象
     * @param Subsystem\Client   $client   连接者对象
     * @param Subsystem\Server   $server   服务器对象
     */
    public function __construct(Subsystem\Database $database, Subsystem\Client $client, Subsystem\Server $server)
    {
        $this->database = $database;
        $this->client = $client;
        $this->server = $server;

        $this->database->setMediator($this);
        $this->client->setMediator($this);
        $this->server->setMediator($this);

    }

    /**
     * MakeRequest function
     * 设置Request内容
     * 
     * @return void
     */
    public function makeRequest()
    {
        $this->server->process();
    }

    /**
     * QueryDb function
     * 查询数据库方法
     *
     * @return string
     */
    public function queryDb(): string
    {
        return $this->database->getData();
    }

    /**
     * SendResponse function
     * 发送Response内容
     *
     * @param mixed $content 内容
     * 
     * @return void
     */
    public function sendResponse($content)
    {
        $this->client->output($content);
    }
}