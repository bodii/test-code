<?php

// 命名空间-实体属性值-测试
namespace design_patterns\entityAttributeValue\test;

use design_patterns\entityAttributeValue\Attribute;
use design_patterns\entityAttributeValue\Value;
use design_patterns\entityAttributeValue\Entity;

// 定义目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload.php
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

use PHPUnit\Framework\TestCase;

/**
 * EAVTest class
 * 实体属性值测试类
 */
class EAVTest extends TestCase
{
    /**
     * TestCanAddAttributeToEntity function
     * 测试能否成功添加属性值
     *
     * @return void
     */
    public function testCanAddAttributeToEntity()
    {
        $colorAttribute = new Attribute('color');
        $colorSilver = new Value($colorAttribute, 'silver');
        $colorBlack = new Value($colorAttribute, 'black');

        $memoryAttribute = new Attribute('memory');
        $memory8Gb = new Value($memoryAttribute, '8GB');

        $entity = new Entity(
            'MacBook Pro', 
            [$colorSilver, $colorBlack, $memory8Gb]
        );

        $this->assertEquals(
            'MacBook Pro, color: silver, color: black, memory: 8GB', 
            (string) $entity
        );
    }
}