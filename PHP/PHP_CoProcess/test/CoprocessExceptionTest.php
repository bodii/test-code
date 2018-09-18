<?php

// 命名空间：PHP协程-测试
namespace PHP_CoProcess\test;

/**
 * Gen function
 */
function gen()
{
    echo "Foo\n";
    try {
        yield;
    } catch (\Exception $e) {
        echo "Excetpion: {$e->getMessage()}\n";
    }

    echo "Bar\n";
}

$gen = gen();
$gen->rewind();
$gen->throw(new \Exception('Test'));
