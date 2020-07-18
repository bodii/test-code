package com.design_pattern.principle.singleresponsibility;

public class WalkMove extends Foot implements IMove {

    @Override
    public String mode() {
        return "用" + use() + "走";
    }
    
}