<?php

namespace design_patterns\builder;

use design_patterns\builder\parts\Vehicle;

/**
 * CarBuilder class
 * 汽车生成器
 */
class CarBuilder implements BuilderInterface
{
    /**
     * @var parts\Car
     */
    private $_car;

    public function addDoors()
    {
        $this->_car->setPart('rightDoor', new parts\Door());
        $this->_car->setPart('leftDoor', new parts\Door());
        $this->_car->setPart('trunkLid', new Parts\Door());
    }

    public function addEngine()
    {
        $this->_car->setPart('engine', new parts\Engine());
    }

    public function addWheel()
    {
        $this->_car->setPart('wheelLF', new parts\Wheel());
        $this->_car->setPart('wheelRF', new parts\Wheel());
        $this->_car->setPart('wheelLR', new parts\Wheel());
        $this->_car->setPart('wheelRR', new Parts\Wheel());
    }

    public function createVehicle()
    {
        $this->_car = new parts\Car();
    }

    public function getVehicle(): Vehicle
    {
        return $this->_car;
    }
}