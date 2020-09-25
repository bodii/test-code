<?php

// 命名空间 享元接口
namespace design_patterns\flyweight;

/**
 * FlyweightInterface interface
 * 享元接口类
 */
interface FlyweightInterface
{
    /**
     * Render function
     *
     * @param string $extrinsicState 传入字符串内容
     * 
     * @return string 返回渲染后的字符串
     */
    public function render(string $extrinsicState): string;
}