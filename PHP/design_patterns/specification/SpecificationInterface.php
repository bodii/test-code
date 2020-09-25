<?php

// 命名空间：规格模式
namespace design_patterns\specification;

/**
 * SpecificationInterface interface
 * 规格模式接口类
 */
interface SpecificationInterface
{
    public function isSatisfiedBy(Item $item): bool;

}