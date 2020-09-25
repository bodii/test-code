<?php

// 命名空间: 享元设计模式测试
namespace design_patterns\flyweight\test;

use design_patterns\flyweight\FlyweightFactory;

// 设置目录分隔符常量
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入 composer vendor autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的PHPUnit测试类
use PHPUnit\Framework\TestCase;

/**
 * FlyweightTest class
 * 享元测试类
 */
class FlyweightTest extends TestCase
{

    /**
     * Characters variable
     * 字符名称集合
     *
     * @var array
     */
    private $characters = [
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
        'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 
        'w', 'x', 'y', 'z'
    ];
    /**
     * Fonts variable
     * 字体名称集合
     *
     * @var array
     */
    private $fonts = ['Arial', 'Times New Roman', 'Verdana', 'Helvetica'];

    /**
     * TestFlyweight function
     * 享元测试方法
     *
     * @return void
     */
    public function testFlyweight()
    {
        $factory = new FlyweightFactory();

        foreach ($this->characters as $char) {
            foreach ($this->fonts as $font) {
                $flyweight = $factory->get($char);
                $rendered = $flyweight->render($font);

                $this->assertEquals(
                    sprintf(
                        'Character %s with font %s',
                        $char, $font
                    ),
                    $rendered
                );
            }
        }
        
        // assertCount(统计的元素个数, 一个数组), 在这里factory是FlyweightFactory的实例
        // 而FlyweightFactory 实例自Countable抽象类,这个抽象类的实现者就是一个可被count调用的
        // 当使用count公共函数统计这个对象时，会自动触发自身的count方法
        $this->assertCount(count($this->characters), $factory);
    }
}   
