<?php

// 定义目录分隔符常量:Win='\', Linux='/'
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 定义design_patterns目录常量
if (!defined('PATTERNS_PATH')) {
    define('PATTERNS_PATH', dirname(__DIR__).DIR_SEP);
}

// 匿名函数：自动注册加载
spl_autoload_register(
    function ($className) {
        require_once PATTERNS_PATH.str_replace('\\', DIR_SEP, $className).'.php';
    }
);

use prototype\Prototype;

// 
$proto = new Prototype();

// 获取一个原型的clone
$protoCloneOne = $proto->copy();
$protoCloneOne->name = 'one';
echo $protoCloneOne->name, PHP_EOL;

$protoCloneTow = $proto->copy();
$protoCloneTow->name = 'two';
echo $protoCloneTow->name, PHP_EOL;

echo $protoCloneOne->name, PHP_EOL;