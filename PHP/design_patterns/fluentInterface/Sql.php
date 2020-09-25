<?php
/**
 * 流接口设计模式
 * 
 * 也是连贯操作
 * 
 */

// 命名空间: 流接口设计模式
namespace design_patterns\fluentInterface;

/**
 * Sql class
 * 简单数据库模型类
 */
class Sql
{
    /**
     * Field variable
     * 要查找的字段集合
     *
     * @var array
     */
    private $fields = [];

    /**
     * From variable
     * 要查的表集合
     *
     * @var array
     */
    private $from = [];

    /**
     * Where variable
     * 查询的条件
     *
     * @var array
     */
    private $where = [];

    /**
     * Select function
     * 设置查询的字段方法
     *
     * @param array $fields 设置查询的字段
     * 
     * @return Sql 返回类自己
     */
    public function select(array $fields): Sql
    {
        $this->fields = $fields;

        return $this;
    }

    /**
     * From function
     * 设置要查的表方法
     *
     * @param string $table 表名
     * @param string $alias 表别名
     * 
     * @return Sql 返回类自已
     */
    public function from(string $table, string $alias): Sql
    {
        $this->from[] = $table.' AS '. $alias;
        
        return $this;
    }

    /**
     * Where function
     * 设置查询条件方法
     *
     * @param string $condition 查询条件
     * 
     * @return Sql 返回类自己
     */
    public function where(string $condition): Sql
    {
        $this->where[] = $condition;

        return $this;
    }

    /**
     * __toString function
     * 格式化SQL语句方法
     * 当打印这个类实例时，生成查询的SQL语句
     *
     * @return string
     */
    public function __toString(): string
    {
        return sprintf(
            'SELECT %s FROM %s WHERE %s',
            join(', ', $this->fields),
            join(', ', $this->from),
            join(' AND ', $this->where)
        );
    }
    
}
