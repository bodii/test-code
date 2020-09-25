<?php

// 命名空间
namespace design_patterns\decorator;

/**
 * RenderInJson class
 * JSON装饰器类
 * 用于将字符串数据, 生成json格式
 */
class RenderInJson extends RendererDecorator
{
    /**
     * RenderData function
     * 渲染数据的装饰方法
     * 生成JSON格式的方法
     *
     * @return string 返回生成后的JSON内容
     */
    public function renderData(): string
    {
        return json_encode($this->wrapped->renderData());
    }
}