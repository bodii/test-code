<?php

// 命名空间：规格设计模式
namespace design_patterns\specification;

/**
 * OrSpecification class
 * 检测是否有满足一个规格类
 */
class OrSpecification implements SpecificationInterface
{
    /**
     * Specifications variable
     * 规格集合
     *
     * @var SpecificationInfterface[]
     */
    private $specifications;

    /**
     * __construct function
     * 实例化方法，传入一个规格集合
     *
     * @param SpecificationInterface ...$specifications 规格集合
     */
    public function __construct(SpecificationInterface ...$specifications)
    {
        $this->specifications = $specifications;
    }

    /**
     * IsSatisfiedBy function
     * 检测商品如果有一个规格满足则返回true
     *
     * @param Item $item 商品对象
     * 
     * @return boolean
     */
    public function isSatisfiedBy(Item $item): bool
    {
        foreach ($this->specifications as $specification) {
            if ($specification->isSatisfiedBy($item)) {
                return true;
            }
        }

        return false;
    }
}