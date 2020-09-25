<?php

// 命名空间
namespace design_patterns\bridge;

/**
 * PlainTextFormatter class
 * 字符串格式化类
 */
class PlainTextFormatter implements FormatterInterface
{
    /**
     * Format function
     * 格式化字符串方法
     *
     * @param string $text 传入要格式化的字符串
     * 
     * @return string 返回格式化后的字符串
     */
    public function format(string $text): string
    {
        return $text;
    }
}
