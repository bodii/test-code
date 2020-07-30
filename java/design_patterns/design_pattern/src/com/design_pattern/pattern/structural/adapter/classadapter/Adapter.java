package com.design_pattern.pattern.structural.adapter.classadapter;

public class Adapter extends Adaptee implements Target{
    @Override
    public void request() {
        super.adapteeRequest();
    }
}
