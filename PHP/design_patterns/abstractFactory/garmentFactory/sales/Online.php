<?php
/**
 * 命名空间 abstractFactory\garmentFactory\sales
 */
namespace abstractFactory\garmentFactory\sales;

use \abstractFactory\garmentFactory\SalesPatterns;

class Online implements SalesPatterns
{
    private static $garmentBrand = 'XXX';
    private static $storeName = 'XXX';

    public function __construct(string $brand, $storeName='')
    {
        if (isset($brand)) {
            static::$garmentBrand = $brand;
        }

        if (!empty($storeName) && is_string($storeName)) {
            static::$storeName = $storeName;
        }

        echo '销售部:', PHP_EOL;
        echo static::$garmentBrand.'牌服装厂线上销售部建立完成.', PHP_EOL;
    }

    public function sell()
    {
        echo '销售部:', PHP_EOL;
        echo static::$garmentBrand.'牌服装在[ '.static::$storeName.' ]线上店开始销售了', PHP_EOL;
    } 
}