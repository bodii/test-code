<?php

// 命名空间-EAV(entity attribute value)实体属性值模型
namespace design_patterns\entityAttributeValue;

use design_patterns\entityAttributeValue\Attribute;

/**
 * ValueInterface interface
 * 值的设置与获取接口类
 */
interface ValueInterface
{
    /**
     * __construct function
     *
     * @param mixed $attribute 属性元素值
     */
    public function __construct(Attribute $attribute, string $name);

}