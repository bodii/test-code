<?php

// 命名空间:迭代器设计模式
namespace design_patterns\iterator;

/**
 * Book class
 * 书本类
 */
class Book
{
    /**
     * Author variable
     * 作者
     * 
     * @var string
     */
    private $author;

    /**
     * Title variable
     * 书名
     *
     * @var string
     */
    private $title;

    /**
     * __construct function
     *
     * @param string $title  书名
     * @param string $author 作者
     */
    public function __construct(string $title, string $author)
    {
        $this->title = $title;
        $this->author = $author;
    }

    /**
     * GetAuthor function
     * 获取作者信息
     *
     * @return string 返回作者信息
     */
    public function getAuthor(): string
    {
        return $this->author;
    }

    /**
     * GetTitle function
     * 获取书名信息
     * 
     * @return string 返回书名信息
     */
    public function getTitle(): string
    {
        return $this->title;
    }

    /**
     * GetAuthorAndTitle function
     * 获取作者和书名信息的字符串
     *  
     * @return string 作者和书名 
     */
    public function getAuthorAndTitle(): string
    {
        return $this->getTitle(). ' by '.$this->getAuthor();
    }
}