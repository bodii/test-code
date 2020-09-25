<?php

// 命名空间
namespace design_patterns\composite;

/**
 * Form class
 * 表单类
 */
class Form implements RenderableInterface
{
    /**
     * Elementes variable
     * 标签集合数组
     *
     * @var object[] 
     */
    private $elements = [];

    /**
     * Render function
     * 显示标签的方法
     *
     * @return string
     */
    public function render(): string
    {
        $formCode = '<form>';

        foreach ($this->elements as $element) {
            $formCode .= $element->render();
        }

        $formCode .= '</form>';

        return $formCode;
    }

    /**
     * AddElement function
     * 添加元素实例到元素标签集合数组
     *
     * @param RenderableInterface $element 元素实例
     * 
     * @return void
     */
    public function addElement(RenderableInterface $element)
    {
        $this->elements[] = $element;
    }
}