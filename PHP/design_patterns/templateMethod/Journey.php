<?php

// 命名空间：模板方法设计模式
namespace design_patterns\templateMethod;

/**
 * Journey class
 * 旅行抽象类
 */
abstract class Journey
{
    /**
     * ThingsToDo variable
     * 想要做的事情
     *
     * @var array
     */
    private $thingsToDo;
    
    /**
     * TakeATrip function
     * 去旅行方法
     *
     * @return void
     */
    final public function takeATrip()
    {
        $this->thingsToDo[] = $this->buyAFlight();
        $this->thingsToDo[] = $this->takePlane();
        $this->thingsToDo[] = $this->enjoyVacation();
        $buyGift = $this->buyGift();

        if ($buyGift !== null) {
            $this->thingsToDo[] = $buyGift;
        }

        $this->thingsToDo[] = $this->takePlane();
    }

    /**
     * EnjoyVacation function
     * 享受假期方法
     *
     * @return void
     */
    abstract protected function enjoyVacation(): string;

    /**
     * BuyGift function
     * 购买礼品
     *
     * @return void
     */
    protected function buyGift()
    {
        return null;
    }

    /**
     * BuyFlight function
     * 买飞机票
     *
     * @return string
     */
    private function buyAFlight(): string
    {
        return 'Buy a flight ticket';
    }

    /**
     * TakePlane function
     * 乘坐飞机
     *
     * @return string
     */
    private function takePlane(): string
    {
        return 'Taking the plane';
    }

    /**
     * GetThingsToDo function
     * 获取想做的事情有哪些
     *
     * @return array
     */
    public function getThingsToDo(): array
    {
        return $this->thingsToDo;
    }

}
