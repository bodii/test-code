<?php

// 命名空间-资源库模式
namespace design_patterns\repository;

use design_patterns\repository\MemoryStorage;

class PostRepository
{
    /**
     * Persistence variable
     *
     * @var MemoryStorage
     */
    private $persistence;

    /**
     * __construct function
     *
     * @param MemoryStorage $persistence 
     */
    public function __construct(MemoryStorage $persistence)
    {
        $this->persistence = $persistence;
    }

    /**
     * FindById function
     *
     * @param integer $id id
     * 
     * @return Post
     */
    public function findById(int $id): Post
    {
        $arrayData = $this->persistence->retrieve($id);

        if (is_null($arrayData)) {
            throw new \InvalIdArgumentException(sprintf('Post with ID %d does not exist', $id));
        }

        return Post::fromState($arrayData);
    }

    /**
     * Save function
     *
     * @param Post $post 
     * 
     * @return void
     */
    public function save(Post $post)
    {
        $id = $this->persistence->persist(
            [
                'text' => $post->getText(),
                'title' => $post->getTitle(),
            ]
        );

        $post->setId($id);
    }
}