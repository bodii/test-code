<?php

// 命名空间：规格设计模式
namespace design_patterns\specification;

/**
 * Item class
 * 商品类
 */
class Item
{
    /**
     * Price variable
     *
     * @var float
     */
    private $price;

    /**
     * __construct function
     *
     * @param float $price 价格
     */
    public function __construct(float $price)
    {
        $this->price = $price;
    }

    /**
     * GetPrice function
     * 获取价格
     *
     * @return float 价格
     */
    public function getPrice(): float
    {
        return $this->price;
    }
}