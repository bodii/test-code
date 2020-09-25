<?php

/**
 * 结构模式
 */

// 命名空间
namespace design_patterns\adapter;

/**
 * BookInterface interface
 * 书的事件接口
 * 
 */
interface BookInterface
{
    /**
     * Open function
     * 打开书本的方法
     */
    public function open();

    /**
     * GetPage function
     * 查看页码的方法
     */
    public function getPage();

    /**
     * TurnPage function
     * 翻页的方法
     */
    public function turnPage();
}
