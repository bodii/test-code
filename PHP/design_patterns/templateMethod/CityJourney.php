<?php

// 命名空间：模板方法设计模式
namespace design_patterns\templateMethod;

/**
 * CityJourney class
 * 城市旅游类
 */
class CityJourney extends Journey
{
    /**
     * EnjoyVacation function
     * 城市旅游享受
     *
     * @return string
     */
    protected function enjoyVacation(): string
    {
        return 'Eat, drink, take photos and sleep';
    }

    /**
     * BuyGift function
     * 购物
     *
     * @return string
     */
    protected function buyGift(): string
    {
        return 'Buy a gift';
    }

}