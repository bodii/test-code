package com.design_pattern.pattern.structural.decorator.v2;

public class AbstractDecorator extends ABattercake {
    private ABattercake aBattercake;

    public AbstractDecorator(ABattercake aBattercake) {
        this.aBattercake = aBattercake;
    }

    @Override
    String getDose() {
        return aBattercake.getDose();
    }

    @Override
    int cost() {
        return aBattercake.cost();
    }
}
