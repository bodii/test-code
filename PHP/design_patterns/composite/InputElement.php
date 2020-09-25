<?php

// 命名空间
namespace design_patterns\composite;

/**
 * InputElement class
 * input 标签元素类
 */
class InputElement implements RenderableInterface
{
    /**
     * Render function
     * 显示input的方法
     *
     * @return string input标签内容
     */
    public function render(): string
    {
        return '<input type="text" />';
    }
}