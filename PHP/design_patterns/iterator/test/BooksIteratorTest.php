<?php

// 命名空间：迭代设计模式-测试
namespace design_patterns\iterator\test;

use design_patterns\iterator\Book;
use design_patterns\iterator\BookList;

// 设置目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autolaod文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用vendor下的PHPUnit单元测试类
use PHPUnit\Framework\TestCase;

/**
 * BooksIteratorTest class
 * 书
 */
class BooksIteratorTest extends TestCase
{
    /**
     * TestCanIterateOverBookList function
     * 测试是否能添加书本对象，并迭代调用书本的相关信息
     *
     * @return void
     */
    public function testCanIterateOverBookList()
    {
        $bookList = new BookList();
        $bookList->addBook(
            new Book('Learning PHP Design Patterns', 'William Sanders')
        );
        $bookList->addBook(
            new Book('Professional Php Design Patterns', 'Aaron Saray')
        );
        $bookList->addBook(
            new Book('Clean Code', 'Robert C. Martin')
        );

        $books = [];

        foreach ($bookList as $book) {
            $books[] = $book->getAuthorAndTitle();
        }

        $this->assertEquals(
            [
                'Learning PHP Design Patterns by William Sanders',
                'Professional Php Design Patterns by Aaron Saray',
                'Clean Code by Robert C. Martin',
            ],
            $books
        );

    }

    /**
     * TsetCanIterateOverBookListAfterRemoveingBook function
     * 测试是否能在迭代前中删除书本对象的方法
     *
     * @return void
     */
    public function tsetCanIterateOverBookListAfterRemoveingBook()
    {
        $book1 = new Book('Clean Code', 'Robert C. Martin');
        $book2 = new Book('Professional Php Design Patterns', 'Aaron Saray');

        $bookList = new BookList();
        $bookList->addBook($book1);
        $bookList->addBook($book2);
        $bookList->removeBook($book1);

        $books = [];
        foreach ($bookList as $book) {
            $books[] = $book->getAuthorAndTitle();
        }

        $this->assertEquals(
            ['Professional Php Design Patterns by Aaron Saray'],
            $books
        );
    }

    /**
     * TestCanAddBookToListAndBookCounts function
     * 测试是否能添加书本对象到书本列表
     *
     * @return void
     */
    public function testCanAddBookToListAndBookCounts()
    {
        $book = new Book('Clean Code', 'Robert C. Martin');

        $bookList = new BookList();
        $bookList->addBook($book);

        $this->assertCount(1, $bookList);
    }

    /**
     * TestCanRemoveBookFromList function
     * 测试能否从列表中添加并删除书本对象
     *
     * @return void
     */
    public function testCanRemoveBookFromList()
    {
        $book = new Book('Clean Code', 'Robert C. Martin');

        $bookList = new BookList();
        $bookList->addBook($book);
        $bookList->removeBook($book);

        $this->assertCount(0, $bookList);
    }
}


