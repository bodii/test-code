package com.design_pattern.pattern.decorator.v2;

public class EggDecorator extends AbstractDecorator {
    public EggDecorator(ABattercake aBattercake) {
        super(aBattercake);
    }

    @Override
    String getDoce() {
        return super.getDoce() + " 加一个鸡蛋";
    }

    @Override
    int cost() {
        return super.cost() + 1;
    }
}
