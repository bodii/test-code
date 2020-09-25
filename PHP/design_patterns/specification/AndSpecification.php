<?php

// 命名空间：规格设计模式
namespace design_patterns\specification;

/**
 * AndSpecification class
 * 检测是否满足所有规格类
 */
class AndSpecification implements SpecificationInterface
{
    /**
     * Specification variable
     * 规格集合
     *
     * @var SpecificationInterface[]
     */
    private $specifications;

    public function __construct(SpecificationInterface ...$specifications)
    {
        $this->specifications = $specifications;
    }

    /**
     * IsSatisfiedBy function
     * 检测商品如果有一个规格不满足，则返回false
     *
     * @param Item $item 商品对象
     * 
     * @return boolean
     */
    public function isSatisfiedBy(Item $item): bool
    {
        foreach ($this->specifications as $specification) {
            if (!$specification->isSatisfiedBy($item)) {
                return false;
            }
        }

        return true;
    }
}