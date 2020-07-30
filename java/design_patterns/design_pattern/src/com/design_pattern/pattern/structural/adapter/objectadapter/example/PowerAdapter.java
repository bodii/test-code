package com.design_pattern.pattern.structural.adapter.objectadapter.example;

public class PowerAdapter implements DC {
    private AC ac220v = new AC();

    @Override
    public int output() {
        int input = ac220v.input();
        int output = input / 44;
        System.out.printf("使用PowerAdapter输入AC: " + input + "V, 输出：" + output + "V");
        return output;
    }
}
