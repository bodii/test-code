package com.design_pattern.pattern.decorator.v2;

public class AbstractDecorator extends ABattercake {
    private ABattercake aBattercake;

    public AbstractDecorator(ABattercake aBattercake) {
        this.aBattercake = aBattercake;
    }

    @Override
    String getDoce() {
        return aBattercake.getDoce();
    }

    @Override
    int cost() {
        return aBattercake.cost();
    }
}
