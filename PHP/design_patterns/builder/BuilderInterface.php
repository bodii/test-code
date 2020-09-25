<?php

namespace design_patterns\builder;

use design_patterns\builder\Parts\Vehicle;

/**
 * BuilderInterface interface
 * 生成器接口
 */
interface BuilderInterface
{
    /**
     * CreateVehicle function
     * 创建车辆方法
     *
     * @return void
     */
    public function createVehicle();

    /**
     * AddWheel function
     * 添加车辆轮子方法
     *
     * @return void
     */
    public function addWheel();

    /**
     * AddEngine function
     * 添加车辆引擎方法
     *
     * @return void
     */
    public function addEngine();

    /**
     * AddDoors function
     * 添加车门方法
     *
     * @return void
     */
    public function addDoors();

    /**
     * GetVehicle function
     * 最后获取这个车辆方法
     *
     * @return void
     */
    public function getVehicle(): Vehicle;
}