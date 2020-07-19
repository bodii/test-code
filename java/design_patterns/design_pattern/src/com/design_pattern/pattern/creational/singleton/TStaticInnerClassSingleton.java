package com.design_pattern.pattern.creational.singleton;

public class TStaticInnerClassSingleton implements Runnable {
    @Override
    public void run() {
        StaticInnerClassSingleton singleton = 
            StaticInnerClassSingleton.getInstance();
        System.out.println(Thread.currentThread().getName() + ": "
            + singleton);
    }
}