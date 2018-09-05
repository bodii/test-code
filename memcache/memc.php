<?php
/**
 * Memc 文件
 * 
 */
namespace memc;

require_once __DIR__ . 'interface/hasher.php';
require_once __DIR__ . 'interface/distribution.php';

use hasher;
use distribution;

/**
 * Memc class
 * memcache 一致性哈希使用
 */
class Memc implements \hasher\IHasher, \distribution\IDistribution
{
    // 节点
    protected $nodes = [];
    // 虚拟槽
    protected $slots = [];
    // 节点生成虚拟槽的倍数
    protected $nul = 64;

    /**
     * MakeHash function
     * 把字符串生成32位的数字
     *
     * @param string $str 一个字符串
     * 
     * @return int 32位的数字
     */
    public function makeHash($str = '')
    {
        // 把字符串转成32位的数字
        return sprintf('%u', crc32($str));
    }

    /**
     * Lookup function
     * 查找最近的虚拟槽并返回
     * 
     * @param string|int $key 要添加到memcache内的字符串值的key
     * 
     * @return int 一个能添加的虚拟槽
     */
    public function lookup($key)
    {
        $val_slot = $this->makeHash($key);

        $node = key($this->slots);

        foreach ($this->slots as $slot => $nd) {
            if ($val_slot <= $slot) {
                $node = $nd;
                break;
            }
        }

        reset($this->slots);
        return $node;
    }

    /**
     * AddNode function
     * 添加节点
     * 
     * @param string|int $node 节点名称
     * 
     * @return void
     */
    public function addNode($node='')
    {
        if (isset($this->nodes[$node])) {
            return ;
        }

        for ($i=0; $i < $this->nul; $i++) {
            $new_slot = $this->makeHash($node . '_' . $i);
            $this->slots[$new_slot] = $node;
            $this->nodes[$node][] = $new_slot;
        }

        $this->sortSlots();
    }

    /**
     * DelNode function
     * 删除节点，包括虚拟节点
     *
     * @param string $node 节点名称
     * 
     * @return void
     */
    public function delNode($node='')
    {
        if (!isset($this->nodes[$node])) {
            return ;
        }

        foreach ($this->nodes[$node] as $slot) {
            unset($this->slots[$slot]);
        }

        unset($this->nodes[$node]);
    }

    /**
     * SortPos function
     * 将现有虚拟节点排序
     * 
     * @return void
     */
    public function sortSlots()
    {
        ksort($this->slots);
    }
}