<?php

// 命名空间：观察者模式
namespace design_patterns\observer;

/**
 * User class
 * 用户类
 * 用于生成一个被观察实例
 */
class User implements \SplSubject
{
    /**
     * Email variable
     *
     * @var string
     */
    private $email;

    /**
     * Observers variable
     *
     * @var \SplObjectStorage
     */
    private $observers;

    /**
     * __construct function
     * 实例方法，用于设置一个观察者对象容器
     */
    public function __construct()
    {
        $this->observers = new \SplObjectStorage();
    }

    /**
     * Attach function
     * 添加
     *
     * @param \SplObserver $observer 观察者对象
     * 
     * @return void
     */
    public function attach(\SplObserver $observer)
    {
        $this->observers->attach($observer);
    }

    /**
     * Detach function
     * 删除
     *
     * @param \SplObserver $observer 观察者对象
     * 
     * @return void
     */
    public function detach(\SplObserver $observer)
    {
        $this->observers->detach($observer);
    }

    /**
     * ChangeEmail function
     *
     * @param string $email email地址
     * 
     * @return void
     */
    public function changeEmail(string $email)
    {
        $this->email = $email;
        $this->notify();
    }

    /**
     * Notify function
     * 通知观察者
     *
     * @return void
     */
    public function notify()
    {
        foreach ($this->observers as $observer) {
            $observer->update($this);
        }
    }
}