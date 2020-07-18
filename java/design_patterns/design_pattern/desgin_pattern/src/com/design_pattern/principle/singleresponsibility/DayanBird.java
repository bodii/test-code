package com.design_pattern.principle.singleresponsibility;

public class DayanBird  extends FlyMove implements IBird {

    @Override
    public String name() {
        return "大雁" + mode();
    }
    
}