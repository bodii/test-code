<?php

// 命名空间
namespace design_patterns\dataMapper;

/**
 * User class
 * 用户类
 */
class User
{
    /**
     * Email variable
     * 用户邮箱
     *
     * @var string
     */
    private $email = '';

    /**
     * Username variable
     * 用户名
     *
     * @var string
     */
    private $username = '';

    /**
     * __construct function
     * 类初始化方法
     * 用于设置用户的邮箱和名子
     *
     * @param string $username 用户名
     * @param string $email    邮箱
     */
    public function __construct(string $username, string $email)
    {
        $this->email = $email;
        $this->username = $username;
    }

    /**
     * GetUsername function
     * 获取用户名方法
     *
     * @return void
     */
    public function getUsername()
    {
        return $this->username;
    }

    /**
     * GetEmail function
     * 获取邮箱方法
     *
     * @return void
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * FromState function
     * 用户名和邮箱实例方法
     *
     * @param array $state [username, email]
     * 
     * @return User 返回一个用户类实例
     */
    public static function fromState(array $state): User
    {
        return new self(
            $state['username'],
            $state['email']
        );
    }

}