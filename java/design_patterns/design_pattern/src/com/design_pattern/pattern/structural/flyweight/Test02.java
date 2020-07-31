package com.design_pattern.pattern.structural.flyweight;

public class Test02 {
    public static void main(String[] args) {
        Integer a = Integer.valueOf(100);
        Integer b = 100;

        Integer c = Integer.valueOf(1000);
        Integer d = 1000;

        System.out.println("a == b :" + (a == b));
        System.out.println("c == b : " + (c == d));

        System.out.println("a equals b :" + (a.equals(b)));
        System.out.println("c equals b : " + (c.equals(d)));
    }
}
