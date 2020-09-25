<?php

// 命名空间-服务定位模式
namespace design_patterns\serviceLocator;

use design_patterns\serviceLocator\ServiceLocatorInterface;


/**
 * ServiceLocator class
 * 服务定位类
 */
class ServiceLocator implements ServiceLocatorInterface
{
    /**
     * Services variable
     * 服务类的参数集合
     * 
     * @var array
     */
    private $services = [];

    /**
     * Instantiated variable
     * 实例化集合
     * 
     * @var array
     */
    private $instantiated = [];

    /**
     * Shared variable
     * 共享集合
     *
     * @var array
     */
    private $shared = [];

    /**
     * AddInstance function
     * 添加一个服务实例方法
     *
     * @param string  $class   实例名称(类名)
     * @param array   $service 服务的参数集合
     * @param boolean $share   是否共享
     * 
     * @return void
     */
    public function add(string $class, array $service = [], bool $share = true)
    {
        $this->instantiated[$class] = new $class();
        $this->services[$class] = $service;
        $this->shared[$class] = $share;
    }

    /**
     * Has function
     * 查看是否存在某个服务
     */
    public function has(string $class): bool
    {
        return isset($this->instantiated[$class]);
    }

    /**
     * Get function
     * 获取一个服务实例
     *
     * @param string $class 类名称
     * 
     * @return void
     */
    public function get(string $class)
    {
        if (isset($this->instantiated[$class]) && $this->shared[$class]) {
            return $this->instantiated[$class];
        }

        $args = $this->services[$class];

        switch(count($args)) {
        case 0:
                $object = new $class();
            break;
        case 1:
                $object = new $class($args[0]);
            break;
        case 2:
                $object = new $class($args[0], $args[1]);
            break;
        case 3:
                $object = new $class($args[0], $args[1], $args[2]);
            break;
        default:
            throw new \OutOfRangeException('Too many arguments given');
        }

        if ($this->shared[$class]) {
            $this->instantiated[$class] = $object;
        }

        return $ojbect;
    }

}