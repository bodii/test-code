<?php

// 命名空间：策略模式-测试
namespace design_patterns\strategy\test;

use design_patterns\strategy\{
    Context,
    DateComparator,
    IdComparator
};

// 目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

use PHPUnit\Framework\TestCase;

/**
 * StrategyTest class
 * 策略测试类
 */
class StrategyTest extends TestCase
{
    /**
     * ProvideIntegters function
     * 提供包含id索引的多维数组
     *
     * @return void
     */
    public function provideIntegers()
    {
        return [
            [
                [['id' => 2], ['id' => 1], ['id' => 3]],
                ['id' => 1],
            ],
            [
                [['id' => 3], ['id' => 2], ['id' => 1]],
                ['id' => 1],
            ],
        ];
    }

    /**
     * ProvideDates function
     * 提供包含'date'索引的多维数组
     *
     * @return void
     */
    public function provideDates()
    {
        return [
            [
                [['date' => '2018-05-07'], ['date' => '2018-02-02'], ['date' => '2018-03-07']],
                ['date' => '2018-02-02'],
            ],
            [
                [['date' => '2018-06-07'], ['date' => '2018-09-12'], ['date' => '2018-10-21']],
                ['date' => '2018-06-07'],
            ],
        ];
    }

    /**
     * TestIdComparator function
     *
     * @param mixed $collection 传入数据
     * @param mixed $expected   预期结果
     * @test
     * @dataProvider provideIntegers phpunit测试数据引入
     * 
     * @return void
     */
    public function testIdComparator($collection, $expected)
    {
        $obj = new Context(new IdComparator());
        $elements = $obj->executeStrategy($collection);

        $firstElement = array_shift($elements);
        $this->assertEquals($expected, $firstElement);
    }

    /**
     * TestDateComparator function
     *
     * @param mixed $collection 传入数据
     * @param mixed $expected   预期结果
     * @test
     * @dataProvider provideDates phpunit测试数据引入
     * 
     * @return void
     */
    public function testDateComparator($collection, $expected)
    {
        $obj = new Context(new DateComparator());
        $elements = $obj->executeStrategy($collection);

        $firstElement = array_shift($elements);

        $this->assertEquals($expected, $firstElement);
    }

}   