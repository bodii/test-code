<?php

// 命名空间-实例属性值模式
namespace design_patterns\entityAttributeValue;

use design_patterns\entityAttributeValue\Value;

interface ValueAccessInterface
{
    public function __toString(): string;
}