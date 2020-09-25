<?php

// 命名空间
namespace design_patterns\bridge;

/**
 * FormatterInterface interface
 * 格式化接口类
 */
interface FormatterInterface
{
    /**
     * Format function
     * 格式化字符串方法
     *
     * @param string $text 传入要格式化的字符串
     * 
     * @return string 返回一个格式化后的字符串
     */
    public function format(string $text): string;
}