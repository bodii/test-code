<?php

// 命名空间
namespace design_patterns\composite;

/**
 * TextElement class
 * 字符串显示元素类
 */
class TextElement implements RenderableInterface
{
    /**
     * Text variable
     *
     * @var string
     */
    private $text = '';

    /**
     * __construct function
     * 构造方法
     * 实例化时传入字符串内容
     *
     * @param string $text
     */
    public function __construct(string $text) 
    {
        $this->text = $text;
    }

    /**
     * Render function
     * 显示内容方法
     *
     * @return string
     */
    public function render(): string
    {
        return $this->text;
    }
}
