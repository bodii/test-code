<?php

if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

if (!defined('PATTERNS_PATH')) {
    define('PATTERNS_PATH', dirname(__DIR__).DIR_SEP);
}

// 匿名函数注册自动加载
spl_autoload_register(
    function ($className) {
        require_once PATTERNS_PATH.str_replace('\\', DIR_SEP, $className).'.php';
    }
);

use factory\farm\Farm;
use factory\SampleFactory;
use factory\orchard\Orchard;

echo '工厂方法(农场):', PHP_EOL;
$farm = new Farm;
$farm->produce('chicken');
$farm->produce('pig');
echo PHP_EOL;

echo '简单工厂，无需要实例化：', PHP_EOL;
SampleFactory::produce('chicken');
SampleFactory::produce('pig');
echo PHP_EOL;

echo '工厂方法(果园):', PHP_EOL;
$orchard = new Orchard;
$orchard->produce('apple');
$orchard->produce('pear');
echo PHP_EOL;