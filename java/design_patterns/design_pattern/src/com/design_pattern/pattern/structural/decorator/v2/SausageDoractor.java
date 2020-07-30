package com.design_pattern.pattern.structural.decorator.v2;

public class SausageDoractor extends AbstractDecorator {
    public SausageDoractor(ABattercake aBattercake) {
        super(aBattercake);
    }

    @Override
    String getDose() {
        return super.getDose() + " 加一根火腿";
    }

    @Override
    int cost() {
        return super.cost() + 1;
    }
}
