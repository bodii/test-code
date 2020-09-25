<?php

// 命名空间
namespace design_patterns\bridge;

/**
 * HtmlFormatter class
 * html格式化类
 */
class HtmlFormatter implements FormatterInterface
{
    /**
     * Format function
     * 将字符串格式化成加p标签的html内容
     *
     * @param string $text 输入要格式化的字符串
     * 
     * @return string 返回格式化后的html标签内容
     */
    public function format(string $text): string
    {
        return sprintf('<p>%s</p>', $text);
    }
}
