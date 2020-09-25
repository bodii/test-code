<?php

namespace design_patterns\builder;

use design_patterns\builder\parts\Vehicle;

/**
 * TruckBuilder class
 * 卡车生成器
 */
class TruckBuilder implements BuilderInterface
{
    /**
     * @var parts\Truck
     */
    private $_truck;

    public function addDoors()
    {
        $this->_truck->setPart('rightDoor', new parts\Door());
        $this->_truck->setPart('leftDoor', new parts\Door());
    }

    public function addEngine()
    {
        $this->_truck->setPart('truckEngine', new parts\Engine());
    }

    public function addWheel()
    {
        $this->_truck->setPart('wheel1', new parts\Wheel());
        $this->_truck->setPart('wheel2', new parts\Wheel());
        $this->_truck->setPart('wheel3', new parts\Wheel());
        $this->_truck->setPart('wheel4', new parts\Wheel());
        $this->_truck->setPart('wheel5', new parts\Wheel());
        $this->_truck->setPart('wheel6', new parts\Wheel());
    }

    public function createVehicle()
    {
        $this->_truck = new parts\Truck();
    }

    public function getVehicle(): Vehicle
    {
        return $this->_truck;
    }
}