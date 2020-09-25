<?php

// 命名空间-实例属性值模式
namespace design_patterns\entityAttributeValue;

use design_patterns\entityAttributeValue\ValueAccessInterface;
use design_patterns\entityAttributeValue\Value;

class Attribute implements ValueAccessInterface
{
    private $values;

    private $name;

    public function __construct(string $name)
    {
        $this->values = new \SplObjectStorage();
        $this->name = $name;
    }

    public function getValues(): \SplObjectStorage
    {
        return $this->values;
    }

    public function addValue(Value $value)
    {
        $this->values->attach($value);
    }

    public function removeValue($value)
    {
        $this->values->detach($value);
    }

    public function __toString(): string
    {
        return $this->name;
    }

    public function getName()
    {
        return $this->name;
    }

    public function setName($name)
    {
        $this->name = $name;
    }
}