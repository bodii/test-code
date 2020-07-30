package com.design_pattern.pattern.decorator.v2;

public class SausageDoractor extends AbstractDecorator {
    public SausageDoractor(ABattercake aBattercake) {
        super(aBattercake);
    }

    @Override
    String getDoce() {
        return super.getDoce() + " 加一根火腿";
    }

    @Override
    int cost() {
        return super.cost() + 1;
    }
}
