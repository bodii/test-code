package com.design_pattern.pattern.structural.decorator.v2;

public class Battercake extends ABattercake {
    @Override
    public String getDose() {
        return "煎饼";
    }

    @Override
    public int cost() {
        return 8;
    }
}
