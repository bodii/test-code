package com.design_pattern.pattern.creational.singleton;

public class Test4 {
    public static void main(String[] args) {
        Thread t1 = new Thread(new TStaticInnerClassSingleton());
        Thread t2 = new Thread(new TStaticInnerClassSingleton());
        t1.start();
        t2.start();

        System.out.println("program end");
    }
}