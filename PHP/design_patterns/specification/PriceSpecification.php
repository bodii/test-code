<?php

// 命名空间：规格设计模式
namespace design_patterns\specification;

/**
 * PriceSpecification class
 * 格式规格类
 */
class PriceSpecification implements SpecificationInterface
{
    /**
     * MaxPrice variable
     *
     * @var float|integer|string
     */
    private $maxPrice;

    /**
     * MinPrice variable
     *
     * @var float|integer|string
     */
    private $minPrice;

    /**
     * __construct function
     *
     * @param float|integer|string $minPrice 是低价格
     * @param float|integer|string $maxPrice 是高价格
     */
    public function __construct($minPrice, $maxPrice)
    {
        $this->minPrice = $minPrice;
        $this->maxPrice = $maxPrice;
    }

    /**
     * IsSatisfiedBy function
     *
     * @param Item $item 商品
     * 
     * @return boolean
     */
    public function isSatisfiedBy(Item $item): bool
    {
        if ($this->maxPrice !== null && $item->getPrice() > $this->maxPrice) {
            return false;
        }
        
        if ($this->minPrice !== null && $item->getPrice() < $this->minPrice) {
            return false;
        }

        return true;
    }
}