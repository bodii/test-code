<?php 

/*
	【 迭代器模式 】
    在不需要了解内部实现的前提下，遍历一个聚合对狗日的内部元素而又不暴露该对象
    的内部表示。

    适用场景
    1。 访问一个聚合对象的内容而无需暴露它的内部表示
    2。支持对聚合对象的多种遍历
    3。为遍历不同的聚合结构提供一个统一的接口
*/

/*
    •Iterator::current — 返回当前元素
    •Iterator::key — 返回当前元素的键
    •Iterator::next — 向前移动到下一个元素
    •Iterator::rewind — 返回到迭代器的第一个元素
    •Iterator::valid — 检查当前位置是否有效

*/

class ConcreteIterator implements Iterator // Iterator是PHP预定义的迭代接口
{
    public $arr = [];
    private $position = 0;

    public function __construct($arrNow)
    {
        $this->arr = $arrNow;
    }

    public function rewind()
    {
        var_dump(__METHOD__);
        $this->position = 0;
    }

    public function current()
    {
        var_dump(__METHOD__);
        return $this->arr[$this->position];
    }

    public function key()
    {
        var_dump(__METHOD__);
        return $this->position;
    }

    public function next()
    {
        var_dump(__METHOD__);
        ++ $this->position;
    }

    public function valid()
    {
        var_dump(__METHOD__);
        return isset($this->arr[$this->position]);
    }
}

$arr = ['xiao hong', 'xiao ming', 'xiao hua'];
$concreteIterator = new ConcreteIterator($arr);
foreach($concreteIterator as $key => $value)
{
    echo  "$key=>$value", '<br>';
}