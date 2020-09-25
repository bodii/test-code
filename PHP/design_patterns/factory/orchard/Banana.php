<?php
/**
 * 命名空间 factory\orchard
 */
namespace factory\orchard;

/**
 * Banana class
 * 定义一个香蕉类
 */
class Banana implements IFruit
{
    public function __construct()
    {
        echo '给你摘一个香蕉.', PHP_EOL;
    }
}