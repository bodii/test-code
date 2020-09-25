<?php

// 命名空间：命令行设计模式
namespace design_patterns\command;

/**
 * Receiver class
 * 接收者类
 */
class Receiver
{
    /**
     * EnableDate variable
     * 是否设置时间
     *
     * @var boolean
     */
    private $enableDate = false;

    /**
     * Output variable
     * 设置输出内容的数组
     *
     * @var array
     */
    private $output = [];

    /**
     * Write function
     *
     * @param string $str 传入要输出的字符串内容
     * 
     * @return void
     */
    public function write(string $str)
    {
        if ($this->enableDate) {
            $str .= ' ['.date('Y-m-d').']';
        }

        $this->output[] = $str;
    }

    /**
     * GetOutput function
     * 获取输出内容
     *
     * @return string
     */
    public function getOutput(): string
    {
        return join("\n", $this->output);
    }

    /**
     * EnableDate function
     * 设置开启输出时添加日期内容
     *
     * @return void
     */
    public function enableDate()
    {
        $this->enableDate = true;
    }

    /**
     * DisableDate function
     * 设置关闭输出时添加日期内容
     *
     * @return void
     */
    public function disableDate()
    {
        $this->enableDate = false;
    }
}
