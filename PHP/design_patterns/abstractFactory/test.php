<?php
// 定义目录分隔符常量:Win='\', Linux='/'
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 定义patterns的目录绝对路径常量
if (!defined('PATTERNS')) {
    define('PATTERNS', dirname(__DIR__).DIR_SEP);
}

// 匿名函数：注册自动加载
spl_autoload_register(
    function ($className) {
        require_once PATTERNS.str_replace('\\', DIR_SEP, $className).'.php';
    }
);

use abstractFactory\garmentFactory\GarmentFactory;
use abstractFactory\garmentFactory\garment as garment;
use abstractFactory\garmentFactory\sales as sales;

//--------------------测试-----------------------------

// 实例化一个服装厂, 品牌“美丽衣橱”
$garmentFactory = new GarmentFactory('美丽衣橱');
// 获取品牌名称
$brand = $garmentFactory->brand;

// 让服装厂制衣部开始生产男装
$garmentFactory::production(new garment\Man($brand));
echo PHP_EOL;

// 让服装厂制衣部开始生产女装
$garmentFactory::production(new garment\Woman($brand));
echo PHP_EOL;

// 开设服装厂线上销售渠道，并上线销售
$storeName = $brand.'天猫旗舰店';
$garmentFactory::sales(new sales\Online($brand, $storeName));
echo PHP_EOL;

// 开设服装厂线下品牌实体店，并开始销售
$storeName = $brand.'品牌连锁';
$garmentFactory::sales(new sales\Store($brand, $storeName));
echo PHP_EOL;


