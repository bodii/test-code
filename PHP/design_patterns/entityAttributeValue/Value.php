<?php

// 命名空间-实体属性值模式
namespace design_patterns\entityAttributeValue;

use design_patterns\entityAttributeValue\Attribute;

class Value implements ValueInterface
{
    private $attribute;
    private $name;

    public function __construct(Attribute $attribute, string $name)
    {
        $this->name = $name;
        $this->attribute = $attribute;

        $attribute->addValue($this);
    }

    public function __toString(): string
    {
        return sprintf('%s: %s', $this->attribute, $this->name);
    }
}