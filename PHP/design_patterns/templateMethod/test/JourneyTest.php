<?php

// 命名空间：模板方法设计模式-测试
namespace design_patterns\templateMethod;

use design_patterns\templateMethod\{
    BeachJourney,
    CityJourney
};

// 目录分隔符
if (!defined('DIR_SEP')) {
    define('DIR_SEP', DIRECTORY_SEPARATOR);
}

// 引入vendor下的autoload文件
require_once dirname(dirname(dirname(__DIR__))).DIR_SEP.
'vendor'.DIR_SEP.'autoload.php';

// 使用单元测试类
use PHPUnit\Framework\TestCase;

/**
 * JourneyTest class
 * 旅行测试类
 */
class JourneyTest extends TestCase
{
    /**
     * TestCanGetOnVacationOnTheBeach function
     * 测试获取在海边游行的内容
     *
     * @return void
     */
    public function testCanGetOnVacationOnTheBeach() 
    {
        $beachJourney = new BeachJourney();
        $beachJourney->takeATrip();

        $this->assertEquals(
            [
                'Buy a flight ticket',
                'Taking the plane',
                'Swimming and sun-bathing',
                'Taking the plane',
            ],
            $beachJourney->getThingsToDo()
        );
    }

    /**
     * TestCanGetOnAJourneyToCity function
     * 测试是否能获取城市旅行内容
     *
     * @return void
     */
    public function testCanGetOnAJourneyToCity()
    {
        $cityJourney = new CityJourney();
        $cityJourney->takeATrip();

        $this->assertEquals(
            [
                'Buy a flight ticket',
                'Taking the plane',
                'Eat, drink, take photos and sleep',
                'Buy a gift',
                'Taking the plane'
            ],
            $cityJourney->getThingsToDo()
        );
    }

}