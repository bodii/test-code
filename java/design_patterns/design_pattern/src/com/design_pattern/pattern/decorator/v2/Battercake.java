package com.design_pattern.pattern.decorator.v2;

public class Battercake extends ABattercake {
    @Override
    public String getDoce() {
        return "煎饼";
    }

    @Override
    public int cost() {
        return 8;
    }
}
