<?php

// 命名空间：PHP协程-测试
namespace PHP_CoProcess\test;

/**
 * Gen function
 *
 * @return void
 */
function gen()
{
    yield 'foo';
    yield 'bar';
}

// 是否是一个迭代生成器
var_dump((gen() instanceof \Generator));

$gen = gen(); // rewind执行，消耗掉一个yield
var_dump($gen->send('something')); // output:bar
