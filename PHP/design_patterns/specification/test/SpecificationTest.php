<?php

// 命名空间：规格设计模式-测试
namespace design_patterns\specification\test;

use design_patterns\specification\{
    Item,
    OrSpecification,
    AndSpecification,
    NotSpecification,
    PriceSpecification
};

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用单元测试类
use PHPUnit\Framework\TestCase;

/**
 * SpecificationTest class
 * 规格测试类
 */
class SpecificationTest extends TestCase
{
    /**
     * TestCanOr function
     * 测试是否满足多个规格的其中一个规格
     *
     * @return void
     */
    public function testCanOr()
    {
        $spec1 = new PriceSpecification(50, 99);
        $spec2 = new PriceSpecification(101, 200);

        $orSpec = new OrSpecification($spec1, $spec2);

        $this->assertFalse($orSpec->isSatisfiedBy(new Item(100))); // 第一、第二规格都不满足
        $this->assertTrue($orSpec->isSatisfiedBy(new Item(51))); // 满足第一个规格
        $this->assertTrue($orSpec->isSatisfiedBy(new Item(150))); // 满足第三个规格
    }

    /**
     * TestCanAnd function
     * 测试是否所有规格
     *
     * @return void
     */
    public function testCanAnd()
    {
        $spec1 = new PriceSpecification(50, 100);
        $spec2 = new PriceSpecification(80, 200);

        $andSpec = new AndSpecification($spec1, $spec2);

        $this->assertFalse($andSpec->isSatisfiedBy(new Item(150))); // 只满足了第二个规格，false
        $this->assertFalse($andSpec->isSatisfiedBy(new Item(1))); // 两个都不满足，false
        $this->assertFalse($andSpec->isSatisfiedBy(new Item(51))); // 只满足了第一个规格, false
        $this->assertTrue($andSpec->isSatisfiedBy(new Item(100))); // 两个都满足,true
    }

    public function testCanNot()
    {
        $spec1 = new PriceSpecification(50, 100);

        $notSpec = new NotSpecification($spec1);

        $this->assertTrue($notSpec->isSatisfiedBy(new Item(150))); // 不满足规格，返回true
        $this->assertFalse($notSpec->isSatisfiedBy(new Item(50))); // 满足规格(和最小金额相同)，返回false
    }
}