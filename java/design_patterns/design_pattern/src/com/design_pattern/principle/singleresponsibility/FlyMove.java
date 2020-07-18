package com.design_pattern.principle.singleresponsibility;

public class FlyMove extends Wing implements IMove {

    @Override
    public String mode() {
        return "用" + use() + "飞";
    }
    
}