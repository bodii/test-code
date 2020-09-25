<?php

// 命名空间
namespace design_patterns\adapter\test;

use design_patterns\adapter\Book;
use design_patterns\adapter\EBookAdapter;
use design_patterns\adapter\Kindle;

// 定义目录分隔符，兼容win和linux
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入composer vendor autoload 文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的phpunit测试类
use PHPUnit\Framework\TestCase;

/**
 * BookAdapterTest class
 * 书本适配器测试类
 */
class BookAdapterTest extends TestCase
{
    /**
     * TestCanTurnPageOnBook function
     * 测试书本翻页方法
     *
     * @return void
     */
    public function testCanTurnPageOnBook()
    {
        $book = new Book();
        $book->open();
        $book->turnPage();

        $this->assertEquals(2, $book->getPage());
    }

    /**
     * TestCanTurnPageOnKindleLikeInANormalBook function
     * 测试适配kindle后获取翻页的方法
     *
     * @return void
     */
    public function testCanTurnPageOnKindleLikeInANormalBook()
    {
        $kindle = new Kindle();
        $book = new EBookAdapter($kindle);

        $book->open();
        $book->turnPage();

        $this->assertEquals(2, $book->getPage());
    }

}
