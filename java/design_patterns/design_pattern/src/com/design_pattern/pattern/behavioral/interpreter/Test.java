package com.design_pattern.pattern.behavioral.interpreter;

public class Test {
    public static void main(String[] args) {
        String str = "6 100 11 + *";
        CalcExpressionParser c = new CalcExpressionParser();
        int result = c.parse(str);
        System.out.printf("解释器计算的结果是： %d \n", result);
    }
}
