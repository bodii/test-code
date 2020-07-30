package com.design_pattern.pattern.structural.adapter.objectadapter;

public class ConcreateTarget implements Target {
    @Override
    public void request() {
        System.out.println("concreate target adapter");
    }
}
