<?php

// 命名空间：规格设计模式
namespace design_patterns\specification;

/**
 * NotSpecification class
 * 不满足规格测试类
 */
class NotSpecification implements SpecificationInterface
{
    /**
     * Specification variable
     *
     * @var SpecificationInterface
     */
    private $specification;

    /**
     * __construct function
     *
     * @param SpecificationInterface $specification 规格对象
     */
    public function __construct(SpecificationInterface $specification)
    {
        $this->specification = $specification;
    }

    /**
     * IsSatisfiedBy function
     * 检测商品的规格是否不满足
     *
     * @param Item $item 商品对象
     *  
     * @return boolean
     */
    public function isSatisfiedBy(Item $item): bool
    {
        return !$this->specification->isSatisfiedBy($item);
    }
}