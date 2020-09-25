<?php

namespace design_patterns\adapter;

class Kindle implements EBookInterface
{
    private $page = 1;

    private $totalPages = 100;

    /**
     * PressNext function
     * 按下一页方法
     *
     * @return void
     */
    public function pressNext()
    {
        $this->page++;
    }

    /**
     * Unlock function
     * 解开锁方法
     *
     * @return void
     */
    public function unlock()
    {

    }

    /**
     * GetPage function
     * 获取当前页和总页数
     * 如:[10, 100], 是当前页是第10页，总页数是100页
     *
     * @return int[]
     */
    public function getPage(): array
    {
        return [$this->page, $this->totalPages];
    }
}
