<?php

// 命名空间：迭代器设计模式
namespace design_patterns\iterator;

class BookList implements \Countable, \Iterator
{
    /**
     * Books variable
     * 书本容器
     *
     * @var array 
     */
    private $books = [];

    /**
     * CurrentIndex variable
     *
     * @var integer 当前书本的索引
     */
    private $currentIndex = 0;

    /**
     * AddBook function
     * 添加一个书本对象到书本容器中
     *
     * @param Book $book 一个书本对象
     * 
     * @return void
     */
    public function addBook(Book $book)
    {
        $this->books[] = $book;
    }

    /**
     * RemoveBook function
     * 将一个书本对象从书本容器中删除
     *
     * @param Book $bookToRemove 要删除的书本对象
     * 
     * @return void
     */
    public function removeBook(Book $bookToRemove)
    {
        foreach ($this->books as $key => $book) {
            if ($book->getAuthorAndTitle() === $bookToRemove->getAuthorAndTitle()) {
                unset($this->books[$key]);
            }
        }

        // 重置一下数组的索引
        $this->books = array_values($this->books);
    }

    /**
     * Count function
     * 获取书本容器的收本对象的总个数
     *
     * @return integer 
     */
    public function count(): int
    {
        return count($this->books);
    }

    /**
     * Current function
     * 获取当前索引下的书本对象
     *
     * @return Book
     */
    public function current(): Book
    {
        return $this->books[$this->currentIndex];
    }

    /**
     * Key function
     * 获取当前书本对象的索引
     *
     * @return integer
     */
    public function key(): int
    {
        return $this->currentIndex;
    }

    /**
     * Next function
     * 获取下一个书本对象的索引
     *
     * @return integer
     */
    public function next(): int
    {
        return $this->currentIndex++;
    }

    /**
     * Rewind function
     * 将当前的书本的索引指针归零
     *
     * @return void
     */
    public function rewind(): void
    {
        $this->currentIndex = 0;
    }

    /**
     * Valid function
     * 返回当前索引的书本对象是否设置
     * 或是否是已经迭代完书本容器
     *
     * @return boolean
     */
    public function valid(): bool
    {
        return isset($this->books[$this->currentIndex]);
    }
}