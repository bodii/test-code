<?php

// 命名空间
namespace design_patterns\decorator;

/**
 * RendererInterface interface
 * 渲染接口
 */
interface RendererInterface
{
    /**
     * RenderData function
     * 渲染数据的方法
     * 
     * @return string 返回渲染后的字符串
     */
    public function renderData(): string;
}