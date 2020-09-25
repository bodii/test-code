<?php

// 命名空间-资源库模式-测试
namespace design_patterns\repository\test;

use design_patterns\repository\{
MemoryStorage,
Post,
PostRepository
};

// 设定目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autload
require_once dirname(dirname(dirname(__DIR__))). DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

use PHPUnit\Framework\TestCase;

/**
 * RepositoryTest class
 * 资源库类
 */
class  RepositoryTest extends TestCase
{
    /**
     * TestCanPersistAndFindPost function
     *
     * @return void
     */
    public function testCanPersistAndFindPost()
    {
        $repository = new PostRepository(new MemoryStorage());
        $post = new Post(null, 'Repository Pattern', 'Design Patterns PHP');

        $repository->save($post);

        $this->assertEquals(1, $post->getId());
        $this->assertEquals($post->getId(), $repository->findById(1)->getId());
    }
}