<?php
/**
 * IHasher散列器命名空间 
 * @example hasher.php require_once hasher.php
 */
namespace hasher;

/**
 * IHasher interface
 * 散列器接口
 * 
 * @category Interface
 * @package  IHasher
 * @author   wong <wong@mail.com>
 * @license  MIT https://github.com
 * @link     https://github.com
 */
interface IHasher
{
    /**
     * MakeHash function
     * 创建hash
     *
     * @param string $str 字符串
     * 
     * @return int 返回数字
     */
    public function makeHash($str='');
}