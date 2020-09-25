<?php

namespace design_patterns\builder;

use design_patterns\builder\parts\Vehicle;

/**
 * Director class
 * 引导类
 */
class Director
{
    /**
     * Build function
     * 车辆生产模型方法
     *
     * @param BuilderInterface $builder 传入要生产的车辆类
     * 
     * @return Vehicle 返回一个生成过的车辆
     */
    public function build(BuilderInterface $builder): Vehicle
    {
        $builder->createVehicle();
        $builder->addDoors();
        $builder->addEngine();
        $builder->addWheel();

        return $builder->getVehicle();
    }
}