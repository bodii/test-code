package com.design_pattern.pattern.creational.singleton;

public class Test3 {
    public static void main(String[] args) {
        Thread thread1 = new Thread(new TLazyCheckDoubleSingleton());
        Thread thread2 = new Thread(new TLazyCheckDoubleSingleton());
        thread1.start();
        thread2.start();

        System.out.println("program end");
    }
}