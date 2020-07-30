package com.design_pattern.pattern.structural.decorator.v1;

public class BattercakeWithEggSausage extends BattercakeWithEgg {
    @Override
    public String getDesc() {
        return super.getDesc() + " 加一根香";
    }

    @Override
    public int cost() {
        return super.cost() + 1;
    }
}
