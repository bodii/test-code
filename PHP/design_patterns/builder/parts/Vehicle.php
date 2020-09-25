<?php

namespace design_patterns\builder\parts;

/**
 * Vehicle class
 * 车辆抽象类
 */
abstract class Vehicle
{
    /**
     * @var array
     */
    private $_data = [];

    /**
     * SetPart function
     * 添加车辆配件方法
     *
     * @param string $key   配件名称
     * @param object $value 配件实物对象
     * 
     * @return void
     */
    public function setPart($key, $value)
    {
        $this->_data[$key] = $value;
    }
}