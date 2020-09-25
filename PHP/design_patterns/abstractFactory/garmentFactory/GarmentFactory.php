<?php
/**
 * 命名空间 abstractFactory\garmentFactory
 */

 namespace abstractFactory\garmentFactory;

 use abstractFactory\Factory;
 use abstractFactory\garmentFactory\GarmentType;
 use abstractFactory\garmentFactory\SalesPatterns;

 /**
  * GarmentFactory class
  * 服装制造厂类
  */
 class GarmentFactory extends Factory
 {
    public $brand = '';

    public function __construct(string $brand='')
    {
        $this->brand = $brand;
    }

    public static function production(GarmentType $type)
    {
        return $type->make();
    }
    
    public static function sales(SalesPatterns $pattern)
    {
        return $pattern->sell();
    }
 }