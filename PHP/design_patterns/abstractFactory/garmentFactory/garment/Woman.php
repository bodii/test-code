<?php 
/**
 * 命名空间 abstractFactory\garmentFactory\garment
 */
namespace abstractFactory\garmentFactory\garment;

use abstractFactory\garmentFactory\GarmentType;

class Woman implements GarmentType
{
    private static $garmentBrand = 'XXX';

    public function __construct(string $brand)
    {
        if (isset($brand)) {
            static::$garmentBrand = $brand;
        }
    }

    public function make()
    {
        echo '制衣部:', PHP_EOL;
        echo '紧急加工'.static::$garmentBrand.'牌女装中...', PHP_EOL;
    }
}