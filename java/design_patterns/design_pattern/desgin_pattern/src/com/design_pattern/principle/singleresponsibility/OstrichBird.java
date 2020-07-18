package com.design_pattern.principle.singleresponsibility;

public class OstrichBird extends WalkMove implements IBird {

    @Override
    public String name() {
        return "鸵鸟" + mode();
    }
    
}