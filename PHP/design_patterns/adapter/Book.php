<?php

// 命名空间
namespace design_patterns\adapter;

/**
 * Book class
 * 书本类
 */
class Book implements BookInterface
{
    /**
     * page variable
     * 页码
     * 
     * @var int
     */
    private $page;

    /**
     * Opne function
     * 打开书本的方法
     *
     * @return void
     */
    public function open()
    {
        $this->page = 1;
    }

    /**
     * TurnPage function
     * 书本翻页的方法
     *
     * @return void
     */
    public function turnPage()
    {
        $this->page ++;
    }

    /**
     * Getpage function
     * 获取书本页码的方法
     *
     * @return void
     */
    public function getPage()
    {
        return $this->page;
    }
}
