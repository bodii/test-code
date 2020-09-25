<?php

// 命名空间
namespace design_patterns\pool;

class StringReverseWorker
{
    /**
     * Undocumented variable
     *
     * @var \DateTime
     */
    private $_createdAt;

    public function __construct()
    {
        $this->_createdAt = new \DateTime();
    }

    public function run(string $text)
    {
        return strrev($text);
    }
}