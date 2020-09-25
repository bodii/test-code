<?php

// 命名空间：观察者设计模式
namespace design_patterns\observer;

/**
 * UserObserver class
 * 用户观察者类
 */
class UserObserver implements \SplObserver
{
    /**
     * ChnangedUsers variable
     *
     * @var array
     */
    private $changedUsers = [];

    /**
     * Update function
     * 更新\添加要修改的用户集合
     *
     * @param \SplSubject $subject 
     * 
     * @return void
     */
    public function update(\SplSubject $subject)
    {
        $this->changedUsers[] = clone $subject;
    }

    /**
     * GetChangedUsers function
     * 获取修改的用户集合
     *
     * @return array
     */
    public function getChangedUsers(): array
    {
        return $this->changedUsers;
    }

}