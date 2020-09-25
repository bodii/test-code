<?php

// 命名空间
namespace design_patterns\adapter;

/**
 * EBookInterface interface
 * 电子书接口类
 */
interface EBookInterface
{
    /**
     * Unlock function
     * 解锁书本
     */
    public function unlock();

    /**
     * PressNext function
     * 按下方法
     */
    public function pressNext();

    /**
     * GetPage function
     * 获取当前页和总页数，如[10, 100]是指当前页是10，总页数是100
     *
     * @return int[]
     */
    public function getPage(): array;
}
