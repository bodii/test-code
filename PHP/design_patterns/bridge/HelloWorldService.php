<?php

// 命名空间
namespace design_patterns\bridge;

/**
 * HelloWorldService class
 * Hello World的服务类
 */
class HelloWorldService extends Service
{
    /**
     * Get function
     * 继承自服务类，通过implementation定义的格式化类的实例，通过格式化方法
     *
     * @return string 返回格式化后的内容
     */
    public function get(): string
    {
        return $this->implementation->format('Hello World');
    } 
}