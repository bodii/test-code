package com.design_pattern.pattern.structural.adapter.objectadapter.example;

public class Test {
    public static void main(String[] args) {
        DC dc5v = new PowerAdapter();
        dc5v.output();
    }
}
