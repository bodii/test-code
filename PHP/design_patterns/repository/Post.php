<?php

// 命名空间-资源库模式
namespace design_patterns\repository;

use phpDocumentor\Reflection\Types\Void_;


class Post
{
    /**
     * Id variable
     *
     * @var int|null
     */
    private $id;

    /**
     * Title variable
     *
     * @var string
     */
    private $title;

    /**
     * Text variable
     * 
     * @var string
     */
    private $text;

    /**
     * FromState function
     * 将一个传入包含类属性的数组，并实例化该类
     *
     * @param array $state 参数值
     * 
     * @return Post
     */
    public static function fromState(array $state): Post
    {
        return new self(
            $state['id'],
            $state['title'],
            $state['text']
        );
    }

    /**
     * __Construct function
     * 类实例化方法
     *
     * @param int|null $id 
     * @param string   $title title
     * @param string   $text  text
     */
    public function __construct($id , string $title, string $text)
    {
        $this->id = $id;
        $this->title = $title;
        $this->text = $text;
    }

    /**
     * GetId function
     *
     * @return integer id
     */
    public function getId(): int
    {
        return $this->id;
    }

    public function setId(int $id): void
    {
        $this->id = $id;
    }

    /**
     * GetText function
     * 获取text内容
     *
     * @return string
     */
    public function getText(): string
    {
        return $this->text;
    }

    /**
     * GetTitle function
     * 获取title
     *
     * @return string
     */
    public function getTitle(): string
    {
        return $this->title;
    }
}