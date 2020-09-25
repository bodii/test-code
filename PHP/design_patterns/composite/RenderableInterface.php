<?php

// 命名空间
namespace design_patterns\composite;

/**
 * RenderableInterface interface
 * 呈现标签接口类
 */
interface RenderableInterface
{
    /**
     * Render function
     * 显示标签内容的方法
     *
     * @return void
     */
    public function render(): string;
}

