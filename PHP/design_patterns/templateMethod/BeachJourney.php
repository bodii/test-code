<?php

// 命名空间：模板方法设计模式
namespace design_patterns\templateMethod;

/**
 * BeachJourney class
 * 海滩旅游类
 */
class BeachJourney extends Journey
{
    /**
     * EnjoyVacation function
     * 海滩旅游享受
     *
     * @return string
     */
    protected function enjoyVacation(): string
    {
        return 'Swimming and sun-bathing';
    }
}