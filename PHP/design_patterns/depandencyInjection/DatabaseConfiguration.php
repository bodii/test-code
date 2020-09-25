<?php

// 命名空间
namespace design_patterns\depandencyInjection;

/**
 * DatabaseConfiguration class
 * 数据库配置类
 */
class DatabaseConfiguration
{
    /**
     * Host variable
     * 数据库的地址
     *
     * @var string
     */
    private $host;

    /**
     * Port variable
     * 数据库端口号
     *
     * @var int
     */
    private $port;

    /**
     * Username variable
     * 数据库登录名
     *
     * @var string
     */
    private $username;

    /**
     * Password variable
     * 数据库登录密码
     *
     * @var string
     */
    private $password;

    /**
     * __construct function
     * 初始化数据库配置项方法
     *
     * @param string  $host     数据库的连接地址
     * @param integer $port     数据库的连接端口号
     * @param string  $username 数据库连接时的登录名
     * @param string  $password 数据库连接时的登录密码
     */
    public function __construct(string $host, int $port, string $username, string $password)
    {
        $this->host = $host;
        $this->port = $port;
        $this->username = $username;
        $this->password = $password;
    }

    /**
     * GetHost function
     * 获取数据库配的地址
     *
     * @return string
     */
    public function getHost(): string
    {
        return $this->host;
    }

    /**
     * GetPort function
     * 获取数据库的端口号
     *
     * @return integer
     */
    public function getPort(): int
    {
        return $this->port;
    }

    /**
     * GetUsername function
     * 获取数据库登录时的用户名
     * 
     * @return string
     */
    public function getUsername(): string
    {
        return $this->username;
    }

    /**
     * GetPassword function
     * 获取数据库登录时的密码
     *
     * @return string
     */
    public function getPassword(): string
    {
        return $this->password;
    }
}
