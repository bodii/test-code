package com.design_pattern.pattern.structural.adapter.objectadapter.example;

public class AC {
    public int input() {
        int input = 220;
        System.out.println("输入交流电：" + input + "V");
        return input;
    }
}
