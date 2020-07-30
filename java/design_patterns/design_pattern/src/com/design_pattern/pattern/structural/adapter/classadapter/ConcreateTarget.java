package com.design_pattern.pattern.structural.adapter.classadapter;

public class ConcreateTarget implements Target {
    @Override
    public void request() {
        System.out.println("concreateTarget target method");
    }
}
