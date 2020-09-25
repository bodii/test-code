<?php
// 测试目录命名空间
namespace design_patterns\builder\test;

// 定义文件目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

use design_patterns\builder\{
    parts\Car,
    parts\Truck,
    TruckBuilder,
    CarBuilder,
    Director,
};

// 引入composer vendor autolaod 文件
require dirname(dirname(dirname(__DIR__))).DIR_SEP.'vendor'.DIR_SEP.'autoload.php';
// 使用phpunit/phpunit PHP测试类
use PHPUnit\Framework\TestCase;


/**
 * DirectorTest 测试类
 */
class DirectorTest extends TestCase
{
    /**
     * TestCanBuildTruck function
     * 卡车测试
     *
     * @return void
     */
    public function testCanBuildTruck()
    {
        $truckBuilder = new TruckBuilder();
        $newVehicle = (new Director())->build($truckBuilder);

        $this->assertInstanceOf(Truck::class, $newVehicle);
    }

    /**
     * TestCanBuildCar function
     * 汽车测试
     *
     * @return void
     */
    public function testCanBuildCar()
    {
        $carBuilder = new CarBuilder();
        $newVehicle = (new Director())->build($carBuilder);

        $this->assertInstanceOf(Car::class, $newVehicle);
    }
}


$director = new DirectorTest();
// 测试汽车
$director->testCanBuildCar();
// 测试卡车
$director->testCanBuildTruck();