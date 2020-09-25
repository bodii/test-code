<?php

// 命名空间
namespace design_patterns\adapter;

/**
 * EBookAdapter class
 * 电子书 适配器类
 */
class EBookAdapter implements BookInterface
{

    /**
     * EBook variable
     * 电子书实例对象
     * 
     * @var EBookInterface
     */
    protected $eBook;

    /**
     * __construct function
     * 实例化引入一个电子书实例
     *
     * @param EBookInterface $eBook
     */
    public function __construct(EBookInterface $eBook)
    {
        $this->eBook = $eBook;
    }

    /**
     * Open function
     * 打开电子书的方法
     *
     * @return void
     */
    public function open()
    {
        $this->eBook->unlock();
    }

    /**
     * TurnPage function
     * 翻页的方法
     *
     * @return void
     */
    public function turnPage()
    {
        $this->eBook->pressNext();
    }

    /**
     * GetPage function
     * 获取当前页的方法
     *
     * @return integer
     */
    public function getPage(): int
    {
        return $this->eBook->getPage()[0];
    }
}
