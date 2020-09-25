<?php
/**
 * 命名空间 factory\orchard
 */
namespace factory\orchard;

/**
 * Apple class
 * 定义一个苹果类
 */
class Apple implements IFruit
{
    public function __construct()
    {
        echo '给你摘一个苹果.', PHP_EOL;
    }
}